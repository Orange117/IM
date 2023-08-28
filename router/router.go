package router

import (
	"IM/middlewares"
	"IM/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 用户登录
	r.POST("/login", service.Login)
	// 发送验证码
	r.POST("/send/code", service.SendCode)

	auth := r.Group("/u", middlewares.AuthCheck())

	//用户详情
	auth.GET("/user/detail", service.UserDetail)
	return r
}
