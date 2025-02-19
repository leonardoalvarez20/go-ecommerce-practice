package routes

import (
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/handlers"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/services"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/config"

	"github.com/gorilla/mux"
)

func AuthUserRoutes(r *mux.Router, cfg *config.Config, service services.AuthService) {
	handler := handlers.NewAuthUserHandler(service, cfg)

	r.HandleFunc("/login", handler.Login).Methods("POST")
}
