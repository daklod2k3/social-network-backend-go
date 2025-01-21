package profile

import "github.com/gin-gonic/gin"

func (s *controller) ApplyRoute(r *gin.RouterGroup, authorized *gin.RouterGroup) {
	path := authorized.Group("/profile")
	path.POST("", s.CreateHdl)
	path.GET("", s.GetHdl)
}
