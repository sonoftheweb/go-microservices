package main

import (
	"gateway/handler"
	pb "gateway/pb"

	"github.com/go-chi/chi"
)

func setupRoutes(router *chi.Mux, authServiceClient pb.AuthServiceClient) {
	authHandler := &handler.AuthServiceHandler{AuthServiceClient: authServiceClient}

	router.Post("/api/auth/login", authHandler.HandleAuthentication)
	router.Post("/api/auth/register", authHandler.HandleRegistration)
}
