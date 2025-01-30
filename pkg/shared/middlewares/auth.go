package middlewares

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	authEntity "shared/entity/auth"
	"shared/internal/global"
	authRpcClient "shared/rpc/client/auth"
	"strings"
)

//type middleware struct {
//	authService *interfaces.AuthService
//	Middleware
//}
//
//type Middleware interface {
//	AuthMiddleware(c *gin.Context) gin.HandlerFunc
//}

//func NewMiddleware(authService *interfaces.AuthService) gin.HandlerFunc {
//
//}

func AuthMiddleware(authService authRpcClient.AuthRpcService) gin.HandlerFunc {
	if authService == nil {
		panic(errors.New("auth service is nil"))
	}
	return func(c *gin.Context) {
		var (
			accessToken  = c.Request.Header.Get("Authorization")
			refreshToken string
		)

		// remove bearer
		accessToken = strings.Replace(accessToken, "Bearer ", "", 1)

		if accessToken == "" {
			cookie, err := c.Request.Cookie("access_token")
			if err == nil {
				accessToken = cookie.Value
			}
		}

		cookie, err := c.Request.Cookie("refresh_token")
		if err == nil {
			refreshToken = cookie.Value
		}

		if refreshToken == "" && accessToken == "" {
			authEntity.ParseError(errors.New("no authorization info found"), 401).WriteError(c)
			return
		}

		//global.Logger.Info(refreshToken)

		session, err := authService.GetSession(&authEntity.SessionRequest{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})

		//logger.Info(fmt.Sprintln(*session))

		if err != nil {
			//authEntity.ParseError(err, 401).WriteError(c)
			global.Logger.Info(err.Error())
			c.AbortWithStatus(401)
			return
		}

		if session.User == nil {
			c.Header("X-Profile", "not-created")
		}

		js, err := json.Marshal(session)
		//global.Logger.Sugar().Infoln(session)
		//fmt.Printf("%#v", session)

		if err != nil {
			global.Logger.Error(err.Error())
		}

		c.Set("session", string(js))
		//logger.Info(session)
		c.Next()

	}
}
