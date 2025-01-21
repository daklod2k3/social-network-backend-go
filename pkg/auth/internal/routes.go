package internal

import (
	"auth/internal/auth"
	"auth/internal/profile"
	"github.com/gin-gonic/gin"
	"net/http"
	"shared/middlewares"
	"shared/route"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	defaultGroup := route.DefaultRouteConfig(r)

	authorizedGroup := defaultGroup.Group("/")

	authorizedGroup.Use(middlewares.AuthMiddleware(s.AuthService))

	auth.NewController().ApplyRoute(defaultGroup, authorizedGroup)
	profile.NewController().ApplyRoute(authorizedGroup, authorizedGroup)

	return r
}
