package service

import (
	"context"
	"flag"
	"time"

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
	c      authService.AuthentificationClient
	ctx    context.Context
	conn   *grpc.ClientConn
	cancel context.CancelFunc
}

func New(cfg *Config) (*Client, error) {
	flag.Parse()
	conn, err := grpc.NewClient(cfg.Host+cfg.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := authService.NewAuthentificationClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	return &Client{c: c, ctx: ctx, conn: conn, cancel: cancel}, nil
}

func (c *Client) Register(user models.User) (*models.AuthData, error) {
	authDataGed, err := c.c.Register(c.ctx, &authService.User{Login: user.Login, Password: user.Password})
	if err != nil {
		return nil, err
	}
	return &models.AuthData{AccessToken: authDataGed.AccessToken, RefreshToken: authDataGed.RefreshToken}, nil
}

func (c *Client) Login(user models.User) (*models.AuthData, error) {
	authDataGed, err := c.c.Login(c.ctx, &authService.User{Login: user.Login, Password: user.Password})
	if err != nil {
		return nil, err
	}
	return &models.AuthData{AccessToken: authDataGed.AccessToken, RefreshToken: authDataGed.RefreshToken}, nil
}

func (c *Client) UpdateTokens(tokens models.AuthData) (*models.AuthData, error) {
	authDataGed, err := c.c.UpdateTokens(c.ctx, &authService.AuthData{AccessToken: tokens.AccessToken, RefreshToken: tokens.RefreshToken})
	if err != nil {
		return nil, err
	}
	return &models.AuthData{AccessToken: authDataGed.AccessToken, RefreshToken: authDataGed.RefreshToken}, nil
}
