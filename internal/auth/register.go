package auth

import (
	"context"

	"github.com/faiyaz032/goplate/internal/domain"
	"go.uber.org/zap"
)

func (s *service) Register(ctx context.Context, data *domain.User) (*domain.User, error) {
	createdUser, err := s.userSvc.Create(ctx, data)
	if err != nil {
		s.log.Error("failed to register user", zap.Error(err))
		return nil, err
	}
	return createdUser, nil
}
