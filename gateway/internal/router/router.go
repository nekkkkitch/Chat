package router

import (
	"Chat/pkg/models"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

type App struct {
	app fiber.App
	cfg *Config
	jwt IJWTManager
	svc IService
}

type Config struct {
	Port string `yaml:"gateway_port" env-prefix:"GATEWAYPORT"`
}

type IService interface {
	Register(user models.User)(int, error)
	Login(user models.User)
}

type IJWTManager interface {
	
}

func New(cfg *Config, service IService, jwt IJWTManager) (*App, error) {
	app := fiber.New()
	app.Use(keyauth.New(keyauth.Config{
		KeyLookup: "cookie:access_token",
		Validator: ,
	})

	)
}
