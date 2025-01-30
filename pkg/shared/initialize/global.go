package initialize

import (
	"shared/config"
	"shared/internal/global"
	"shared/logger"
)

type Type struct {
	Config *config.Configuration
	Logger *logger.LoggerZap
}

func InitGlobal(g *Type) {
	global.Config = g.Config
	global.Logger = g.Logger
}
