package auth

import (
	"context"

	"github.com/faiyaz032/goplate/internal/domain"
	authhandler "github.com/faiyaz032/goplate/internal/rest/handler/auth"

	"github.com/google/uuid"
)

type Service interface {
	authhandler.Service
}

type UserSvc interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	Create(ctx context.Context, record *domain.User) (*domain.User, error)
}
