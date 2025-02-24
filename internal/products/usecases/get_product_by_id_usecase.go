package usecases

import (
	"context"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/converters"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/repositories"
)

type GetProductByIdUsecase struct {
	productsRepository repositories.ProductsRepository
}

func NewGetProductByIdUsecase(productsRepository repositories.ProductsRepository) *GetProductByIdUsecase {
	return &GetProductByIdUsecase{productsRepository: productsRepository}
}

func (s *GetProductByIdUsecase) Execute(ctx context.Context, id string) (*dtos.ProductResponse, error) {
	product, err := s.productsRepository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}
	response := converters.ToProductResponse(product)

	return &response, nil
}
