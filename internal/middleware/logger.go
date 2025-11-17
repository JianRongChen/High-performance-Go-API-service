package middleware

import (
	"bgame/internal/util"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		// 根据状态码选择日志级别
		logMsg := fmt.Sprintf("[%s] %s %s %d %v %s",
			clientIP,
			method,
			path,
			statusCode,
			latency,
			c.Errors.String(),
		)

		// 错误状态码（4xx, 5xx）记录到 error 日志
		if statusCode >= 400 {
			util.LogError(logMsg)
		} else {
			util.Info(logMsg)
		}
	}
}

