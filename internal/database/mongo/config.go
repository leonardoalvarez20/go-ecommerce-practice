package mongo

import "os"

type Config struct {
	MongoURI     string
	DatabaseName string
}

func NewConfig() Config {
	return Config{
		MongoURI:     getEnv("MONGO_URI", "uri"),
		DatabaseName: getEnv("MONGO_DB_NAME", "ecommerce"),
	}
}

func getEnv(key, fallback string) string {
 	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
