package server

import (
	"auth/handler"
	"auth/pb"
	"context"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	Handler *handler.Handler
}

func NewAuthServer(db *gorm.DB, redisClient *redis.Client) *AuthServer {
	return &AuthServer{
		Handler: handler.NewHandler(db, redisClient),
	}
}

func (s *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	return s.Handler.Register(ctx, req)
}

func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	return s.Handler.Login(ctx, req)
}
