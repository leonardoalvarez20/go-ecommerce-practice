package dtos

type CreateUserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Phone     string           `json:"phone"`
	Password  string           `json:"password"`
	Addresses []AddressRequest `json:"addresses"`
}

type UserResponse struct {
	ID        string            `json:"id"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	FullName  string            `json:"full_name"`
	Email     string            `json:"email"`
	Phone     string            `json:"phone"`
	Addresses []AddressResponse `json:"addresses"`
}
