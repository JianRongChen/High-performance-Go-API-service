#!/bin/bash

# 开发环境设置脚本
# 使用方法: bash scripts/setup_dev.sh

echo "=== 设置开发环境 ==="

echo "1. 安装项目依赖..."
go mod download
go mod tidy

echo "2. 安装开发工具..."
echo "   安装 Air (热重载工具)..."
go install github.com/air-verse/air@latest

echo "   安装 Swag (Swagger 文档生成工具)..."
go install github.com/swaggo/swag/cmd/swag@latest

echo "3. 生成 Swagger 文档..."
swag init -g cmd/server/main.go -o docs

# 获取 GOPATH 和 bin 目录
GOPATH=$(go env GOPATH)
BIN_DIR="$GOPATH/bin"

echo ""
echo "=== 设置完成 ==="
echo ""
echo "重要提示："
echo "  如果 'air' 或 'swag' 命令找不到，请确保 Go bin 目录在 PATH 中："
echo "  export PATH=\$PATH:$BIN_DIR"
echo ""
echo "  或者使用完整路径："
echo "  $BIN_DIR/air"
echo "  $BIN_DIR/swag init -g cmd/server/main.go -o docs"
echo ""
echo "接下来可以："
echo "  - 运行 'air' 启动热重载开发（如果 PATH 已配置）"
echo "  - 或运行 '$BIN_DIR/air' 启动热重载开发"
echo "  - 运行 'swag init -g cmd/server/main.go -o docs' 重新生成文档"
echo "  - 访问 http://localhost:8080/swagger/index.html 查看 API 文档"

