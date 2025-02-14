package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/leonardoalvarez20/go-ecommerce-practice/common/models"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/services"
)

type AuthHandler struct {
	service *services.AuthService
}

func CreateAuthUserHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var authUserRequest dtos.AuthUserRequest
	err := json.NewDecoder(r.Body).Decode(&authUserRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.Login(ctx, &authUserRequest)
	w.Header().Set("Content-Type", "application/json")
	var apiResponse models.ApiResponse
	statusCode := http.StatusCreated
	if err != nil {
		statusCode = http.StatusUnauthorized
		apiResponse = models.NewErrorResponse(err.Error(), http.StatusUnauthorized)
	} else {
		apiResponse = models.NewSuccessResponse(response)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(apiResponse)
}
