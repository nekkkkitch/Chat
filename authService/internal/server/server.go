package authserver

import (
	"context"
	"fmt"
	"log"
	"net"

	"Chat/pkg/crypt"
	pb "Chat/pkg/grpc/pb"
	"Chat/pkg/models"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Config struct {
	Port string `yaml:"auth_port" env-prefix:"AUTHPORT"`
}

type IJWTManager interface {
	CreateTokens(user_id int) (string, string, error)
	GetIDFromToken(token string) (int, error)
}

type IDBManager interface {
	CheckSameLogin(login string) (bool, error)
	AddUser(user models.User) (int, error)
	GetUserByID(id int) (models.User, error)
	GetUserByLogin(login string) (models.User, error)
}

type server struct {
	pb.UnimplementedAuthentificationServer
	jwt IJWTManager
	db  IDBManager
}

type Service struct {
	AuthServer *grpc.Server
	Listener   *net.Listener
	cfg        *Config
	jwt        IJWTManager
	db         IDBManager
}

func NewService(cfg *Config, jwt IJWTManager, db IDBManager) (*Service, error) {
	log.Println(cfg.Port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()
	pb.RegisterAuthentificationServer(s, &server{jwt: jwt, db: db})
	log.Printf("Auth server listening at %v\n", lis.Addr())
	return &Service{AuthServer: s, Listener: &lis, cfg: cfg, jwt: jwt, db: db}, nil
}

func (s *server) Register(_ context.Context, in *pb.User) (*pb.AuthData, error) {
	if in.Login == "" || in.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing login or password")
	}
	if same, err := s.db.CheckSameLogin(in.Login); err != nil || same {
		if same {
			return nil, status.Errorf(codes.AlreadyExists, "login occupied")
		}
		return nil, status.Errorf(codes.Internal, "%v", err)
	}
	cryptedPassword, err := crypt.CryptPassword(in.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}
	id, err := s.db.AddUser(models.User{Login: in.Login, Password: string(cryptedPassword)})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}
	access, refresh, err := s.jwt.CreateTokens(id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}
	return &pb.AuthData{AccessToken: access, RefreshToken: refresh}, nil
}

func (s *server) Login(_ context.Context, in *pb.User) (*pb.AuthData, error) {
	if in.Login == "" || in.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing login or password")
	}
	if same, err := s.db.CheckSameLogin(in.Login); err != nil || !same {
		if !same {
			return nil, status.Errorf(codes.AlreadyExists, "login does not exist")
		}
		return nil, status.Errorf(codes.Internal, "%v", err)
	}
	user, err := s.db.GetUserByLogin(in.Login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		if err.Error() == bcrypt.ErrMismatchedHashAndPassword.Error() {
			return nil, status.Errorf(codes.InvalidArgument, "wrong password")
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	access, refresh, err := s.jwt.CreateTokens(user.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}
	return &pb.AuthData{AccessToken: access, RefreshToken: refresh}, nil
}

func (s *server) UpdateTokens(_ context.Context, in *pb.AuthData) (*pb.AuthData, error) {
	user_id, err := s.jwt.GetIDFromToken(in.AccessToken)
	if user_id == 0 {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	access, refresh, err := s.jwt.CreateTokens(user_id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}
	return &pb.AuthData{AccessToken: access, RefreshToken: refresh}, nil
}
