package authhandler

import (
	"context"

	"github.com/faiyaz032/goplate/internal/domain"
)

type Service interface {
	Register(context.Context, *domain.User) (*domain.User, error)
}
