package auth

import (
	authRepositories "github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/repositories"
	authServices "github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/services"
	authUsecases "github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/usecases"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthContainer struct {
	AuthRepositories authRepositories.AuthRepository
	AuthServices authServices.AuthService
	LoginUsecase authUsecases.Usecase
}

func NewAuthContainer(db *mongo.Database) *AuthContainer{
	authRepositories := authRepositories.NewAuthRepository(db)
	loginUsecase := authUsecases.NewLoginUsecase(authRepositories)
	authServices := authServices.NewAuthService(loginUsecase)

	return &AuthContainer{
		AuthRepositories: authRepositories,
		AuthServices: authServices,
		LoginUsecase: loginUsecase,
	}
}