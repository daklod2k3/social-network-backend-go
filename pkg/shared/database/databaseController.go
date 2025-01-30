package database

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *service) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.Health())
}
