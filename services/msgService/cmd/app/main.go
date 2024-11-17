package main

import (
	pg "Chat/pkg/db"
	"Chat/services/msgService/internal/kafka"
	server "Chat/services/msgService/internal/server"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DBConfig        *pg.Config     `yaml:"db" env-prefix:"DB_"`
	MsgServerConfig *server.Config `yaml:"msg"`
	KFKConfig       *kafka.Config  `yaml:"kfk"`
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
	log.Println(cfg.DBConfig)
	db, err := pg.New(cfg.DBConfig)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connected successfully")
	broker, err := kafka.New(cfg.KFKConfig)
	if err != nil {
		log.Fatalln("Failed to establish broker connection: " + err.Error())
	}
	log.Println("Broker connected successfully")
	service, err := server.New(cfg.MsgServerConfig, db, broker)
	if err != nil {
		log.Fatalln("Failed to create service:", err)
	}
	if err := service.MsgServer.Serve(*service.Listener); err != nil {
		log.Fatalln("Failed to start service:", err)
	}
	log.Println("Service connected successfully")
}
