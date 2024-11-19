package server

import (
	pb "Chat/pkg/grpc/pb/msgService"
	"Chat/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Config struct {
	Port string `yaml:"msg_port" env-prefix:"AUTHPORT"`
}

type IDBManager interface {
	CheckSameLogin(login string) (bool, error)
	AddUser(user models.User) (int, error)
	GetUserByID(id int) (models.User, error)
	GetUserByLogin(login string) (models.User, error)
	AddMessage(msg models.BeautifiedMessage, sendTime time.Time) error
	GetMessages(firstId, secondId int) ([]models.BeautifiedMessage, error)
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

// Создание сервера сервиса сообщений
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
	svc.EnterChat()
	return &svc, nil
}

// Получение и возврат из БД истории переписки пользователей
func (s *server) GetMessages(_ context.Context, in *pb.Message) (*pb.Chat, error) {
	reciever, err := s.db.GetUserByLogin(in.Reciever)
	if err != nil {
		return nil, err
	}
	msgs, err := s.db.GetMessages(int(in.Sender), reciever.ID)
	if err != nil {
		return nil, err
	}
	jsonedMsgs, err := json.Marshal(msgs)
	if err != nil {
		return nil, err
	}
	log.Println("Successfully got messages between user, returning them")
	return &pb.Chat{JsonedChat: jsonedMsgs}, nil
}

// Отправка измененного сообщения в брокер и сохранение сообщения в БД
func (s *Service) SendMessage(_ context.Context, in *pb.Message) error {
	log.Println("Got message:", in)
	beautifiedmsg := models.BeautifiedMessage{Sender: int(in.Sender)}
	sender, err := s.db.GetUserByID(int(in.Sender))
	if err != nil {
		return err
	}
	reciever, err := s.db.GetUserByLogin(in.Reciever)
	if err != nil {
		return err
	}
	beautifiedmsg.Reciever = reciever.ID
	beautifiedmsg.Text = fmt.Sprintf("%v\n%v:\n%v", in.SendTime.AsTime().Format("02.01 15:04"), sender.Login, in.Text)
	log.Println("Sending message:", beautifiedmsg)
	err = s.db.AddMessage(beautifiedmsg, in.SendTime.AsTime())
	if err != nil {
		log.Println("Failed to save message:", err.Error())
	}
	err = s.broker.SendMessage(beautifiedmsg, "shared_message")
	if err != nil {
		log.Println("Failed to send message via broker:", err.Error())
	}
	return nil
}

// Запуск чтения сообщений из брокера
func (s *Service) EnterChat() error {
	log.Println("Opening chat")
	go s.broker.OpenMessageTube(&broadcast, "sent_message")
	go s.ReadMessages()
	return nil
}

// Чтение сообщений из брокера
func (s *Service) ReadMessages() {
	for msg := range broadcast {
		log.Println("Got message:", msg)
		err := s.SendMessage(context.Background(), &pb.Message{Sender: int32(msg.Sender), Reciever: msg.Reciever,
			Text: msg.Text, SendTime: timestamppb.New(msg.SendTime)})
		if err != nil {
			log.Println("Failed to send message: " + err.Error())
		}
	}
}
