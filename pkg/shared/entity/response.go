package entity

import "github.com/gin-gonic/gin"

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func WriteError(c *gin.Context, status int, message string) {
	if status == 0 {
		status = 500
	}

	c.AbortWithStatusJSON(status, Response{
		status,
		message,
	})
}

func WriteSuccess(c *gin.Context, status int, message string) {
	if status == 0 {
		status = 200
	}
	c.AbortWithStatusJSON(status, Response{
		status,
		message,
	})
}
