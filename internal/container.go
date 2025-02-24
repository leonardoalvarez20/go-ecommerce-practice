package internal

import (
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users"
)

type Container struct {
	Auth     *auth.AuthContainer
	Products *products.ProductsContainer
	Users    *users.UsersContainer
}

func NewContainer(
	authContainer *auth.AuthContainer,
	productsContainer *products.ProductsContainer,
	usersContainer *users.UsersContainer,
) *Container {
	return &Container{
		Auth:     authContainer,
		Products: productsContainer,
		Users:    usersContainer,
	}
}
