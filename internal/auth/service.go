package auth

import (
	"context"

	"github.com/faiyaz032/goplate/internal/domain"
)

type service struct {
	userSvc UserSvc
}

func NewService(userSvc UserSvc) Service {
	return &service{
		userSvc,
	}
}

func (s *service) Register(ctx context.Context, data *domain.User) (*domain.User, error) {
	createdUser, err := s.userSvc.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
