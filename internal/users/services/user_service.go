package services

import (
	"context"
	"time"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/security"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/converters"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/repositories"
)

type UserService struct {
	repo           *repositories.UserRepository
	passwordHasher security.PasswordHasher
}

func CreateUserServices(repo *repositories.UserRepository, passwordHasher security.PasswordHasher) *UserService {
	return &UserService{
		repo:           repo,
		passwordHasher: passwordHasher,
	}
}

func (s *UserService) Create(ctx context.Context, u *dtos.CreateUserRequest) (dtos.UserResponse, error) {
	user := converters.ToUserDatabaseModel(u)
	user.Password, _ = s.passwordHasher.HashPassword(u.Password)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err := s.repo.Create(ctx, &user)

	if err != nil {
		return dtos.UserResponse{}, err
	}

	response := converters.ToUserResponse(&user)
	return response, nil
}

func (s *UserService) GetById(ctx context.Context, id string) (dtos.UserResponse, error) {
	user, err := s.repo.GetById(ctx, id)

	if err != nil {
		return dtos.UserResponse{}, err
	}
	response := converters.ToUserResponse(&user)

	return response, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id string, u *dtos.UpdateUserRequest) (dtos.UserResponse, error) {
	user := converters.ToUpdateUserDatabaseModel(u)
	user.UpdatedAt = time.Now()

	err := s.repo.UpdateUser(ctx, id, &user)

	if err != nil {
		return dtos.UserResponse{}, err
	}

	response := converters.ToUserResponse(&user)
	return response, nil
}
