package main

import (
	"fmt"
	"log"
	"net/http"

	authContainer "github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth"
	authRouter "github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/routes"
	productsRepositories "github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/repositories"
	productsRoutes "github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/routes"
	productsServices "github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/services"
	usersRepositories "github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/repositories"
	usersRoutes "github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/routes"
	usersServices "github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/services"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/config"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/database/mongo"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/security"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	config := config.NewConfig()
	database, err := mongo.ConnectMongo(&config.Mongo)
	passwordHasher := &security.BcryptHasher{}

	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}

	productsRepo := productsRepositories.NewProductRepository(database.DB)
	productsService := productsServices.CreateProductService(productsRepo)
	productsRoutes.ProductRoutes(r, productsService)

	usersRepo := usersRepositories.NewUserRepository(database.DB)
	usersService := usersServices.CreateUserServices(usersRepo, passwordHasher)
	usersRoutes.UsersRoutes(r, config, usersService)

	authContainer := authContainer.NewAuthContainer(database.DB, passwordHasher)
	authRouter.AuthUserRoutes(r, config, authContainer.AuthServices)

	// Iniciar servidor HTTP en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
