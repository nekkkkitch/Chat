package router

import (
	"Chat/pkg/models"
	"crypto/rsa"
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

type Router struct {
	App    *fiber.App
	Config *Config
	jwt    IJWTManager
	asvc   IAuthService
	msgs   IMsgService
	broker IBroker
}

type Config struct {
	Host string `yaml:"router_host" env-prefix:"ROUTERHOST"`
	Port string `yaml:"router_port" env-prefix:"ROUTERPORT"`
}

type IBroker interface {
	SendMessage(msg models.Message, topic string) error
	OpenMessageTube(ch *chan models.BeautifiedMessage, topic string) error
}

type IAuthService interface {
	Register(user models.User) (*models.AuthData, error)
	Login(user models.User) (*models.AuthData, error)
	UpdateTokens(tokens models.AuthData) (*models.AuthData, error)
}

type IMsgService interface {
	OpenChat() error
}

type IJWTManager interface {
	GetPublicKey() *rsa.PublicKey
	GetIDFromToken(token string) (int, error)
	ValidateToken(c *fiber.Ctx, key string) (bool, error)
	AuthFilter(c *fiber.Ctx) bool
	RefreshFilter(c *fiber.Ctx) bool
}

type Client struct {
	conn *websocket.Conn
	id   int
}

var (
	clients    = make(map[*websocket.Conn]int)
	register   = make(chan Client)
	broadcast  = make(chan models.BeautifiedMessage)
	unregister = make(chan *websocket.Conn)
)

func New(cfg *Config, auservice IAuthService, msgservice IMsgService, jwt IJWTManager, broker IBroker) (*Router, error) {
	app := fiber.New()
	router := Router{App: app, Config: cfg, jwt: jwt, asvc: auservice, msgs: msgservice, broker: broker}
	router.App.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			err := c.Next()
			if err != nil {
				log.Println(err.Error())
				return err
			}
		}
		return nil
	})
	router.App.Use(cors.New(cors.Config{
		AllowHeaders: "X-Access-Token, X-Refresh-Token",
	}))
	router.App.Use(keyauth.New(keyauth.Config{
		Next:         router.jwt.AuthFilter,
		KeyLookup:    "header:X-Access-Token",
		Validator:    router.jwt.ValidateToken,
		ErrorHandler: router.ErrorHandler(),
	}))
	router.App.Use(keyauth.New(keyauth.Config{
		Next:         router.jwt.RefreshFilter,
		KeyLookup:    "header:X-Refresh-Token",
		Validator:    router.jwt.ValidateToken,
		ErrorHandler: router.ErrorHandler(),
	}))
	go runHub()
	go router.broker.OpenMessageTube(&broadcast, "shared_message")
	router.App.Get("/chat", router.RegisterClient())
	router.App.Post("/login", router.Login())
	router.App.Post("/register", router.Register())
	router.App.Get("/refresh", router.UpdateTokens())
	router.App.Get("/ping", Ping)
	err := router.msgs.OpenChat()
	if err != nil {
		return nil, err
	}
	return &router, nil
}

func (r *Router) Listen() {
	r.App.Listen(r.Config.Host + r.Config.Port)
}

func Ping(c *fiber.Ctx) error {
	log.Println("Ping")
	return c.JSON("Ping")
}

func (r *Router) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.User
		err := json.Unmarshal(c.Body(), &user)
		if err != nil {
			return err
		}
		log.Printf("User to login: %s\n", user.Login)
		authData, err := r.asvc.Login(user)
		if err != nil {
			switch err.Error() {
			case "rpc error: code = AlreadyExists desc = login occupied":
				c.Status(fiber.StatusBadRequest)
				return nil
			}
			return err
		}
		log.Printf("Tokens to return:\nAccess token: %s\nRefresh token: %s", authData.AccessToken[:20], authData.RefreshToken[:20])
		c.Context().Response.Header.Set("access_token", authData.AccessToken)
		c.Context().Response.Header.Set("refresh_token", authData.RefreshToken)
		return nil
	}
}

func (r *Router) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.User
		err := json.Unmarshal(c.Body(), &user)
		if err != nil {
			return err
		}
		log.Printf("User to register: %s\n", user.Login)
		authData, err := r.asvc.Register(user)
		if err != nil {
			return err
		}
		log.Printf("Tokens to return:\nAccess token: %s\nRefresh token: %s", authData.AccessToken[:20], authData.RefreshToken[:20])
		c.Context().Response.Header.Set("access_token", authData.AccessToken)
		c.Context().Response.Header.Set("refresh_token", authData.RefreshToken)
		return nil
	}
}

func (r *Router) UpdateTokens() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authData := models.AuthData{AccessToken: c.GetReqHeaders()["X-Access-Token"][0], RefreshToken: c.GetReqHeaders()["X-Refresh-Token"][0]}
		log.Printf("Got tokens:\nAccess token: %s\nRefresh token: %s", authData.AccessToken[:20], authData.RefreshToken[:20])
		authDataResp, err := r.asvc.UpdateTokens(authData)
		if err != nil {
			return err
		}

		log.Printf("Tokens to return:\nAccess token: %s\nRefresh token: %s", authDataResp.AccessToken[:20], authDataResp.RefreshToken[:20])
		c.Context().Response.Header.Set("X-Acess-Token", authDataResp.AccessToken)
		c.Context().Response.Header.Set("X-Refresh-Token", authDataResp.RefreshToken)
		return nil
	}
}

func (r *Router) ErrorHandler() func(c *fiber.Ctx, err error) error {
	return func(c *fiber.Ctx, err error) error {
		log.Println("Bad access token: ", c.GetReqHeaders()["X-Access-Token"])
		log.Println("Bad refresh token: ", c.GetReqHeaders()["X-Refresh-Token"])
		log.Println("Wrong jwts: " + err.Error())
		return err
	}
}

// Создаёт websocket соединение с клиентом, передает его id и соединение в канал регистрации и ожидает сообщение от него(см runHub()).
// При получении сообщение передается в брокер(см kafka.SendMessage())
func (r *Router) RegisterClient() fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		defer func() {
			unregister <- c
			c.Close()
		}()
		access_token := c.Query("access_token")
		log.Println("Trying to connect client, access token: " + access_token)
		if _, err := r.jwt.ValidateToken(nil, access_token); err != nil {
			log.Println("Bad token")
			c.WriteMessage(websocket.TextMessage, []byte("Bad jwt token: "+err.Error()))
			c.Close()
			return
		}
		client := Client{conn: c}
		id, _ := r.jwt.GetIDFromToken(access_token)
		client.id = id
		register <- client
		for {
			_, message, err := c.ReadMessage()
			log.Println("Read message:", string(message))
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("Error when reading message:", err.Error())
				}
				return
			}
			msg := models.Message{}
			err = json.Unmarshal(message, &msg)
			if err != nil {
				log.Println("Cannot unmarshal message: " + err.Error())
				return
			}
			msg.Sender = id
			err = r.broker.SendMessage(msg, "sent_message")
			if err != nil {
				log.Println("Cannot send message: " + err.Error())
				return
			}
		}
	})
}

// Создаёт хаб с клиентами, которые получает от вебсокета. При получении сообщения от брокера возвращает его пользователю с
// id, указанным в сообщении.
// Также, удаляет клиента из мапы клиентов при его отключении
func runHub() {
	log.Println("Opened clients hub")
	for {
		select {
		case connectedClient := <-register:
			clients[connectedClient.conn] = connectedClient.id
			log.Println("Connected client: ", connectedClient.id)
		case connection := <-unregister:
			log.Println("Client disconnected: ", clients[connection])
			delete(clients, connection)
		case message := <-broadcast:
			log.Println("Received message: ", message)
			for conn, client := range clients {
				if client == message.Reciever {
					if err := conn.WriteMessage(websocket.TextMessage, []byte(message.MessageText)); err != nil {
						log.Println("Write message error, cannot sent message to client: " + err.Error())
						unregister <- conn
						conn.WriteMessage(websocket.CloseMessage, []byte{})
						conn.Close()
					}
				}
			}
		}
	}
}
