package routes

import (
	"net/http"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/config"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/security"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/handlers"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/services"

	"github.com/gorilla/mux"
)

func UsersRoutes(r *mux.Router, config *config.Config, service *services.UserService) {
	handler := handlers.CreateUserHandler(service)

	r.HandleFunc("/users", handler.Create).Methods("POST")
	r.Handle("/users/{id}", security.JWTMiddleware(config, http.HandlerFunc(handler.GetById))).Methods("GET")
	r.Handle("/users/{id}", security.JWTMiddleware(config, http.HandlerFunc(handler.UpdateUser))).Methods("PATCH")
}
