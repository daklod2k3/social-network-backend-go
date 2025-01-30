package database

import "github.com/gin-gonic/gin"

func (s *service) ApplyRoute(r *gin.RouterGroup) {
	r.GET("/health", s.HealthHandler)
}
