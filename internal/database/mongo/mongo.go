package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// ConnectMongo establece la conexión con MongoDB
func ConnectMongo(cfg *config.MongoConfig) (*MongoDatabase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		return nil, fmt.Errorf("error conectando a MongoDB: %w", err)
	}

	// Verificar conexión
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error haciendo ping a MongoDB: %w", err)
	}

	log.Println("Conectado a MongoDB en:", cfg.MongoURI)
	return &MongoDatabase{
		Client: client,
		DB:     client.Database(cfg.DatabaseName),
	}, nil
}
