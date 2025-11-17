package middleware

import (
	"strings"

	"bgame/internal/util"
	"github.com/gin-gonic/gin"
)

// AuthUser 用户认证中间件
func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c)
		if token == "" {
			util.Unauthorized(c, "未提供认证token")
			c.Abort()
			return
		}

		claims, err := util.ParseToken(token)
		if err != nil {
			util.Unauthorized(c, "无效的token")
			c.Abort()
			return
		}

		if claims.Type != "user" {
			util.Unauthorized(c, "token类型错误")
			c.Abort()
			return
		}

		// 将用户信息存入上下文
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}

// AuthAdmin 管理员认证中间件
func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c)
		if token == "" {
			util.Unauthorized(c, "未提供认证token")
			c.Abort()
			return
		}

		claims, err := util.ParseToken(token)
		if err != nil {
			util.Unauthorized(c, "无效的token")
			c.Abort()
			return
		}

		if claims.Type != "admin" {
			util.Unauthorized(c, "token类型错误")
			c.Abort()
			return
		}

		// 将管理员信息存入上下文
		c.Set("admin_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}

// RequireRole 要求特定角色的中间件
func RequireRole(requiredRole int) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			util.Forbidden(c, "未获取到角色信息")
			c.Abort()
			return
		}

		adminRole, ok := role.(int)
		if !ok {
			util.Forbidden(c, "角色信息格式错误")
			c.Abort()
			return
		}

		// 超级管理员(1)拥有所有权限
		if adminRole == 1 {
			c.Next()
			return
		}

		// 检查角色权限
		if adminRole > requiredRole {
			util.Forbidden(c, "权限不足")
			c.Abort()
			return
		}

		c.Next()
	}
}

// extractToken 从请求头中提取token
func extractToken(c *gin.Context) string {
	// 优先从 Authorization header 获取
	bearerToken := c.GetHeader("Authorization")
	if bearerToken != "" {
		parts := strings.Split(bearerToken, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			return parts[1]
		}
	}

	// 从 query 参数获取
	token := c.Query("token")
	if token != "" {
		return token
	}

	return ""
}

