package security

import (
	"context"
	"net/http"
	"strings"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/config"
)

// Middleware para proteger rutas con JWT
func JWTMiddleware(config *config.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, config.JWT.JWTBearerPrefix) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, config.JWT.JWTBearerPrefix)
		claims, err := ValidateJWT(config, tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Pasar claims al contexto
		ctx := context.WithValue(r.Context(), "userClaims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
