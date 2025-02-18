package config

import (
	"sync"
	"time"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/shared/utils"
)

type Config struct {
	Server ServerConfig
	Mongo  MongoConfig
	JWT    JWTConfig
}

type ServerConfig struct {
	Port      string
	AppName   string
	DebugMode bool
}

type MongoConfig struct {
	MongoURI     string
	DatabaseName string
}

type JWTConfig struct {
	JWTSecret       []byte
	JWTExpiration   time.Duration
	JWTBearerPrefix string
}

var config *Config
var once sync.Once

func NewConfig() *Config {
	once.Do(func() {
		config = &Config{
			Server: ServerConfig{
				Port:      utils.GetEnv("PORT", "8080"),
				AppName:   utils.GetEnv("APP_NAME", "Go Ecommerce Practice"),
				DebugMode: utils.GetEnv("DEBUG_MODE", "true") == "true",
			},
			Mongo: MongoConfig{
				MongoURI:     utils.GetEnv("MONGO_URI", "mongodb://localhost:27017"),
				DatabaseName: utils.GetEnv("MONGO_DB_NAME", "ecommerce"),
			},
			JWT: JWTConfig{
				JWTSecret:       []byte(utils.GetEnv("JWT_SECRET", "SUPER_SECRET_KEY")),
				JWTExpiration:   time.Hour * 2,
				JWTBearerPrefix: "Bearer ",
			},
		}
	})
	return config
}
