package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
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

	port := fmt.Sprintf(":%s", os.Getenv("USER_SERVICE_PORT"))
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

}
