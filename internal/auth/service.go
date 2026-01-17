package auth

import (
	"go.uber.org/zap"
)

type service struct {
	userSvc UserSvc
	log     *zap.Logger
}

func NewService(userSvc UserSvc, log *zap.Logger) Service {
	return &service{
		userSvc,
		log,
	}
}
