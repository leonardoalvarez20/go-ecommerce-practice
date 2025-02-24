package users

import (
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/security"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/repositories"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/services"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/usecases"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersContainer struct {
	UsersRepositories repositories.UsersRepository
	UserService       *services.UserService
}

func NewUsersContainer(db *mongo.Database, passwordHasher security.PasswordHasher) *UsersContainer {
	usersRepositories := repositories.NewUserRepository(db)
	usersServiceUsecases := services.UserServiceUsecases{
		CreateUserUsecase:  usecases.NewCreateUserUsecase(usersRepositories, passwordHasher),
		GetUserByIdUsecase: usecases.NewGetUserByIdUsecase(usersRepositories),
		UpdateUserUsecase:  usecases.NewUpdateUserUsecase(usersRepositories),
	}
	userService := services.CreateUserServices(usersRepositories, usersServiceUsecases)

	return &UsersContainer{
		UsersRepositories: usersRepositories,
		UserService:       userService,
	}
}
