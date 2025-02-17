package routes

import (
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/handlers"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/services"

	"github.com/gorilla/mux"
)

func AuthUserRoutes(r *mux.Router, service services.AuthService) {
	handler := handlers.NewAuthUserHandler(service)

	r.HandleFunc("/login", handler.Login).Methods("POST")
}
