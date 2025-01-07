package auth

import (
	"core/internal/server/auth/entity"
	shared "core/shared/entity"
	"github.com/gin-gonic/gin"
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
	*service
}

func NewController() *Controller {
	return &Controller{
		NewService(),
	}
}

func (ctl *Controller) LoginHandler(c *gin.Context) {
	var form entity.LoginEmail
	if err := c.ShouldBindJSON(&form); err != nil {
		c.AbortWithError(400, err)
	}
	//fmt.Println(form)
	rs, err := ctl.service.Login(form)
	if err != nil {
		shared.WriteError(c, 400, ctl.service.Error(err).Msg)
		return
	}
	c.JSON(200, rs)
}

func (ctl *Controller) RegisterHandler(c *gin.Context) {
	var form entity.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		c.AbortWithError(400, err)
	}
	//fmt.Println(form)
	rs, err := ctl.service.Register(form)
	if err != nil {
		shared.WriteError(c, 400, ctl.service.Error(err).Msg)
		return
	}
	c.JSON(200, rs)
}

func (ctl *Controller) GetSessionHandler(c *gin.Context) {
}
