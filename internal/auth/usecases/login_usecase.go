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
}

func NewLoginUsecase(authRepository repositories.AuthRepository) *loginUsecase {
	return &loginUsecase{
		authRepository: authRepository,
	}
}

func (s *loginUsecase) Execute(ctx context.Context, loginRequest *dtos.LoginRequest) (*dtos.LoginResponse, error) {
	authUser, err := s.authRepository.GetUserByEmail(ctx, loginRequest.Email)
	if err != nil || security.ComparePasswords(authUser.Password, loginRequest.Email) {
		return &dtos.LoginResponse{}, errors.New("invalid credentials")
	}
	response := converters.ToAuthUserResponse(authUser)
	return &response, nil
}
