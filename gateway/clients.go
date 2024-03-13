package main

import (
	pb "gateway/pb"

	"github.com/go-chi/chi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SetupAuthClient(r *chi.Mux) (*grpc.ClientConn, error) {
	// Set up a connection to the gRPC server when the application starts.
	conn, err := grpc.Dial(
		"auth:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	authServiceClient = pb.NewAuthServiceClient(conn)

	// Define your route and handler
	setupAuthRoutes(r, authServiceClient)

	return conn, nil
}
