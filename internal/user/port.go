package user

import (
	"context"

	"github.com/faiyaz032/goplate/internal/domain"
	userHandler "github.com/faiyaz032/goplate/internal/rest/handler/user"
	"github.com/google/uuid"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(context.Context, *domain.User) (*domain.User, error)
	FindByID(context.Context, uuid.UUID) (*domain.User, error)
}
