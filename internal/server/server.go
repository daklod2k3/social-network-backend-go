package server

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"shared/database"
	"shared/logger"

	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port   int
	db     database.Service
	Logger *zap.Logger
}

func NewServer() *http.Server {
	port := viper.GetInt("port")
	NewServer := &Server{
		port:   port,
		Logger: logger.GetLogger(),
		db:     database.New(),
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
