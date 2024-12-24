package controllers

import (
	"github.com/gin-gonic/gin"
	"os"
	"social-backend/internal/server/utils"
)

var supabaseUrl = os.Getenv("SUPABASE_URL")
var reversePath = supabaseUrl + "/auth/v1"

func Login(c *gin.Context) {
	utils.ReverseProxy(c, reversePath+"/token?grant_type=password", "POST")
}

func Register(c *gin.Context) {
	utils.ReverseProxy(c, reversePath+"/signup", "POST")
}
