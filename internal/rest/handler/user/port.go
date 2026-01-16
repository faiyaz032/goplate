package userhandler

import (
	"context"

	"github.com/faiyaz032/goplate/internal/domain"
	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, record *domain.User) (*domain.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
}
