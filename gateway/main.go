package main

import (
	pb "gateway/pb"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token,omitempty"`
	Error   string `json:"error,omitempty"`
}

var authServiceClient pb.AuthServiceClient

func main() {
	r := chi.NewRouter()

	// Set up a connection to the gRPC server when the application starts.
	conn, err := grpc.Dial(
		"auth:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("failed to connect to auth service: %v", err)
	}
	defer conn.Close()

	authServiceClient = pb.NewAuthServiceClient(conn)

	// Define your route and handler
	setupRoutes(r, authServiceClient)

	// Start the HTTP server
	log.Println("api gateway running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
