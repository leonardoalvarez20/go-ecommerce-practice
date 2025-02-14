package services

import (
	"context"
	"time"

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
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	err := s.repo.Create(ctx, &product)

	if err != nil {
		return dtos.ProductResponse{}, err
	}

	response := converters.ToProductResponse(&product)
	return response, nil
}

func (s ProductService) GetById(ctx context.Context, id string) (dtos.ProductResponse, error) {
	product, err := s.repo.GetById(ctx, id)

	if err != nil {
		return dtos.ProductResponse{}, err
	}
	response := converters.ToProductResponse(&product)

	return response, nil
}

func (s ProductService) GetAll(ctx context.Context) ([]dtos.ProductResponse, error) {
	products, err := s.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	var response []dtos.ProductResponse
	for _, p := range products {
		response = append(response, converters.ToProductResponse(&p))
	}

	return response, nil
}
