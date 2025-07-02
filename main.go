package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"starzeng.com/gin-demo/config"
	"starzeng.com/gin-demo/docs"
	"starzeng.com/gin-demo/handler"
	"starzeng.com/gin-demo/middleware"
)

// @title Gin JWT RBAC Demo API
// @version 1.0
// @description 示例项目：JWT + RBAC + Redis 黑名单 + Swagger 接口文档
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// 加载配置
	config.LoadConfig()

	// 初始化
	config.InitMySQL()
	config.InitRedis()

	gin.SetMode(config.AppConfig.Server.Mode)
	gin.ForceConsoleColor()
	r := gin.Default()

	// Swagger 配置
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 全局错误处理
	r.Use(gin.Logger(), middleware.RecoveryWithJSON())

	r.POST("/login", handler.Login)

	auth := r.Group("/api", middleware.JWTAuth())
	auth.GET("/profile", handler.Profile)
	auth.GET("/logout", handler.Logout)
	auth.GET("/admin", middleware.RequireRole("admin"), handler.AdminOnly)
	auth.POST("/write", middleware.RequirePermission("write"), handler.WriteData)

	err := r.Run(config.AppConfig.Server.Port)
	if err != nil {
		return
	}

}
