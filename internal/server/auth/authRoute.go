package auth

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoute(r *gin.RouterGroup) {
	rAuth := r.Group("auth")
	controller := NewController()

	rAuth.POST("/login", controller.LoginHandler)

	rAuth.POST("/register", controller.RegisterHandler)

	rAuth.GET("/session", controller.GetSessionHandler)
}
