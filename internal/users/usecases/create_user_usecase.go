package usecases

import (
	"context"
	"time"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/security"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/converters"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/repositories"
)

type CreateUserUsecase struct {
	userRepository repositories.UsersRepository
	passwordHasher security.PasswordHasher
}

func NewCreateUserUsecase(userRepository repositories.UsersRepository, passwordHasher security.PasswordHasher) *CreateUserUsecase {
	return &CreateUserUsecase{
		userRepository: userRepository,
		passwordHasher: passwordHasher,
	}
}

func (s *CreateUserUsecase) Execute(ctx context.Context, u *dtos.CreateUserRequest) (*dtos.UserResponse, error) {
	user := converters.ToUserDatabaseModel(u)
	user.Password, _ = s.passwordHasher.HashPassword(u.Password)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err := s.userRepository.Create(ctx, &user)

	if err != nil {
		return nil, err
	}

	response := converters.ToUserResponse(&user)
	return &response, nil
}
