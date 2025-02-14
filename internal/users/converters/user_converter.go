package converters

import (
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/models"
)

func ToUserDatabaseModel(dto *dtos.CreateUserRequest) models.User {
	return models.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  dto.Password,
	}
}

func ToUserResponse(model *models.User) dtos.UserResponse {
	var addresses = []dtos.AddressResponse{}
	for _, address := range model.Addresses {
		addresses = append(addresses, ToAddressResponse(&address))
	}

	return dtos.UserResponse{
		ID:        model.ID.Hex(),
		FirstName: model.FirstName,
		LastName:  model.LastName,
		FullName:  model.FirstName + " " + model.LastName,
		Email:     model.Email,
		Phone:     model.Phone,
		Addresses: addresses,
	}
}

func ToUpdateUserDatabaseModel(dto *dtos.UpdateUserRequest) models.User {
	var addresses = []models.Address{}
	for _, address := range dto.Addresses {
		addresses = append(addresses, ToAddressDatabaseModel(&address))
	}
	return models.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Phone:     dto.Phone,
		Password:  dto.Password,
		Addresses: addresses,
	}
}
