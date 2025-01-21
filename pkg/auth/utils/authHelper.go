package authUtils

import (
	"github.com/gin-gonic/gin"
	"shared/entity"
	authEntity "shared/entity/auth"
	logger2 "shared/logger"
	"shared/utils"
)

var logger = logger2.GetLogger()

func GetUserFromContext(c *gin.Context) *entity.User {
	if c.GetHeader("x-profile") == "not created" {
		c.AbortWithStatusJSON(403, gin.H{"message": "User Not Found"})
		return nil
	}
	return GetSessionFromContext(c).User
}

func GetSessionFromContext(c *gin.Context) (session *authEntity.AuthResponse) {
	sessionStr := c.GetString("session")
	if sessionStr == "" {
		c.AbortWithStatus(401)
		return nil
	}
	//logger.Info(sessionStr)
	utils.Deserialize(sessionStr, &session)
	//logger.Info(sessionStr)
	return session
}
