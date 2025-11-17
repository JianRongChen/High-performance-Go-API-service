package util

import (
	"errors"
	"time"

	"bgame/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Type     string `json:"type"` // "user" or "admin"
	Role     int    `json:"role,omitempty"` // admin role
	jwt.RegisteredClaims
}

// GenerateUserToken 生成用户token
func GenerateUserToken(userID uint, username string) (string, error) {
	cfg := config.Cfg
	expireTime := time.Now().Add(time.Duration(cfg.JWT.UserExpire) * time.Second)

	claims := &Claims{
		UserID:   userID,
		Username: username,
		Type:     "user",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.Secret))
}

// GenerateAdminToken 生成管理员token
func GenerateAdminToken(adminID uint, username string, role int) (string, error) {
	cfg := config.Cfg
	expireTime := time.Now().Add(time.Duration(cfg.JWT.AdminExpire) * time.Second)

	claims := &Claims{
		UserID:   adminID,
		Username: username,
		Type:     "admin",
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.Secret))
}

// ParseToken 解析token
func ParseToken(tokenString string) (*Claims, error) {
	cfg := config.Cfg
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("无效的签名方法")
		}
		return []byte(cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的token")
}

