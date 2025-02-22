package usecases

import (
	"context"
	"errors"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/converters"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/repositories"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/security"
)

type loginUsecase struct {
	authRepository repositories.AuthRepository
	passwordHasher security.PasswordHasher
}

func NewLoginUsecase(authRepository repositories.AuthRepository, passwordHasher security.PasswordHasher) *loginUsecase {
	return &loginUsecase{
		authRepository: authRepository,
		passwordHasher: passwordHasher,
	}
}

func (s *loginUsecase) Execute(ctx context.Context, loginRequest *dtos.LoginRequest) (*dtos.LoginResponse, error) {
	authUser, err := s.authRepository.GetUserByEmail(ctx, loginRequest.Email)
	if err != nil || !s.passwordHasher.ComparePasswords(authUser.Password, loginRequest.Password) {
		return nil, errors.New("invalid credentials")
	}
	response := converters.ToAuthUserResponse(authUser)
	return &response, nil
}
