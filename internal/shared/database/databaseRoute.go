package database

import "github.com/gin-gonic/gin"

func ApplyRoute(r *gin.RouterGroup) {
	r.GET("/health", HealthHandler)
}
