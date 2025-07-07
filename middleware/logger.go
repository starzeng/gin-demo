package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"io"
	"starzeng.com/gin-demo/common"
	"starzeng.com/gin-demo/pkg/logger"
	"time"
)

// 自定义 ResponseWriter 包装器
type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)                  // 写入 buffer
	return w.ResponseWriter.Write(b) // 正常写回响应体
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		traceID := c.Request.Header.Get(common.TraceIDHeader)
		if traceID == "" {
			traceID = uuid.New().String()
		}

		// 读取请求体
		var requestBody string
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			requestBody = string(bodyBytes)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 恢复 body
		}

		// 包装 ResponseWriter 捕获响应体
		bw := &bodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bw
		c.Next()

		// 日志输出
		logger.Info(
			"HTTP Request",
			zap.String(common.TraceIDKey, traceID),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("client_ip", c.ClientIP()),
			//zap.String("user_agent", c.Request.UserAgent()),
			//zap.Any("headers", c.Request.Header),
			zap.String("referer", c.Request.Referer()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("body", requestBody),
			zap.String("response", bw.body.String()),
			zap.Duration("latency", time.Since(start)),
		)
	}
}
