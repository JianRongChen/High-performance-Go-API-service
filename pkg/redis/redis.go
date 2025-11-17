package redis

import (
	"context"
	"fmt"
	"time"

	"bgame/internal/config"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client
var ctx = context.Background()

func Init() error {
	cfg := config.Cfg
	Client = redis.NewClient(&redis.Options{
		Addr:         cfg.GetRedisAddr(),
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.Database,
		PoolSize:     cfg.Redis.PoolSize,
		MinIdleConns: cfg.Redis.MinIdleConns,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	// 测试连接
	if err := Client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("连接Redis失败: %w", err)
	}

	return nil
}

func Close() error {
	if Client != nil {
		return Client.Close()
	}
	return nil
}

func Get(ctx context.Context) *redis.Client {
	return Client
}

