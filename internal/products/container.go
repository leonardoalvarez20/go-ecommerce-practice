package products

import (
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/repositories"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/services"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/usecases"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductsContainer struct {
	ProductsRepositories repositories.ProductsRepository
	ProductService       *services.ProductService
}

func NewProductsContainer(db *mongo.Database) *ProductsContainer {
	productsRepositories := repositories.NewProductRepository(db)
	productsServiceUsecases := services.ProductServiceUsecases{
		CreateProductUsecase:  usecases.NewCreateProductUsecase(productsRepositories),
		GetProductByIdUsecase: usecases.NewGetProductByIdUsecase(productsRepositories),
		GetAllProductUsecase:  usecases.NewGetAllProductUsecase(productsRepositories),
	}
	productService := services.NewProductService(productsServiceUsecases)

	return &ProductsContainer{
		ProductsRepositories: productsRepositories,
		ProductService:       productService,
	}
}
