package main

import (
	"Chat/pkg/jwt"
	rtr "Chat/services/gateway/internal/router"
	aus "Chat/services/gateway/internal/services/authService"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	JWTConfig *jwt.Config `yaml:"jwt" env-prefix:"JWT_"`
	AUSConfig *aus.Config `yaml:"aus" env-prefix:"AUS_"`
	RTRConfig *rtr.Config `yaml:"rtr" env-prefix:"RTR_"`
}

func readConfig(filename string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(filename, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func main() {
	cfg, err := readConfig("./cfg.yml")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Config file read successfully")
	jwt, err := jwt.New(cfg.JWTConfig)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("JWT created successfully")
	authService, err := aus.New(cfg.AUSConfig)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Auth service connected successfully")
	router := rtr.New(cfg.RTRConfig, authService, &jwt)
	router.Listen()
	log.Println("Router started successfully")
}
