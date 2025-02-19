package services

import (
	"context"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/config"
)

type AuthService interface {
	Login(ctx context.Context, config *config.Config, loginRequest *dtos.LoginRequest) (*dtos.LoginResponse, error)
}
