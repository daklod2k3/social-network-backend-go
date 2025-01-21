package auth

import (
	"github.com/gin-gonic/gin"
)

func (ctl *Controller) ApplyRoute(r *gin.RouterGroup, authorized *gin.RouterGroup) {
	rAuth := r.Group("auth")
	rAuthWithAuthorized := authorized.Group("auth")

	rAuth.POST("/login", ctl.LoginHandler)

	rAuth.POST("/register", ctl.RegisterHandler)

	rAuthWithAuthorized.GET("/session", ctl.GetSessionHandler)

	rAuth.GET("/health", ctl.Health)
}
