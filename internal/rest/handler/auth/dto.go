package authhandler

import "github.com/faiyaz032/goplate/internal/domain"

type RegisterUserDTO struct {
	Username string `json:"username" validate:"required,min=3,max=32,alphanum"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

func (r *RegisterUserDTO) toDomain() *domain.User {
	return &domain.User{
		Username: r.Username,
		Email:    r.Email,
	}
}
