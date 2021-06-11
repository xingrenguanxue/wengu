package routes

import (
	"wengu/conf"
	"wengu/routes/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() {
	gin.SetMode(conf.RunMode)
	r := gin.New()
	r.POST("signup", v1.SignUp)
	r.Run(conf.HttpPort)
}