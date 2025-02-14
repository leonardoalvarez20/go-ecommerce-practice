package dtos

type AddressRequest struct {
	Street  string `json:"street" validate:"required"`
	Number  string `json:"number" validate:"required"`
	ZipCode string `json:"zip_code" validate:"required"`
	City    string `json:"city" validate:"required"`
	State   string `json:"state" validate:"required"`
}

type AddressResponse struct {
	ID      string `json:"id"`
	Street  string `json:"street"`
	Number  string `json:"number"`
	ZipCode string `json:"zip_code"`
	City    string `json:"city"`
	State   string `json:"state"`
}
