package router

import (
	"bgame/docs"
	"bgame/internal/config"
	"bgame/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	// 根据配置设置gin模式
	mode := config.Cfg.Server.Mode
	switch mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// 全局中间件
	r.Use(middleware.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())
	r.Use(middleware.RateLimit())

	// 健康检查
	// @Summary      健康检查
	// @Description  检查服务是否正常运行
	// @Tags         系统
	// @Accept       json
	// @Produce      json
	// @Success      200  {object}  map[string]string
	// @Router       /health [get]
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Swagger 文档
	docs.SwaggerInfo.Title = "bGame API 文档"
	docs.SwaggerInfo.Description = "高性能 Go API 服务，单机 QPS > 20000"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 设置路由
	setupUserRoutes(r)
	setupAdminRoutes(r)

	return r
}

