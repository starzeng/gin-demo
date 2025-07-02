package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	CodeSuccess       = 0
	CodeInvalidParams = 1001
	CodeUnauthorized  = 1002
	CodeForbidden     = 1003
	CodeInternalError = 1004
	CodeTokenExpired  = 1005
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{CodeSuccess, "success", data})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{code, msg, nil})
}
