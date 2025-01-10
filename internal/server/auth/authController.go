package auth

import (
	"auth/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	shared "shared/entity"
	"strings"
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
	service Service
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
		ErrorHandler(c, err.Error())
		return
	}
	c.JSON(200, rs)
}

func (ctl *Controller) RegisterHandler(c *gin.Context) {
	var form entity.RegisterEmail
	if err := c.ShouldBindJSON(&form); err != nil {
		c.AbortWithError(400, err)
	}
	//fmt.Println(form)
	rs, err := ctl.service.Register(form)
	if err != nil {
		ErrorHandler(c, err.Error())
		return
	}
	c.JSON(200, rs)
}

func (ctl *Controller) GetSessionHandler(c *gin.Context) {
}

func ErrorHandler(c *gin.Context, err string) {
	//fmt.Println(err)
	switch {
	case strings.IndexAny(err, "code = Unknown") == -1 && strings.IndexAny(err, "code = Unavailable") > -1:
		fmt.Println(err)
		shared.WriteError(c, 500, "fail to connect auth service")

	default:
		//fmt.Println("spErr")
		spErr := entity.Error(err)
		shared.WriteError(c, spErr.Code, spErr.Msg)
	}
}
