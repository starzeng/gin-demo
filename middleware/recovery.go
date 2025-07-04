package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
	"starzeng.com/gin-demo/common"
)

func RecoveryWithJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[PANIC] %v\n%s", err, debug.Stack())
				common.Error(c, common.CodeInternalError, "服务器内部错误")
				c.Abort()
			}
		}()
		c.Next()
	}
}
