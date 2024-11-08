package service

import (
	"context"
	"flag"
	"log"

	authService "Chat/pkg/grpc/pb"
	"Chat/pkg/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Host string `yaml:"auth_host" env-prefix:"AUTHHOST"`
	Port string `yaml:"auth_port" env-prefix:"AUTHPORT"`
}

type Client struct {
	client authService.AuthentificationClient
	conn   *grpc.ClientConn
}

func New(cfg *Config) (*Client, error) {
	flag.Parse()
	conn, err := grpc.NewClient(cfg.Host+cfg.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := authService.NewAuthentificationClient(conn)
	log.Println("Connecting to aus on " + cfg.Host + cfg.Port)
	return &Client{client: c, conn: conn}, nil
}

func (c *Client) Register(user models.User) (*models.AuthData, error) {
	authDataGed, err := c.client.Register(context.Background(), &authService.User{Login: user.Login, Password: user.Password})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.AuthData{AccessToken: authDataGed.AccessToken, RefreshToken: authDataGed.RefreshToken}, nil
}

func (c *Client) Login(user models.User) (*models.AuthData, error) {
	authDataGed, err := c.client.Login(context.Background(), &authService.User{Login: user.Login, Password: user.Password})
	if err != nil {
		return nil, err
	}
	return &models.AuthData{AccessToken: authDataGed.AccessToken, RefreshToken: authDataGed.RefreshToken}, nil
}

func (c *Client) UpdateTokens(tokens models.AuthData) (*models.AuthData, error) {
	authDataGed, err := c.client.UpdateTokens(context.Background(), &authService.AuthData{AccessToken: tokens.AccessToken, RefreshToken: tokens.RefreshToken})
	if err != nil {
		return nil, err
	}
	return &models.AuthData{AccessToken: authDataGed.AccessToken, RefreshToken: authDataGed.RefreshToken}, nil
}
