package entity

import (
	"go.uber.org/zap"
	"shared/database"
	logger2 "shared/logger"
)

type Service struct {
	Logger *zap.Logger
	Db     database.Service
}

func NewService() *Service {
	return &Service{
		logger2.GetLogger(),
		database.New(),
	}
}
