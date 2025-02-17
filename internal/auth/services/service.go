package services

import (
	"context"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/dtos"
)

type AuthService interface {
	Login(ctx context.Context, loginRequest *dtos.LoginRequest) (*dtos.LoginResponse, error)
}
