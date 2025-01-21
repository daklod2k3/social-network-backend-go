package auth

import (
	"auth/entity"
	authUtils "auth/utils"
	"github.com/gin-gonic/gin"
	authEntity "shared/entity/auth"
	"shared/interfaces"
)

//type Controller struct {
//	reversePath string
//}
//
//func (a *Controller) LoginHandler(c *gin.Context) {
//	reverseProxyWithSupabase(c, a.reversePath+"/token?grant_type=password", "POST")
//}
//
//func (a *Controller) RegisterHandler(c *gin.Context) {
//	reverseProxyWithSupabase(c, a.reversePath+"/signup", "POST")
//}
//
//func (a *Controller) GetSessionHandler(c *gin.Context) {
//	reverseProxyWithSupabase(c, a.reversePath+"/user", "GET")
//}
//
//func getReversePath() string {
//	var supabaseUrl = os.Getenv("SUPABASE_URL")
//	if supabaseUrl == "" {
//		panic("Supabase URL not found")
//	}
//	return supabaseUrl + "/auth/v1"
//}
//
//func NewController() *Controller {
//	//println(getReversePath())
//	return &Controller{
//		reversePath: getReversePath(),
//	}
//}
//
//func reverseProxyWithSupabase(c *gin.Context, path string, method string) {
//	supabaseKey := os.Getenv("SUPABASE_KEY")
//	c.Header("apikey", supabaseKey)
//	utils.ReverseProxy(c, path, method)
//}

type Controller struct {
	service interfaces.AuthService
}

func NewController() *Controller {
	return &Controller{
		NewService(),
	}
}

func (ctl *Controller) LoginHandler(c *gin.Context) {
	var form entity.LoginMail
	if err := c.ShouldBindJSON(&form); err != nil {
		c.AbortWithError(400, err)
	}
	//fmt.Println(form)
	rs, err := ctl.service.Login(&form)
	if err != nil {
		authEntity.ParseError(err, -1).WriteError(c)
		return
	}
	c.JSON(200, rs)
}

func (ctl *Controller) RegisterHandler(c *gin.Context) {
	var form entity.RegisterMail
	if err := c.ShouldBindJSON(&form); err != nil {
		c.AbortWithError(400, err)
	}
	//fmt.Println(form)
	rs, err := ctl.service.Register(&form)
	if err != nil {
		authEntity.ParseError(err, -1).WriteError(c)
		return
	}

	c.JSON(200, rs)
}

func (ctl *Controller) Health(c *gin.Context) {
	res, err := ctl.service.Health()
	if err != nil {
		authEntity.ParseError(err, -1).WriteError(c)
	}
	c.JSON(200, res)
}

func (ctl *Controller) GetSessionHandler(c *gin.Context) {
	//sessionStr := c.GetString("session")
	//if sessionStr == "" {
	//	c.AbortWithStatus(401)
	//	return
	//}
	//
	//var session authEntity.AuthResponse
	//utils.Deserialize(sessionStr, &session)
	session := authUtils.GetSessionFromContext(c)

	c.JSON(200, session)

}
