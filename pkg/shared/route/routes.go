package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func DefaultRouteConfig(r *gin.Engine) *gin.RouterGroup {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	return r.Group("api/v1")
}
