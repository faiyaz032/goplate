package repository

import (
	"context"

	"github.com/faiyaz032/goplate/internal/domain"
	db "github.com/faiyaz032/goplate/internal/infrastructure/db/sqlc"
	"github.com/faiyaz032/goplate/internal/user"
	"github.com/google/uuid"
)

type UserRepository interface {
	user.UserRepo
}

type repository struct {
	queries *db.Queries
}

func NewUserRepository(queries *db.Queries) UserRepository {
	return &repository{
		queries: queries,
	}
}

func (r *repository) Create(ctx context.Context, record *domain.User) (*domain.User, error) {
	createdRecord, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Username: record.Username,
		Email:    record.Email,
		Password: record.Password,
	})

	if err != nil {
		return nil, err
	}

	return toDomain(createdRecord), nil
}

func (r *repository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	record, err := r.queries.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return toDomain(record), nil
}

func toDomain(record db.User) *domain.User {
	return &domain.User{
		ID:        record.ID,
		Username:  record.Username,
		Email:     record.Email,
		Password:  record.Password,
		CreatedAt: timeFromTimestamptz(record.CreatedAt),
	}
}
