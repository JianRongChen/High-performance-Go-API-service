package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bgame/internal/config"
	"bgame/internal/model"
	"bgame/internal/router"
	"bgame/internal/util"
	"bgame/pkg/mysql"
	"bgame/pkg/redis"
)

// @title           bGame API 文档
// @version         1.0
// @description     高性能 Go API 服务，单机 QPS > 20000
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// 加载配置
	configPath := "config.yaml"
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	if err := config.LoadConfig(configPath); err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化日志系统
	if err := util.InitLogger(); err != nil {
		log.Fatalf("初始化日志系统失败: %v", err)
	}
	util.Info("日志系统初始化成功")

	// 初始化MySQL
	if err := mysql.Init(); err != nil {
		util.LogError("初始化MySQL失败: %v", err)
		log.Fatalf("初始化MySQL失败: %v", err)
	}
	defer mysql.Close()
	util.Info("MySQL 连接成功")

	// 初始化Redis
	if err := redis.Init(); err != nil {
		util.LogError("初始化Redis失败: %v", err)
		log.Fatalf("初始化Redis失败: %v", err)
	}
	defer redis.Close()
	util.Info("Redis 连接成功")

	// 自动迁移数据库表
	if err := autoMigrate(); err != nil {
		util.LogError("数据库迁移失败: %v", err)
		log.Fatalf("数据库迁移失败: %v", err)
	}
	util.Info("数据库表迁移完成")

	// 设置路由
	r := router.SetupRouter()

	// 创建HTTP服务器
	srv := &http.Server{
		Addr:           config.Cfg.GetServerAddr(),
		Handler:        r,
		ReadTimeout:    config.Cfg.GetReadTimeout(),
		WriteTimeout:   config.Cfg.GetWriteTimeout(),
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	// 启动服务器（goroutine）
	go func() {
		log.Printf("服务器启动在 %s", config.Cfg.GetServerAddr())
		log.Printf("API 文档地址: http://%s/swagger/index.html", config.Cfg.GetServerAddr())
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			util.LogError("服务器启动失败: %v", err)
			log.Fatalf("服务器启动失败: %v", err)
		}
	}()

	// 等待中断信号以优雅关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	util.Info("正在关闭服务器...")
	log.Println("正在关闭服务器...")

	// 5秒超时关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		util.LogError("服务器强制关闭: %v", err)
		log.Fatalf("服务器强制关闭: %v", err)
	}

	util.Info("服务器已关闭")
	log.Println("服务器已关闭")
}

// autoMigrate 自动迁移数据库表
func autoMigrate() error {
	if err := mysql.DB.AutoMigrate(
		&model.User{},
		&model.Admin{},
	); err != nil {
		return fmt.Errorf("数据库迁移失败: %w", err)
	}
	return nil
}
