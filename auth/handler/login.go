package handler

import (
	"auth/model"
	"auth/pb"
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const sessionTokenLength = 32

func generateSessionToken() string {
	randomBytes := make([]byte, sessionTokenLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("failed to generate random session token: ", err)
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}

func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	// Find user by email
	var user model.User
	if err := h.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return &pb.AuthResponse{Success: false, Error: "user not found"}, nil
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return &pb.AuthResponse{Success: false, Error: "invalid credentials"}, nil
	}

	// Generate a session token
	sessionToken := generateSessionToken()

	// Store session token in Redis with user identifier
	err := h.RedisClient.Set(ctx, sessionToken, req.Email, 24*time.Hour).Err()
	if err != nil {
		log.Printf("error saving session to Redis: %v", err)
		return &pb.AuthResponse{Success: false, Error: "failed to create session"}, nil
	}

	return &pb.AuthResponse{Success: true, Token: sessionToken}, nil
}
