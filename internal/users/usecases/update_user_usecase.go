package usecases

import (
	"context"
	"time"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/converters"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/repositories"
)

type UpdateUserUsecase struct {
	userRepository repositories.UsersRepository
}

func NewUpdateUserUsecase(userRepository repositories.UsersRepository) *UpdateUserUsecase {
	return &UpdateUserUsecase{userRepository: userRepository}
}
func (s *UpdateUserUsecase) Execute(ctx context.Context, id string, u *dtos.UpdateUserRequest) (*dtos.UserResponse, error) {
	user := converters.ToUpdateUserDatabaseModel(u)
	user.UpdatedAt = time.Now()

	err := s.userRepository.UpdateUser(ctx, id, &user)

	if err != nil {
		return nil, err
	}

	response := converters.ToUserResponse(&user)
	return &response, nil
}
