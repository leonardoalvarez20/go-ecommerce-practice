package auth

import (
	authRepositories "github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/repositories"
	authServices "github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/services"
	authUsecases "github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/usecases"
	security "github.com/leonardoalvarez20/go-ecommerce-practice/internal/security"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthContainer struct {
	AuthRepositories authRepositories.AuthRepository
	AuthService      authServices.AuthService
	LoginUsecase     authUsecases.Usecase
}

func NewAuthContainer(db *mongo.Database, passwordHasher security.PasswordHasher) *AuthContainer {
	authRepositories := authRepositories.NewAuthRepository(db)
	loginUsecase := authUsecases.NewLoginUsecase(authRepositories, passwordHasher)
	authServices := authServices.NewAuthService(loginUsecase)

	return &AuthContainer{
		AuthRepositories: authRepositories,
		AuthService:      authServices,
		LoginUsecase:     loginUsecase,
	}
}
