package authhandler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	svc      Service
	validate *validator.Validate
}

func NewHandler(svc Service, validate *validator.Validate) *Handler {
	return &Handler{
		svc:      svc,
		validate: validate,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	dto := new(RegisterUserDTO)
	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(dto); err != nil {
		http.Error(w, "validation failed: "+err.Error(), http.StatusUnprocessableEntity)
		return
	}

	createdUser, err := h.svc.Register(r.Context(), dto.toDomain())
	if err != nil {
		http.Error(w, "failed to register user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}
