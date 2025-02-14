package main

import (
	"fmt"
	"log"
	"net/http"

	userAuthRepositories "github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/repositories"
	userAuthRoutes "github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/routes"
	userAuthServices "github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/services"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/database/mongo"
	productsRepositories "github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/repositories"
	productsRoutes "github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/routes"
	productsServices "github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/services"
	usersRepositories "github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/repositories"
	usersRoutes "github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/routes"
	usersServices "github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/services"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	database, err := mongo.ConnectMongo()
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}

	productsRepo := productsRepositories.NewProductRepository(database.DB)
	productsService := productsServices.CreateProductService(productsRepo)
	productsRoutes.ProductRoutes(r, productsService)

	usersRepo := usersRepositories.NewUserRepository(database.DB)
	usersService := usersServices.CreateUserServices(usersRepo)
	usersRoutes.UsersRoutes(r, usersService)

	userAuthRepo := userAuthRepositories.NewAuthRepository(database.DB)
	userAuthService := userAuthServices.NewAuthService(userAuthRepo)
	userAuthRoutes.AuthUserRoutes(r, userAuthService)

	// Iniciar servidor HTTP en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
