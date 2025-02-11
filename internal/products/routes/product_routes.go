package routes

import (
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/handlers"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/services"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router, service *services.ProductService) {
	handler := handlers.CreateProductHandler(service)

	r.HandleFunc("/products", handler.Create).Methods("POST")
	r.HandleFunc("/products/{id}", handler.GetById).Methods("GET")
}
