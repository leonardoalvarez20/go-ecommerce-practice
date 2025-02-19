package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/config"
)

// Generar un nuevo token JWT
func GenerateJWT(config *config.Config, username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(config.JWT.JWTExpiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWT.JWTSecret)
}

// Validar JWT y extraer claims
func ValidateJWT(config *config.Config, tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.JWT.JWTSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
