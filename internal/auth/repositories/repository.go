package repositories

import (
	"context"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/models"
)

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*models.AuthUser, error)
}
