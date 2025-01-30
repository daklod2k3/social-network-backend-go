package internal

import (
	"core/internal/global"
	"fmt"
	"net/http"
	"shared/database"
	"shared/initialize"
	"shared/logger"
	authRpcClient "shared/rpc/client/auth"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port        int
	db          database.Service
	Logger      *logger.LoggerZap
	AuthService authRpcClient.AuthRpcService
}

func NewServer() *http.Server {
	// init global of this pkg
	global.InitGlobal()

	// init global of shared pkg
	initialize.InitGlobal(&initialize.Type{
		Config: global.Config,
		Logger: global.Logger,
	})

	authService := authRpcClient.NewClient()

	port := global.Config.Port
	NewServer := &Server{
		port:        port,
		Logger:      global.Logger,
		db:          database.New(),
		AuthService: authService,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
