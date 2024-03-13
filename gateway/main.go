package main

import (
	pb "gateway/pb"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

var (
	authServiceClient pb.AuthServiceClient
)

func main() {
	r := chi.NewRouter()

	conn, err := SetupAuthClient(r)
	if err != nil {
		log.Fatalf("failed to setup auth client: %v", err)
	}
	defer conn.Close()

	// Start the HTTP server
	log.Println("api gateway running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
