package usecases

import (
	"context"
	"errors"
	"time"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/converters"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/repositories"
)

type CreateProductUsecase struct {
	productsRepository repositories.ProductsRepository
}

func NewCreateProductUsecase(productsRepository repositories.ProductsRepository) *CreateProductUsecase {
	return &CreateProductUsecase{productsRepository: productsRepository}
}

func (s *CreateProductUsecase) Execute(ctx context.Context, createProductRequest *dtos.CreateProductRequest) (*dtos.ProductResponse, error) {
	product := converters.ToDatabaseModel(createProductRequest)
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	err := s.productsRepository.Create(ctx, &product)
	if err != nil {
		return nil, errors.New("error creating product")
	}

	response := converters.ToProductResponse(&product)
	return &response, nil
}
