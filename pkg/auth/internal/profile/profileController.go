package profile

import (
	"auth/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shared/database"
	"shared/entity"
	authEntity "shared/entity/auth"
	"shared/global"
)

var (
	logger = global.Logger
)

type controller struct {
	service  *profileService
	profiles *entity.UserRepo
}

func NewController() *controller {

	db := database.New()

	return &controller{
		profiles: entity.NewRepo(db.GetSchema()),
	}
}

func (s *controller) CreateHdl(c *gin.Context) {
	var form Post
	if err := c.Bind(&form); err != nil {
		authEntity.ParseError(err, 400).WriteError(c)
		return
	}

	logger.Info(fmt.Sprintf("%+v", form))

	session := authUtils.GetSessionFromContext(c)

	createdUser, err := s.profiles.CreateUser(session.UserId, form.DisplayName, form.AvatarPath)
	if err != nil {
		authEntity.ParseError(err, 400).WriteError(c)
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}

func (s *controller) GetHdl(c *gin.Context) {
	user := authUtils.GetUserFromContext(c)
	if user == nil {
		c.AbortWithStatus(403)
		c.JSON(-1, gin.H{
			"message": "User not found",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}
