package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/models"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/usecases"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/security"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mockAuthRepository struct {
	mock.Mock
}

func (m *mockAuthRepository) GetUserByEmail(ctx context.Context, email string) (*models.AuthUser, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.AuthUser), args.Error(1)
}

func TestLoginUsecase_Execute(t *testing.T) {
	ctx := context.Background()

	t.Run("should return response when credentials are valid", func(t *testing.T) {
		mockRepo := new(mockAuthRepository)
		hasher := &security.BcryptHasher{}

		userID, _ := primitive.ObjectIDFromHex("67b19fc4ea6f7933149ecc59")
		hashedPassword, _ := hasher.HashPassword("leo")
		authModel := &models.AuthUser{
			ID:       userID,
			Email:    "test@example.com",
			FirsName: "Leonardo",
			LastName: "Alvarez",
			Password: hashedPassword,
		}
		mockRepo.On("GetUserByEmail", ctx, "test@example.com").Return(authModel, nil)

		usecase := usecases.NewLoginUsecase(mockRepo, hasher)

		loginRequest := &dtos.LoginRequest{Email: "test@example.com", Password: "leo"}

		resp, err := usecase.Execute(ctx, loginRequest)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "test@example.com", resp.Email)
		assert.Equal(t, "Leonardo Alvarez", resp.FullName)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when user is not found", func(t *testing.T) {
		mockRepo := new(mockAuthRepository)
		hasher := &security.BcryptHasher{}

		mockRepo.On("GetUserByEmail", ctx, "notfound@example.com").Return(nil, errors.New("user not found"))

		usecase := usecases.NewLoginUsecase(mockRepo, hasher)

		loginRequest := &dtos.LoginRequest{Email: "notfound@example.com", Password: "password"}

		resp, err := usecase.Execute(ctx, loginRequest)

		assert.Error(t, err)
		assert.Equal(t, "invalid credentials", err.Error())
		assert.Nil(t, resp)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when password is incorrect", func(t *testing.T) {
		mockRepo := new(mockAuthRepository)
		hasher := &security.BcryptHasher{}
		hashedPassword, _ := hasher.HashPassword("hashedpassword")

		mockRepo.On("GetUserByEmail", ctx, "test@example.com").Return(&models.AuthUser{
			Email:    "test@example.com",
			Password: hashedPassword,
		}, nil)

		usecase := usecases.NewLoginUsecase(mockRepo, hasher)

		loginRequest := &dtos.LoginRequest{Email: "test@example.com", Password: "wrongpassword"}

		resp, err := usecase.Execute(ctx, loginRequest)

		assert.Error(t, err)
		assert.Equal(t, "invalid credentials", err.Error())
		assert.Nil(t, resp)

		mockRepo.AssertExpectations(t)
	})
}
