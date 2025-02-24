package usecases

import (
	"context"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/converters"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/repositories"
)

type GetUserByIdUsecase struct {
	userRepository repositories.UsersRepository
}

func NewGetUserByIdUsecase(userRepository repositories.UsersRepository) *GetUserByIdUsecase {
	return &GetUserByIdUsecase{userRepository: userRepository}
}
func (s *GetUserByIdUsecase) GetById(ctx context.Context, id string) (*dtos.UserResponse, error) {
	user, err := s.userRepository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	response := converters.ToUserResponse(user)

	return &response, nil
}
