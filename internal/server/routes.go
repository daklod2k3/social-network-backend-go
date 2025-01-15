package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shared/database"
	"shared/route"
)

func (s *Server) RegisterRoutes() http.Handler {

	r := gin.Default()

	defaultGroup := route.DefaultRouteConfig(r)

	database.ApplyRoute(defaultGroup)

	return r
}
