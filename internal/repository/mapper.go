package repository

import (
	"errors"
	"fmt"

	"github.com/faiyaz032/goplate/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func MapDBError(err error, contextMsg string) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return &domain.AppError{
			Err:     domain.ErrNotFound,
			Message: fmt.Sprintf("%s not found", contextMsg),
			Raw:     err,
		}
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return &domain.AppError{
				Err:     domain.ErrConflict,
				Message: fmt.Sprintf("%s already exists", contextMsg),
				Raw:     err,
			}
		case "23503":
			return &domain.AppError{
				Err:     domain.ErrBadRequest,
				Message: fmt.Sprintf("invalid reference for %s", contextMsg),
				Raw:     err,
			}
		}
	}

	return &domain.AppError{
		Err:     domain.ErrInternal,
		Message: "an internal database error occurred",
		Raw:     err,
	}
}
