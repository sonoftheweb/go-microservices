package server

import (
	"auth/pb"
	"context"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	// Placeholder logic for user authentication
	// In a real application, you would query your user database

	if req.Email == "test@example.com" && req.Password == "password" {
		return &pb.AuthResponse{Success: true, Token: "some_token"}, nil
	}

	return &pb.AuthResponse{Success: false, Error: "invalid credentials"}, nil
}
