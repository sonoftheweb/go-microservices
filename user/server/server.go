package server

import (
	"context"
	"user/handler"
	pb "user/pb"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	Handler *handler.Handler
}

func NewUserServer(db *gorm.DB, redisClient *redis.Client) *UserServer {
	return &UserServer{
		Handler: handler.NewHandler(db, redisClient),
	}
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	return s.Handler.CreateUser(ctx, req)
}
