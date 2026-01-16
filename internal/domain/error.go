package domain

import "errors"

var (
	ErrNotFound      = errors.New("resource not found")
	ErrConflict      = errors.New("resource already exists")
	ErrInternal      = errors.New("internal server error")
	ErrBadRequest    = errors.New("bad request")
	ErrUnprocessable = errors.New("validation failed")
)

type AppError struct {
	Err     error
	Message string
	Raw     error
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}
