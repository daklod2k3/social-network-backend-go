package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"shared/entity"
	authEntity "shared/entity/auth"
	"shared/utils"
)

func GetUserFromContext(c *gin.Context) (*entity.User, error) {
	if c.GetHeader("x-profile") == "not created" {
		c.AbortWithStatusJSON(403, gin.H{"message": "User Not Found"})
		return nil, errors.New("User Not Found")
	}
	session, err := GetSessionFromContext(c)
	if err != nil {
		return nil, err
	}
	return session.User, nil
}

func GetSessionFromContext(c *gin.Context) (session *authEntity.AuthResponse, _ error) {
	sessionStr := c.GetString("session")
	if sessionStr == "" {
		c.AbortWithStatus(401)
		return nil, errors.New("Session Not Found")
	}
	err := utils.Deserialize(sessionStr, &session)
	if err != nil {
		return nil, err
	}
	//fmt.Println(session)
	//logger.Info(sessionStr)
	return session, nil
}

func AbortNoUserCreated(c *gin.Context) {
	c.AbortWithStatusJSON(403, gin.H{"message": "Please create User first"})
}
