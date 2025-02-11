package models

// ApiResponse es la estructura común para todas las respuestas JSON
type ApiResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Code    int         `json:"code,omitempty"`
}

// NewSuccessResponse crea una respuesta de éxito
func NewSuccessResponse(data interface{}) ApiResponse {
	return ApiResponse{
		Status: "success",
		Data:   data,
	}
}

// NewErrorResponse crea una respuesta de error
func NewErrorResponse(message string, code int) ApiResponse {
	return ApiResponse{
		Status:  "error",
		Message: message,
		Code:    code,
	}
}
