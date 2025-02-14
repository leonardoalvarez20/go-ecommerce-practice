package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthUser struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	FirsName string             `bson:"first_name"`
	LastName string             `bson:"last_name"`
	Password string             `bson:"password"`
}
