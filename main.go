package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"starzeng.com/gin-demo/config"
	_ "starzeng.com/gin-demo/controller"
	"starzeng.com/gin-demo/docs"
	"starzeng.com/gin-demo/middleware"
	"starzeng.com/gin-demo/router"
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

	r := gin.Default()

	// 初始化路由
	router.InitRouter(r)

	// Swagger 配置
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 全局错误处理
	r.Use(gin.Logger(), middleware.RecoveryWithJSON())

	err := r.Run(config.GetServerAddr())
	if err != nil {
		return
	}

}
