package router

import (
	"Chat/pkg/models"
	"crypto/rsa"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

type Router struct {
	App    *fiber.App
	Config *Config
	jwt    IJWTManager
	asvc   IAuthService
}

type Config struct {
	Host string `yaml:"router_host" env-prefix:"ROUTERHOST"`
	Port string `yaml:"router_port" env-prefix:"ROUTERPORT"`
}

type IAuthService interface {
	Register(user models.User) (*models.AuthData, error)
	Login(user models.User) (*models.AuthData, error)
	UpdateTokens(tokens models.AuthData) (*models.AuthData, error)
}

type IJWTManager interface {
	GetPublicKey() *rsa.PublicKey
	ValidateToken(c *fiber.Ctx, key string) (bool, error)
	AuthFilter(c *fiber.Ctx) bool
	RefreshFilter(c *fiber.Ctx) bool
}

func New(cfg *Config, auservice IAuthService, jwt IJWTManager) *Router {
	app := fiber.New()
	router := &Router{App: app, Config: cfg, jwt: jwt, asvc: auservice}
	/*
		app.Use(jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{
				JWTAlg: jwtware.RS256,
				Key:    jwt.GetPublicKey(),
			},
		}))
	*/
	app.Use(keyauth.New(keyauth.Config{
		Next:         router.jwt.AuthFilter,
		KeyLookup:    "header:X-Access-Token",
		Validator:    router.jwt.ValidateToken,
		ErrorHandler: router.ErrorHandler(),
	}))
	app.Use(keyauth.New(keyauth.Config{
		Next:         router.jwt.RefreshFilter,
		KeyLookup:    "header:X-Access-Token",
		Validator:    router.jwt.ValidateToken,
		ErrorHandler: router.ErrorHandler(),
	}))
	router.App.Post("/login", router.Login())
	router.App.Post("/register", router.Register())
	router.App.Get("/refresh", router.UpdateTokens())
	router.App.Get("/", Ping)
	return router
}

func (r *Router) Listen() {
	r.App.Listen(r.Config.Host + r.Config.Port)
}

func Ping(c *fiber.Ctx) error {
	log.Println("Ping")
	return c.JSON("aboba")
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
		authData := models.AuthData{}
		if err := c.CookieParser(&authData); err != nil {
			return err
		}
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
		return err
	}
}
