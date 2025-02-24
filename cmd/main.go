package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal"
	authContainer "github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth"
	authRouter "github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/routes"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/config"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/database/mongo"
	productsContainer "github.com/leonardoalvarez20/go-ecommerce-practice/internal/products"
	productsRoutes "github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/routes"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/security"
	usersContainer "github.com/leonardoalvarez20/go-ecommerce-practice/internal/users"
	usersRoutes "github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/routes"
)

func main() {
	r := mux.NewRouter()
	config := config.NewConfig()
	mongoDatabase := initMongoDatabse(config)
	container := intitContainer(mongoDatabase)
	initRoutes(r, config, container)
	initServer(r, config)
}

func initMongoDatabse(config *config.Config) *mongo.MongoDatabase {
	fmt.Println("Iniciando MongoDatabase...")
	database, err := mongo.ConnectMongo(&config.Mongo)
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}

	return database
}

func intitContainer(mongoDatabase *mongo.MongoDatabase) *internal.Container {
	fmt.Println("Iniciando contenedor...")
	passwordHasher := &security.BcryptHasher{}
	productsContainer := productsContainer.NewProductsContainer(mongoDatabase.DB)
	usersContainer := usersContainer.NewUsersContainer(mongoDatabase.DB, passwordHasher)
	authContainer := authContainer.NewAuthContainer(mongoDatabase.DB, passwordHasher)
	return &internal.Container{
		Auth:     authContainer,
		Products: productsContainer,
		Users:    usersContainer,
	}
}

func initRoutes(r *mux.Router, config *config.Config, container *internal.Container) {
	fmt.Println("Iniciando rutas...")
	productsRoutes.ProductRoutes(r, container.Products.ProductService)
	usersRoutes.UsersRoutes(r, config, container.Users.UserService)
	authRouter.AuthUserRoutes(r, config, container.Auth.AuthService)
}

func initServer(r *mux.Router, cfg *config.Config) {
	fmt.Println("Iniciando " + cfg.Server.AppName + "...")
	// Iniciar servidor HTTP en el puerto 8080
	log.Println("Servidor escuchando en http://localhost:" + cfg.Server.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, r))
}
