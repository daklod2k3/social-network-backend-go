package internal

import (
	"auth/internal/auth"
	"auth/internal/profile"
	"github.com/gin-gonic/gin"
	"net/http"
	"shared/route"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	defaultGroup := route.DefaultRouteConfig(r)

	auth.NewController().ApplyRoute(defaultGroup)
	profile.NewController().ApplyRoute(defaultGroup)

	return r
}
