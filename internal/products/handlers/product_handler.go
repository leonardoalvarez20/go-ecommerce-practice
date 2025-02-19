package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/dtos"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/products/services"
	"github.com/leonardoalvarez20/go-ecommerce-practice/internal/shared/models"
)

type ProductHandler struct {
	service *services.ProductService
}

func CreateProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var createProductRequest dtos.CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&createProductRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.Create(ctx, &createProductRequest)
	w.Header().Set("Content-Type", "application/json")
	var apiResponse models.ApiResponse
	statusCode := http.StatusCreated
	if err != nil {
		statusCode = http.StatusInternalServerError
		apiResponse = models.NewErrorResponse(err.Error(), http.StatusInternalServerError)
	} else {
		apiResponse = models.NewSuccessResponse(response)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(apiResponse)
}

func (h ProductHandler) GetById(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(apiResponse)
}

func (h ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response, err := h.service.GetAll(ctx)

	var apiResponse models.ApiResponse
	statusCode := http.StatusOK
	if err != nil {
		statusCode = http.StatusInternalServerError
		apiResponse = models.NewErrorResponse(err.Error(), http.StatusInternalServerError)
	} else {
		apiResponse = models.NewSuccessResponse(response)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(apiResponse)
}
