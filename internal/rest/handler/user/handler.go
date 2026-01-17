package userhandler

import (
	"encoding/json"
	"net/http"

	"github.com/faiyaz032/goplate/internal/domain"
	"github.com/faiyaz032/goplate/internal/rest/response"
	"go.uber.org/zap"
)

type Handler struct {
	svc Service
	log *zap.Logger
}

func NewHandler(svc Service, log *zap.Logger) *Handler {
	return &Handler{
		svc,
		log,
	}
}

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.HandleError(w, h.log, &domain.AppError{
			Err:     domain.ErrBadRequest,
			Message: "Invalid request body",
			Raw:     err,
		})
		return
	}

	createdUser, err := h.svc.Create(r.Context(), &user)
	if err != nil {
		response.HandleError(w, h.log, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}
