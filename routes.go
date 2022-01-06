package main

import (
	"gin_vue/controller"
	"gin_vue/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	//用户信息的路由
	r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
