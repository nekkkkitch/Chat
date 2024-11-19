package service

import (
	"context"
	"encoding/json"
	"flag"
	"log"

	msgService "Chat/pkg/grpc/pb/msgService"
	"Chat/pkg/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Host string `yaml:"msg_host" env-prefix:"AUTHHOST"`
	Port string `yaml:"msg_port" env-prefix:"AUTHPORT"`
}

type Client struct {
	client msgService.MessagesClient
	conn   *grpc.ClientConn
}

// Создание клиента для msgService
func New(cfg *Config) (*Client, error) {
	flag.Parse()
	conn, err := grpc.NewClient(cfg.Host+cfg.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	log.Println("Connecting to msg on " + cfg.Host + cfg.Port)
	c := msgService.NewMessagesClient(conn)
	return &Client{client: c, conn: conn}, nil
}

// Вызов функции для получения всей переписки
func (c *Client) GetMessages(msg models.Message) ([]models.BeautifiedMessage, error) {
	resp, err := c.client.GetMessages(context.Background(), &msgService.Message{Sender: int32(msg.Sender), Reciever: msg.Reciever})
	if err != nil {
		return nil, err
	}
	msgs := []models.BeautifiedMessage{}
	err = json.Unmarshal(resp.JsonedChat, &msgs)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}
