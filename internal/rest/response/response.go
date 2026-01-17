package response

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/faiyaz032/goplate/internal/domain"
	"go.uber.org/zap"
)

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PaginatedData struct {
	Items      interface{} `json:"items"`
	TotalCount int         `json:"total_count"`
	Page       int         `json:"page"`
	Size       int         `json:"size"`
	HasNext    bool        `json:"has_next"`
}

func JSON(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func PaginatedJSON(w http.ResponseWriter, status int, message string, data PaginatedData) {
	JSON(w, status, message, data)
}

func HandleError(w http.ResponseWriter, log *zap.Logger, err error) {
	var appErr *domain.AppError
	statusCode := http.StatusInternalServerError
	message := "Internal server error"

	if errors.As(err, &appErr) {

		if appErr.Raw != nil {
			log.Error("API error", zap.String("message", appErr.Message), zap.Error(appErr.Raw))
		}

		switch {
		case errors.Is(appErr.Err, domain.ErrBadRequest):
			statusCode = http.StatusBadRequest
			message = appErr.Message
		case errors.Is(appErr.Err, domain.ErrUnprocessable):
			statusCode = http.StatusUnprocessableEntity
			message = appErr.Message
		case errors.Is(appErr.Err, domain.ErrNotFound):
			statusCode = http.StatusNotFound
			message = appErr.Message
		case errors.Is(appErr.Err, domain.ErrConflict):
			statusCode = http.StatusConflict
			message = appErr.Message
		default:

			statusCode = http.StatusInternalServerError
			message = "Internal server error"
		}
	} else {
		log.Error("unexpected error", zap.Error(err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(BaseResponse{
		Success: false,
		Message: message,
	})
}
