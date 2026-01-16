package authhandler

import (
	"encoding/json"
	"net/http"

	"github.com/faiyaz032/goplate/internal/domain"
	"github.com/faiyaz032/goplate/internal/rest/response"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	validate *validator.Validate
	svc      Service
}

func NewHandler(validate *validator.Validate, svc Service) *Handler {
	return &Handler{
		validate: validate,
		svc:      svc,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	dto := new(RegisterUserDTO)
	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		response.HandleError(w, &domain.AppError{
			Err:     domain.ErrBadRequest,
			Message: "Invalid request body",
			Raw:     err,
		})
		return
	}

	if err := h.validate.Struct(dto); err != nil {
		response.HandleError(w, &domain.AppError{
			Err:     domain.ErrUnprocessable,
			Message: err.Error(),
			Raw:     err,
		})
		return
	}

	user, err := h.svc.Register(r.Context(), dto.toDomain())
	if err != nil {
		response.HandleError(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, "User registered successfully", user)
}
