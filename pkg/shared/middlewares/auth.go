package middlewares

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	authEntity "shared/entity/auth"
	logger2 "shared/logger"
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

var (
	logger = logger2.GetLogger()
)

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

		//logger.Info(accessToken + " " + refreshToken)

		session, err := authService.GetSession(&authEntity.SessionRequest{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})

		//logger.Info(fmt.Sprintln(*session))

		if err != nil {
			//authEntity.ParseError(err, 401).WriteError(c)
			logger.Error(err.Error())
			c.AbortWithStatus(401)
			return
		}

		if session.User == nil {
			c.Header("X-Profile", "not-created")
		}

		js, err := json.Marshal(session)

		if err != nil {
			logger.Error(err.Error())
		}

		c.Set("session", string(js))
		//logger.Info(session)
		c.Next()

	}
}
