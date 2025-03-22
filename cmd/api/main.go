package main

import (
	"log"
	"go_server/internal/config"
	"go_server/internal/handler"
	"go_server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	if err := config.Load(); err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// 初始化路由
	r := gin.Default()

	// 注册中间件
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(middleware.Cors())

	// 注册路由
	api := r.Group("/api")
	{
		// 认证相关路由
		auth := api.Group("/auth")
		{
			auth.POST("/register", handler.Register)
			auth.POST("/login", handler.Login)
			auth.POST("/send-code", handler.SendVerificationCode)
		}

		// 用户相关路由
		user := api.Group("/user")
		user.Use(middleware.Auth())
		{
			user.GET("/profile", handler.GetProfile)
			user.PUT("/profile", handler.UpdateProfile)
		}
	}

	// 启动服务器
	r.Run(":3000")
} 