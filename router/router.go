package router

import (
	"IM/middlewares"
	"IM/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 用户注册
	r.POST("/register", service.Register)

	// 用户登录
	r.POST("/login", service.Login)
	// 发送验证码
	r.POST("/send/code", service.SendCode)

	auth := r.Group("/u", middlewares.AuthCheck())

	// 用户详情
	auth.GET("/user/detail", service.UserDetail)
	// 查询指定用户的个人信息
	auth.GET("/user/query", service.UserQuery)
	// 发送、接受消息
	auth.GET("/websocket/message", service.WebsocketMessage)
	// 聊天记录列表
	auth.GET("/chat/list", service.ChatList)

	return r
}
