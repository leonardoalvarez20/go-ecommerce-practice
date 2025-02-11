package services

import (
	"context"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/converters"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/repositories"
)

type ProductService struct {
	repo *repositories.ProductsRepository
}

func CreateProductService(repo *repositories.ProductsRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s ProductService) Create(ctx context.Context, p *dtos.CreateProductRequest) (dtos.ProductResponse, error) {
	product := converters.ToDatabaseModel(p)
	err := s.repo.Create(ctx, &product)

	if err != nil {
		return dtos.ProductResponse{}, err
	}

	response := converters.ToProductResponse(&product)
	return response, nil
}

func (s ProductService) GetById(id string) (dtos.ProductResponse, error) {
	product, err := s.repo.GetById(id)

	if err != nil {
		return dtos.ProductResponse{}, err
	}
	response := converters.ToProductResponse(&product)

	return response, nil
}
