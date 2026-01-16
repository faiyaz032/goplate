package user

import (
	"context"

	"github.com/faiyaz032/goplate/internal/domain"
)

func (s *service) Create(ctx context.Context, record *domain.User) (*domain.User, error) {
	createdUser, err := s.userRepo.Create(ctx, record)
	if err != nil {
		return nil, err
	}
	return createdUser, err
}
