package repositories

import (
	"context"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/models"
)

type UsersRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetById(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, id string, user *models.User) error
}
