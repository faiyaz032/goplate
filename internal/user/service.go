package user

import (
	"go.uber.org/zap"
)


type service struct {
	userRepo UserRepo
	log      *zap.Logger
}

func NewService(userRepo UserRepo, log *zap.Logger) Service {
	return &service{
		userRepo,
		log,
	}
}
