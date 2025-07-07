package main

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"starzeng.com/gin-demo/config"
	"starzeng.com/gin-demo/middleware"
	"starzeng.com/gin-demo/pkg/db"
	"starzeng.com/gin-demo/pkg/redis"
	"starzeng.com/gin-demo/router"

	_ "starzeng.com/gin-demo/internal/book/controller"
	_ "starzeng.com/gin-demo/internal/user/controller"
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

	// 初始化MySQL数据库
	db.InitMySQL()
	// 集中统一迁移所有模型
	if err := db.AutoMigrate(db.DB); err != nil {
		panic("模型迁移失败: " + err.Error())
	}

	// 初始化redis
	redis.InitRedis()

	r := gin.Default()
	r.Use(middleware.RecoveryWithJSON())

	// 初始化路由
	router.InitRouter(r)

	// Swagger 配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(config.GetServerAddr())
	if err != nil {
		return
	}

}
