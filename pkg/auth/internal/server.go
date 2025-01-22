package internal

import (
	"auth/internal/auth"
	"auth/internal/global"
	"fmt"
	"net/http"
	"shared/interfaces"
	"time"
)

type Server struct {
	port int
	interfaces.AuthService
}

func NewServer() *http.Server {
	port := global.Config.Port
	NewServer := &Server{
		port:        port,
		AuthService: auth.NewService(),
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
