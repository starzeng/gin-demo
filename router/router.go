package router

import (
	"github.com/gin-gonic/gin"
	"starzeng.com/gin-demo/config"
)

// InitRouter 初始化路由
func InitRouter(r *gin.Engine) {
	api := r.Group(config.AppConfig.Server.RelativePath)
	loadRoutes(api)
}

// RouteRegister 定义controller注册接口
type RouteRegister interface {
	RouteRegister(group *gin.RouterGroup)
}

// 控制器注册列表
var routeRegisters []RouteRegister

// Register 注册, 初始化使用
func Register(r RouteRegister) {
	routeRegisters = append(routeRegisters, r)
}

// loadRoutes 注册入口
func loadRoutes(rg *gin.RouterGroup) {
	for _, register := range routeRegisters {
		register.RouteRegister(rg)
	}
}
