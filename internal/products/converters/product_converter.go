package converters

import (
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/models"
)

func ToDatabaseModel(dto *dtos.CreateProductRequest) models.Product {
	return models.Product{
		Name:        dto.Name,
		Price:       dto.Price,
		Description: dto.Description,
		Stock:       dto.Stock,
	}
}

func ToProductResponse(model *models.Product) dtos.ProductResponse {
	return dtos.ProductResponse{
		ID:          model.ID.Hex(),
		Name:        model.Name,
		Price:       model.Price,
		Description: model.Description,
		Stock:       model.Stock,
	}
}
