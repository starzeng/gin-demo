package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
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
	gin.SetMode(gin.DebugMode)
	// 强制日志颜色化
	gin.ForceConsoleColor()

	r := gin.Default()

	// Swagger 配置
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(gin.Logger(), middleware.RecoveryWithJSON())

	r.POST("/login", handler.Login)

	auth := r.Group("/api", middleware.JWTAuth())
	auth.GET("/profile", handler.Profile)
	auth.GET("/logout", handler.Logout)
	auth.GET("/admin", middleware.RequireRole("admin"), handler.AdminOnly)
	auth.POST("/write", middleware.RequirePermission("write"), handler.WriteData)

	err := r.Run(":8080")
	if err != nil {
		return
	}

}

func getTest(router *gin.Engine) gin.IRoutes {
	return router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
			"hello":   123,
		})
	})
}
