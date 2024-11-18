package service

import (
	"context"
	"flag"
	"log"

	msgService "Chat/pkg/grpc/pb/msgService"

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

func (c *Client) OpenChat() error {
	_, err := c.client.EnterChat(context.Background(), &msgService.Entering{})
	return err
}