package profile

import "github.com/gin-gonic/gin"

func (s *controller) ApplyRoute(r *gin.RouterGroup) {
	path := r.Group("/profile")
	path.POST("", s.CreateHdl)
}
