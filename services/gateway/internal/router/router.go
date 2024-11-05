package router

import (
	"Chat/pkg/models"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

type Router struct {
	app *fiber.App
	cfg *Config
	jwt IJWTManager
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
	ValidateToken(c *fiber.Ctx, key string) (bool, error)
	AuthFilter(c *fiber.Ctx) bool
	RefreshFilter(c *fiber.Ctx) bool
}

var (
	asvc IAuthService
)

func New(cfg *Config, auservice IAuthService, jwt IJWTManager) *Router {
	app := fiber.New()
	app.Use(keyauth.New(keyauth.Config{
		Next:      jwt.AuthFilter,
		KeyLookup: "cookie:access_token",
		Validator: jwt.ValidateToken,
	}))
	app.Use(keyauth.New(keyauth.Config{
		Next:      jwt.RefreshFilter,
		KeyLookup: "cookie:refresh_token",
		Validator: jwt.ValidateToken,
	}))
	app.Post("/login", Login)
	app.Post("/register", Register)
	app.Post("/refresh", UpdateTokens)
	asvc = auservice
	return &Router{app: app, cfg: cfg, jwt: jwt}
}

func (r *Router) Listen() {
	r.app.Listen(r.cfg.Host + r.cfg.Port)
}

func Login(c *fiber.Ctx) error {
	var user models.User
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		return err
	}
	log.Printf("User to login: %s\n", user.Login)
	authData, err := asvc.Login(user)
	if err != nil {
		return err
	}
	log.Printf("Tokens to return:\nAccess token: %s\nRefresh token: %s", authData.AccessToken[:20], authData.RefreshToken[:20])
	return c.JSON(authData)
}

func Register(c *fiber.Ctx) error {
	var user models.User
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		return err
	}
	log.Printf("User to register: %s\n", user.Login)
	authData, err := asvc.Register(user)
	if err != nil {
		return err
	}
	log.Printf("Tokens to return:\nAccess token: %s\nRefresh token: %s", authData.AccessToken[:20], authData.RefreshToken[:20])
	return c.JSON(authData)
}

func UpdateTokens(c *fiber.Ctx) error {
	authData := models.AuthData{}
	if err := c.CookieParser(&authData); err != nil {
		return err
	}
	log.Printf("Got tokens:\nAccess token: %s\nRefresh token: %s", authData.AccessToken[:20], authData.RefreshToken[:20])
	authDataResp, err := asvc.UpdateTokens(authData)
	if err != nil {
		return err
	}
	log.Printf("Tokens to return:\nAccess token: %s\nRefresh token: %s", authDataResp.AccessToken[:20], authDataResp.RefreshToken[:20])
	return c.JSON(authDataResp) //переписать возвращение токенов в c.Cookie
}
