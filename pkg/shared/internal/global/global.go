package global

import (
	"shared/config"
	"shared/logger"
)

var (
	Config *config.Configuration
	Logger *logger.LoggerZap
)
