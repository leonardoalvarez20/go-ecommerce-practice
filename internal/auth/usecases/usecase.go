package usecases

import (
	"context"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/dtos"
)

type Usecase interface {
	Execute(ctx context.Context, loginRequest *dtos.LoginRequest) (*dtos.LoginResponse, error)
}
