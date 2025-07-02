package controller

import (
	"github.com/gin-gonic/gin"
	"starzeng.com/gin-demo/common"
	"starzeng.com/gin-demo/middleware"
	"starzeng.com/gin-demo/model"
	"starzeng.com/gin-demo/router"
	"starzeng.com/gin-demo/utils"
	"time"
)

type userController struct{}

func (u userController) RouteRegister(group *gin.RouterGroup) {
	group.POST("/login", Login)

	auth := group.Group("/user", middleware.JWTAuth())

	auth.GET("/profile", Profile)
	auth.GET("/logout", Logout)
	auth.GET("/admin", middleware.RequireRole("admin"), AdminOnly)
	auth.POST("/write", middleware.RequirePermission("write"), WriteData)

}

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login 用户登录
// @Summary 用户登录
// @Description 使用用户名和密码进行登录，返回 JWT Token
// @Tags 用户
// @Accept json
// @Produce json
// @Param data body LoginReq true "登录请求参数"
// @Success 200 {object} common.Response
// @Failure 401 {object} common.Response
// @Router /api/login [post]
func Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, common.CodeInvalidParams, err.Error())
		return
	}
	u, ok := model.Users[req.Username]
	if !ok || u.Password != req.Password {
		common.Error(c, common.CodeUnauthorized, "用户名或密码错误")
		return
	}
	token, jti, _ := middleware.GenerateToken(u)
	common.Success(c, gin.H{
		"token":  token,
		"expire": middleware.TokenExpireDuration.Seconds(),
		"jti":    jti,
	})
}

// Logout 用户登出
// @Summary 用户登出
// @Description 将当前 JWT 加入黑名单
// @Tags 用户
// @Security BearerAuth
// @Produce json
// @Success 200 {object} common.Response
// @Router /api/logout [get]
func Logout(c *gin.Context) {
	jti := c.GetString("jti")
	exp := time.Minute * 5
	utils.BlacklistAdd(jti, int64(exp.Seconds()))
	common.Success(c, gin.H{"message": "退出成功"})
}

// Profile 获取用户信息
// @Summary 获取用户信息
// @Description 返回当前登录用户的用户名与角色
// @Tags 用户
// @Security BearerAuth
// @Produce json
// @Success 200 {object} common.Response
// @Router /api/profile [get]
func Profile(c *gin.Context) {
	utils.Set("hello world")
	common.Success(c, gin.H{"username": c.GetString("username"), "role": c.GetString("role")})
}

// AdminOnly 管理员接口
// @Summary 仅管理员可访问
// @Description 仅角色为 admin 的用户可访问
// @Tags 权限
// @Security BearerAuth
// @Produce json
// @Success 200 {object} common.Response
// @Failure 403 {object} common.Response
// @Router /api/admin [get]
func AdminOnly(c *gin.Context) {
	common.Success(c, gin.H{"msg": "admin 访问成功"})
}

// WriteData 写入数据接口
// @Summary 写数据
// @Description 需要拥有 write 权限的用户才可以访问
// @Tags 权限
// @Security BearerAuth
// @Produce json
// @Success 200 {object} common.Response
// @Failure 403 {object} common.Response
// @Router /api/write [post]
func WriteData(c *gin.Context) {
	common.Success(c, gin.H{"msg": "写入成功"})
}

func init() {
	router.Register(&userController{})
}
