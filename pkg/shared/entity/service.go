package entity

import (
	"shared/database"
)

type Service struct {
	Db database.Service
}

func NewService(connectString string) *Service {
	return &Service{
		database.New(connectString),
	}
}
