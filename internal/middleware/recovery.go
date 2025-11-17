package middleware

import (
	"bgame/internal/util"
	"github.com/gin-gonic/gin"
)

// Recovery 恢复panic中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				util.LogError("Panic recovered: %v", err)
				util.ErrorWithCode(c, 500, "服务器内部错误")
				c.Abort()
			}
		}()
		c.Next()
	}
}

