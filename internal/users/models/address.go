package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	ID        primitive.ObjectID `bson:"_id"`
	Street    string             `bson:"street,omitempty"`
	Number    string             `bson:"number,omitempty"`
	ZipCode   string             `bson:"zip_code,omitempty"`
	City      string             `bson:"city,omitempty"`
	State     string             `bson:"state,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
