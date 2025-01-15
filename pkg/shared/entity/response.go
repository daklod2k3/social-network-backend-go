package entity

import "github.com/gin-gonic/gin"

type ResponseJson struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (res ResponseJson) WriteError(c *gin.Context) {
	if res.Status == 0 {
		res.Status = 500
	}

	c.AbortWithStatusJSON(res.Status, res)
}

func (res ResponseJson) WriteSuccess(c *gin.Context) {
	if res.Status == 0 {
		res.Status = 200
	}
	c.AbortWithStatusJSON(res.Status, res)
}
