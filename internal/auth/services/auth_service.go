package services

import (
	"context"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/usecases"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/config"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/security"
)

type authService struct {
	login_usecase usecases.Usecase
}

func NewAuthService(login_usecase usecases.Usecase) *authService {
	return &authService{
		login_usecase: login_usecase,
	}
}

func (s *authService) Login(ctx context.Context, config *config.Config, loginRequest *dtos.LoginRequest) (*dtos.LoginResponse, error) {
	response, err := s.login_usecase.Execute(ctx, loginRequest)
	if err != nil {
		return nil, err
	}

	jwtToken, err := security.GenerateJWT(config, response.ID)
	if err != nil {
		return nil, err
	}

	response.Token = jwtToken
	return response, nil
}
