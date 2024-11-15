package main

import (
	"Chat/pkg/jwt"
	kfk "Chat/services/gateway/internal/kafka"
	rtr "Chat/services/gateway/internal/router"
	aus "Chat/services/gateway/internal/services/authService"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	JWTConfig *jwt.Config `yaml:"jwt" env-prefix:"JWT_"`
	AUSConfig *aus.Config `yaml:"aus" env-prefix:"AUS_"`
	RTRConfig *rtr.Config `yaml:"rtr" env-prefix:"RTR_"`
	KFKConfig *kfk.Config `yaml:"kfk"`
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
	authService, err := aus.New(cfg.AUSConfig)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Auth service connected successfully")
	key, err := authService.GetPrivateKey()
	if err != nil {
		log.Fatalln("Problem with getting key: " + err.Error())
	}
	jwt, err := jwt.NewWithKey(cfg.JWTConfig, key)
	if err != nil {
		log.Fatalln("Failed to create jwt: " + err.Error())
	}
	log.Println("JWT created successfully")
	broker, err := kfk.New(cfg.KFKConfig)
	if err != nil {
		log.Fatalln("Failed to establish broker connection: " + err.Error())
	}
	log.Println("Broker connected successfully")
	router := rtr.New(cfg.RTRConfig, authService, &jwt, broker)
	log.Printf("Router is listening on %v:%v\n", router.Config.Host, router.Config.Port)
	router.Listen()
}
