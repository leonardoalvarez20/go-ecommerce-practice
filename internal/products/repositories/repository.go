package repositories

import (
	"context"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/models"
)

type ProductsRepository interface {
	Create(ctx context.Context, product *models.Product) error
	GetById(ctx context.Context, id string) (*models.Product, error)
	GetAll(ctx context.Context) ([]models.Product, error)
}
