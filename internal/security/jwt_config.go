package security

import (
	"time"

	"github.com/leonardoalvarez20/go-ecommerce-practice/pkg/utils"
)

type Config struct {
	JWTSecret       []byte        // Cargar desde variable de entorno
	JWTExpiration   time.Duration // Expira en 2 horas
	JWTBearerPrefix string
}

func NewConfig() *Config {
	return &Config{
		JWTSecret:       []byte(utils.GetEnv("JWT_SECRET", "SUPER_SECRET_KEY")),
		JWTExpiration:   time.Hour * 2,
		JWTBearerPrefix: "Bearer",
	}
}
