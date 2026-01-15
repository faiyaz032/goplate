package user

import (
	"encoding/json"
	"net/http"

	"github.com/faiyaz032/goplate/internal/domain"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{
		svc,
	}
}

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	createdUser, err := h.svc.Create(r.Context(), &user)
	if err != nil {
		http.Error(w, "failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}
