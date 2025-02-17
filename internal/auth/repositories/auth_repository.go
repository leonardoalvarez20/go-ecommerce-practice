package repositories

import (
	"context"
	"errors"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type authRepository struct {
	authUser *mongo.Collection
}

func NewAuthRepository(db *mongo.Database) AuthRepository {
	return &authRepository{
		authUser: db.Collection("users"),
	}
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (*models.AuthUser, error) {
	var user models.AuthUser

	filter := bson.M{"email": email}
	err := r.authUser.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.AuthUser{}, errors.New("user not found")
		}
		return &models.AuthUser{}, err
	}

	return &user, nil
}
