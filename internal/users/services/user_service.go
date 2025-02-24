package services

import (
	"context"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/repositories"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/usecases"
)

type UserServiceUsecases struct {
	CreateUserUsecase  *usecases.CreateUserUsecase
	GetUserByIdUsecase *usecases.GetUserByIdUsecase
	UpdateUserUsecase  *usecases.UpdateUserUsecase
}

type UserService struct {
	repo     repositories.UsersRepository
	usecases UserServiceUsecases
}

func CreateUserServices(repo repositories.UsersRepository, usecases UserServiceUsecases) *UserService {
	return &UserService{
		repo:     repo,
		usecases: usecases,
	}
}

func (s *UserService) Create(ctx context.Context, u *dtos.CreateUserRequest) (*dtos.UserResponse, error) {
	return s.usecases.CreateUserUsecase.Execute(ctx, u)
}

func (s *UserService) GetById(ctx context.Context, id string) (*dtos.UserResponse, error) {
	return s.usecases.GetUserByIdUsecase.GetById(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, id string, u *dtos.UpdateUserRequest) (*dtos.UserResponse, error) {
	return s.usecases.UpdateUserUsecase.Execute(ctx, id, u)
}
