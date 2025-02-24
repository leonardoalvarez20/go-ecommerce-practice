package services

import (
	"context"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/usecases"
)

type ProductServiceUsecases struct {
	CreateProductUsecase  *usecases.CreateProductUsecase
	GetProductByIdUsecase *usecases.GetProductByIdUsecase
	GetAllProductUsecase  *usecases.GetAllProductUsecase
}

type ProductService struct {
	productServiceUsecases ProductServiceUsecases
}

func NewProductService(productServiceUsecase ProductServiceUsecases) *ProductService {
	return &ProductService{productServiceUsecases: productServiceUsecase}
}

func (s ProductService) Create(ctx context.Context, p *dtos.CreateProductRequest) (*dtos.ProductResponse, error) {
	return s.productServiceUsecases.CreateProductUsecase.Execute(ctx, p)
}

func (s ProductService) GetById(ctx context.Context, id string) (*dtos.ProductResponse, error) {
	return s.productServiceUsecases.GetProductByIdUsecase.Execute(ctx, id)
}

func (s ProductService) GetAll(ctx context.Context) ([]dtos.ProductResponse, error) {
	return s.productServiceUsecases.GetAllProductUsecase.Execute(ctx)
}
