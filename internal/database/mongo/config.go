package mongo

import (
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/shared/utils"
)

type Config struct {
	MongoURI     string
	DatabaseName string
}

func NewConfig() Config {

	return Config{
		MongoURI:     utils.GetEnv("MONGO_URI", "uri"),
		DatabaseName: utils.GetEnv("MONGO_DB_NAME", "ecommerce"),
	}
}
