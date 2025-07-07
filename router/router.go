package router

import (
	"github.com/gin-gonic/gin"
	"starzeng.com/gin-demo/config"
	"strconv"
)

// InitRouter 初始化路由
func InitRouter(r *gin.Engine) {
	path := config.AppConfig.Server.RelativePath
	api := r.Group(path)
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

	println("注册的 controller 数量：" + strconv.Itoa(len(routeRegisters)))

	for _, register := range routeRegisters {
		register.RouteRegister(rg)
	}
}
