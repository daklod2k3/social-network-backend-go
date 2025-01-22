package global

import (
	"shared/config"
	"shared/logger"
)

type Type struct {
	Config *config.Configuration
	Logger *logger.LoggerZap
}

var (
	Config *config.Configuration
	Logger *logger.LoggerZap
)

func InitGlobal(global *Type) {
	Config = global.Config
	Logger = global.Logger
}
