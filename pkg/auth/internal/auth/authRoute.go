package auth

import (
	"github.com/gin-gonic/gin"
)

func (ctl *Controller) ApplyRoute(r *gin.RouterGroup) {
	rAuth := r.Group("auth")

	rAuth.POST("/login", ctl.LoginHandler)

	rAuth.POST("/register", ctl.RegisterHandler)

	rAuth.POST("/session", ctl.GetSessionHandler)

	rAuth.GET("/health", ctl.Health)
}
