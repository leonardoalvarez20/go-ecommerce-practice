package services

import (
	"context"
	"errors"

	"github.com/leonardoalvarez20/go-ecommerce-practice/common/security"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/converters"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/repositories"
)

type AuthService struct {
	authRepository *repositories.AuthRepository
}

func NewAuthService(authRepository *repositories.AuthRepository) *AuthService {
	return &AuthService{
		authRepository: authRepository,
	}
}

func (s *AuthService) Login(ctx context.Context, authUserRequest *dtos.AuthUserRequest) (dtos.AuthUserResponse, error) {
	authUser, err := s.authRepository.GetUserByEmail(ctx, authUserRequest.Email)
	if err != nil || security.ComparePasswords(authUser.Password, authUserRequest.Email) {
		return dtos.AuthUserResponse{}, errors.New("invalid credentials")
	}

	return converters.ToAuthUserResponse(&authUser), nil
}
