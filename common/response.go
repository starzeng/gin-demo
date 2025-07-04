package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 定义了接口响应的通用结构
type Response struct {
	Code    int         `json:"code"`    // Code 表示响应状态码
	Message string      `json:"message"` // Message 表示响应消息
	Data    interface{} `json:"data"`    // Data 表示响应携带的数据
}

// 定义接口响应的状态码常量
const (
	CodeSuccess       = 0    // CodeSuccess 表示请求成功
	CodeInvalidParams = 1001 // CodeInvalidParams 表示请求参数无效
	CodeUnauthorized  = 1002 // CodeUnauthorized 表示未授权访问
	CodeForbidden     = 1003 // CodeForbidden 表示禁止访问
	CodeInternalError = 1004 // CodeInternalError 表示服务器内部错误
	CodeTokenExpired  = 1005 // CodeTokenExpired 表示令牌过期
)

// Success 用于返回成功响应
// 参数 c 是 gin 的上下文对象，用于处理 HTTP 请求和响应
// 参数 data 是需要返回给客户端的数据
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{CodeSuccess, "success", data})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{code, msg, nil})
}
