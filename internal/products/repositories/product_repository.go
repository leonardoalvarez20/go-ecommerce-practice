package repositories

import (
	"context"
	"time"

	"errors"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductsRepository struct {
	products *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductsRepository {
	return &ProductsRepository{
		products: db.Collection("products"),
	}
}

func (r *ProductsRepository) Create(ctx context.Context, product *models.Product) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := r.products.InsertOne(ctx, product)
	if err != nil {
		return err
	}

	product.ID = result.InsertedID.(primitive.ObjectID)

	return nil
}

func (r *ProductsRepository) GetById(ctx context.Context, id string) (models.Product, error) {
	var product models.Product

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Product{}, errors.New("invalid id format")
	}

	filter := bson.M{"_id": objectID}
	err = r.products.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Product{}, errors.New("product not found")
		}
		return models.Product{}, err
	}

	return product, nil
}

func (r *ProductsRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	filter := bson.M{}

	cursor, err := r.products.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
