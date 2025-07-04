package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"starzeng.com/gin-demo/config"
	"starzeng.com/gin-demo/docs"
	_ "starzeng.com/gin-demo/internal/book/controller"
	_ "starzeng.com/gin-demo/internal/user/controller"
	"starzeng.com/gin-demo/middleware"
	"starzeng.com/gin-demo/pkg/db"
	"starzeng.com/gin-demo/pkg/redis"
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
	db.InitMySQL()
	redis.InitRedis()

	// 集中统一迁移所有模型
	if err := db.AutoMigrate(db.DB); err != nil {
		panic("模型迁移失败: " + err.Error())
	}

	r := gin.Default()

	// 初始化路由
	router.InitRouter(r)

	// Swagger 配置
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	println("API文档: http://localhost:8080/swagger/index.html")

	// 全局错误处理
	r.Use(gin.Logger(), middleware.RecoveryWithJSON())

	err := r.Run(config.GetServerAddr())
	if err != nil {
		return
	}

}
