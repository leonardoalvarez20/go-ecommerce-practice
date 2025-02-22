package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/shared/handlers"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/shared/models"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/users/services"
)

type UserHandler struct {
	service *services.UserService
	handlers.BaseHandler
}

func CreateUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var createUserRequest dtos.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&createUserRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.Create(ctx, &createUserRequest)
	w.Header().Set("Content-Type", "application/json")
	var apiResponse models.ApiResponse
	statusCode := http.StatusCreated
	if err != nil {
		statusCode = http.StatusInternalServerError
		apiResponse = models.NewErrorResponse(err.Error(), http.StatusInternalServerError)
	} else {
		apiResponse = models.NewSuccessResponse(response)
	}

	h.WriteJSONResponse(w, statusCode, apiResponse)
}

func (h UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	response, err := h.service.GetById(ctx, id)

	var apiResponse models.ApiResponse
	statusCode := http.StatusOK
	if err != nil {
		statusCode = http.StatusNotFound
		apiResponse = models.NewErrorResponse(err.Error(), http.StatusNotFound)
	} else {
		apiResponse = models.NewSuccessResponse(response)
	}

	h.WriteJSONResponse(w, statusCode, apiResponse)
}

func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	var updateUserRequest dtos.UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&updateUserRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.UpdateUser(ctx, id, &updateUserRequest)
	w.Header().Set("Content-Type", "application/json")
	var apiResponse models.ApiResponse
	statusCode := http.StatusOK
	if err != nil {
		statusCode = http.StatusInternalServerError
		apiResponse = models.NewErrorResponse(err.Error(), http.StatusInternalServerError)
	} else {
		apiResponse = models.NewSuccessResponse(response)
	}

	h.WriteJSONResponse(w, statusCode, apiResponse)
}
