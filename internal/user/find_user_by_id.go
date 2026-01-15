package user

import (
	"context"

	"github.com/faiyaz032/goplate/internal/domain"
	"github.com/google/uuid"
)

func (s *service) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
