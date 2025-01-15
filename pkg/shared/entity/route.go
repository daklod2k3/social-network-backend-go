package entity

import "github.com/gin-gonic/gin"

type Route interface {
	ApplyRoute(group *gin.RouterGroup)
}
