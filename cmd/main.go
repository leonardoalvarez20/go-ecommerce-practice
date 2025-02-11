package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/database/mongo"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/repositories"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/routes"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/services"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	database, err := mongo.ConnectMongo()
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}

	productRepo := repositories.NewProductRepository(database.DB)
	productService := services.CreateProductService(productRepo)
	routes.ProductRoutes(r, productService)

	// Iniciar servidor HTTP en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
