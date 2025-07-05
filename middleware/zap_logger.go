package middleware

import (
	"github.com/gin-gonic/gin"
	"starzeng.com/gin-demo/pkg/logger"
	"time"
)

func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		c.Next()
		cost := time.Since(start)
		logger.Infow("HTTP Request",
			"method", c.Request.Method,
			"path", path,
			"status", c.Writer.Status(),
			"latency", cost.String(),
			"ip", c.ClientIP(),
		)
	}
}
