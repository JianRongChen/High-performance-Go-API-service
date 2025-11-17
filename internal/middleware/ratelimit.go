package middleware

import (
	"context"
	"fmt"
	"time"

	"bgame/internal/config"
	redisPkg "bgame/pkg/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// RateLimit 基于Redis的限流中间件
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !config.Cfg.RateLimit.Enabled {
			c.Next()
			return
		}

		// 使用IP作为限流key
		clientIP := c.ClientIP()
		key := fmt.Sprintf("ratelimit:%s", clientIP)

		ctx := context.Background()
		rdb := redisPkg.Client

		// 使用滑动窗口算法
		now := time.Now().Unix()
		windowStart := now - int64(config.Cfg.RateLimit.RPS)

		// 清理过期记录
		rdb.ZRemRangeByScore(ctx, key, "0", fmt.Sprintf("%d", windowStart))

		// 获取当前窗口内的请求数
		count, err := rdb.ZCard(ctx, key).Result()
		if err != nil {
			c.Next()
			return
		}

		// 检查是否超过限制
		if count >= int64(config.Cfg.RateLimit.Burst) {
			c.JSON(429, gin.H{
				"code":    429,
				"message": "请求过于频繁，请稍后再试",
			})
			c.Abort()
			return
		}

		// 添加当前请求记录
		rdb.ZAdd(ctx, key, &redis.Z{
			Score:  float64(now),
			Member: fmt.Sprintf("%d", now),
		})

		// 设置key过期时间
		rdb.Expire(ctx, key, time.Duration(config.Cfg.RateLimit.RPS)*time.Second)

		c.Next()
	}
}

