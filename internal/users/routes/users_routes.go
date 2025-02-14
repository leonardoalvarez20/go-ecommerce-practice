package routes

import (
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/handlers"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/services"

	"github.com/gorilla/mux"
)

func UsersRoutes(r *mux.Router, service *services.UserService) {
	handler := handlers.CreateUserHandler(service)

	r.HandleFunc("/users", handler.Create).Methods("POST")
	r.HandleFunc("/users/{id}", handler.GetById).Methods("GET")
	r.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PATCH")
}
