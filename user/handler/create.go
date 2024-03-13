package handler

import (
	"context"
	"user/pb"
)

func (h *Handler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	// logic to create user here

	return &pb.UserResponse{Success: true}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *pb.UpdateDeleteViewUserRequest) (*pb.UserResponse, error) {
	// logic to update user here

	return &pb.UserResponse{Success: true}, nil
}

//func (h *Handler) View
