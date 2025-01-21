package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shared/database"
	"shared/middlewares"
	"shared/route"
)

func (s *Server) RegisterRoutes() http.Handler {

	r := gin.Default()

	defaultGroup := route.DefaultRouteConfig(r)

	authorizedGroup := defaultGroup.Group("/")

	authorizedGroup.Use(middlewares.AuthMiddleware(s.AuthService))
	{
		database.ApplyRoute(authorizedGroup)
	}

	return r
}
