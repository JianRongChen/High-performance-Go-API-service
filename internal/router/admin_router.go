package router

import (
	"bgame/internal/handler/admin"
	"bgame/internal/middleware"

	"github.com/gin-gonic/gin"
)

func setupAdminRoutes(r *gin.Engine) {
	adminHandler := admin.NewAdminHandler()
	adminGroup := r.Group("/api/admin")
	{
		// 公开接口
		adminGroup.POST("/login", adminHandler.Login)
		adminGroup.GET("/roles", adminHandler.GetRoles)
		adminGroup.POST("/create", adminHandler.CreateAdmin)

		// 需要认证的接口
		adminGroup.Use(middleware.AuthAdmin())
		{
			adminGroup.GET("/info", adminHandler.GetAdminInfo)

			// 需要超级管理员权限的接口
			adminGroup.Use(middleware.RequireRole(1))
			{
				// adminGroup.POST("/create", adminHandler.CreateAdmin)
			}
		}
	}
}
