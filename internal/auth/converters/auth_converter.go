package converters

import (
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/models"
)

func ToAuthUserResponse(u *models.AuthUser) dtos.LoginResponse {
	return dtos.LoginResponse{
		ID:       u.ID.Hex(),
		Email:    u.Email,
		FullName: u.FirsName + " " + u.LastName,
	}
}
