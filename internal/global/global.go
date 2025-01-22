package global

import (
	"shared/config"
	"shared/logger"
)

var (
	Config *config.Configuration
	Logger *logger.LoggerZap
)

func InitGlobal() {
	Config = config.NewConfig()
	Logger = logger.NewLogger(Config.LogLevel)
}
