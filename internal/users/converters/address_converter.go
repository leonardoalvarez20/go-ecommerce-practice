package converters

import (
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/models"
)

func ToAddressDatabaseModel(dto *dtos.AddressRequest) models.Address {
	return models.Address{
		Street:  dto.Street,
		Number:  dto.Number,
		City:    dto.City,
		State:   dto.State,
		ZipCode: dto.ZipCode,
	}
}

func ToAddressResponse(model *models.Address) dtos.AddressResponse {
	return dtos.AddressResponse{
		ID:      model.ID.Hex(),
		Street:  model.Street,
		Number:  model.Number,
		City:    model.City,
		State:   model.State,
		ZipCode: model.ZipCode,
	}
}
