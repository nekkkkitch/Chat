package server

import (
	pb "Chat/pkg/grpc/pb/msgService"
	"Chat/pkg/models"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Config struct {
	Port string `yaml:"auth_port" env-prefix:"AUTHPORT"`
}

type IDBManager interface {
	CheckSameLogin(login string) (bool, error)
	AddUser(user models.User) (int, error)
	GetUserByID(id int) (models.User, error)
	GetUserByLogin(login string) (models.User, error)
}

type IBroker interface {
	SendMessage()
}

type server struct {
	pb.UnimplementedMessagesServer
	db     IDBManager
	broker IBroker
}

type Service struct {
	AuthServer *grpc.Server
	Listener   *net.Listener
	cfg        *Config
	db         IDBManager
	broker     IBroker
}

func New(cfg *Config, db IDBManager, brkr IBroker) (*Service, error) {
	log.Println(cfg.Port)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s", cfg.Port))
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()
	pb.RegisterMessagesServer(s, &server{db: db, broker: brkr})
	log.Printf("Auth server listening at %v\n", lis.Addr())
	return &Service{AuthServer: s, Listener: &lis, cfg: cfg, broker: brkr, db: db}, nil
}

