package profile

import (
	entity2 "auth/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"shared/database"
	"shared/entity"
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
	var user Post
	if err := c.ShouldBind(&user); err != nil {
		entity2.ParseError(err).WriteError(c)
		return
	}
	createdUser, err := s.profiles.CreateUser(&user.UserId, &user.DisplayName)
	if err != nil {
		entity2.ParseError(err).WriteError(c)
	}

	c.JSON(http.StatusCreated, createdUser)
}
