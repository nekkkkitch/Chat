package server

import (
	pb "Chat/pkg/grpc/pb/msgService"
	"Chat/pkg/models"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Config struct {
	Port string `yaml:"msg_port" env-prefix:"AUTHPORT"`
}

type IDBManager interface {
	CheckSameLogin(login string) (bool, error)
	AddUser(user models.User) (int, error)
	GetUserByID(id int) (models.User, error)
	GetUserByLogin(login string) (models.User, error)
}

type IBroker interface {
	SendMessage(msg models.BeautifiedMessage, topic string) error
	OpenMessageTube(ch *chan models.Message, topic string) error
}

type server struct {
	pb.UnimplementedMessagesServer
	db     IDBManager
	broker IBroker
}

type Service struct {
	MsgServer *grpc.Server
	Listener  *net.Listener
	cfg       *Config
	db        IDBManager
	broker    IBroker
}

var (
	broadcast = make(chan models.Message)
)

func New(cfg *Config, db IDBManager, brkr IBroker) (*Service, error) {
	log.Println(cfg.Port)
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()
	pb.RegisterMessagesServer(s, &server{db: db, broker: brkr})
	svc := Service{MsgServer: s, Listener: &lis, cfg: cfg, broker: brkr, db: db}
	log.Printf("Auth server listening at %v\n", lis.Addr())
	return &svc, nil
}

func (s *server) SendMessage(_ context.Context, in *pb.Message) (*pb.Status, error) {
	log.Println("Got message:", in)
	beautifiedmsg := models.BeautifiedMessage{Sender: int(in.Sender)}
	sender, err := s.db.GetUserByID(int(in.Sender))
	if err != nil {
		return nil, err
	}
	reciever, err := s.db.GetUserByLogin(in.Reciever)
	if err != nil {
		return nil, err
	}
	beautifiedmsg.Reciever = reciever.ID
	beautifiedmsg.MessageText = fmt.Sprintf("Sender: %v\n%v", sender.Login, in.Text)
	s.broker.SendMessage(beautifiedmsg, "shared_message")
	return nil, nil
}

func (s *server) EnterChat(_ context.Context, _ *pb.Entering) (*pb.Status, error) {
	log.Println("Opening chat")
	go s.broker.OpenMessageTube(&broadcast, "sent_message")
	go s.ReadMessages()
	return nil, nil
}

func (s *server) ReadMessages() {
	for msg := range broadcast {
		log.Println("Got message:", msg)
		_, err := s.SendMessage(context.Background(), &pb.Message{Sender: int32(msg.Sender), Reciever: msg.Reciever, Text: msg.MessageText})
		if err != nil {
			log.Println("Failed to send message: " + err.Error())
		}
	}
}
