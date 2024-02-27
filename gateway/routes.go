package main

import (
	"github.com/go-chi/chi"
)

func setupRoutes(r *chi.Mux) {
	r.Post("/api/auth/login", loginHandler)
}
