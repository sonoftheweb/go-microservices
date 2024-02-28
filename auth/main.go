package main

import (
	"auth/pb"
	"auth/server"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// connect to DB
	db, err := ConnectDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Initialize Redis
	redis, err := ConnectRedis()
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	authServer := server.NewAuthServer(db, redis)
	pb.RegisterAuthServiceServer(s, authServer)

	// register reflection service on gRPC server
	// remove on production
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
