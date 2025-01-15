package internal

import (
	"fmt"
	"net/http"
	"shared/config"
	"time"
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	port := config.GetConfig().Auth.Port
	NewServer := &Server{
		port: port,
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
