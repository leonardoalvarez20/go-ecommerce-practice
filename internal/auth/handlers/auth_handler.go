package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/auth/services"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/config"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/shared/models"
)

type AuthHandler struct {
	service services.AuthService
	config  *config.Config
}

func NewAuthUserHandler(service services.AuthService, config *config.Config) *AuthHandler {
	return &AuthHandler{
		service: service,
		config:  config,
	}
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var authUserRequest dtos.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&authUserRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.Login(ctx, h.config, &authUserRequest)
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
