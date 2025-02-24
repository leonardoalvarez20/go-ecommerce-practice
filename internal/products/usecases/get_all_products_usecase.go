package usecases

import (
	"context"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/converters"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/repositories"
)

type GetAllProductUsecase struct {
	productsRepository repositories.ProductsRepository
}

func NewGetAllProductUsecase(productsRepository repositories.ProductsRepository) *GetAllProductUsecase {
	return &GetAllProductUsecase{productsRepository: productsRepository}
}

func (s *GetAllProductUsecase) Execute(ctx context.Context) ([]dtos.ProductResponse, error) {
	products, err := s.productsRepository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	var response []dtos.ProductResponse
	for _, p := range products {
		response = append(response, converters.ToProductResponse(&p))
	}

	return response, nil
}
