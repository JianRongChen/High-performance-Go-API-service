package main

import (
	"fmt"
	"log"

	"bgame/internal/util"
)

// 用于生成管理员密码的工具脚本
// 使用方法: go run scripts/create_admin.go
func main() {
	password := "admin123"
	hashed, err := util.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("原始密码: %s\n", password)
	fmt.Printf("加密后的密码: %s\n", hashed)
	fmt.Println("\n可以将此密码用于数据库初始化脚本")
}

