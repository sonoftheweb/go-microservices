package handler

import (
	"auth/model"
	"auth/pb"
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return nil, err
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	// Save the user in the DB
	result := h.DB.Create(&user)
	if result.Error != nil {
		log.Printf("failed to register user: %v", result.Error)
		return &pb.AuthResponse{Success: false, Error: "failed to register user"}, nil
	}

	return &pb.AuthResponse{Success: true}, nil
}
