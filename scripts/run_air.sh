#!/bin/bash

# 快速启动 Air 热重载
# 使用方法: bash scripts/run_air.sh

# 获取 GOPATH 和 bin 目录
GOPATH=$(go env GOPATH)
BIN_DIR="$GOPATH/bin"

# 确保 PATH 包含 Go bin 目录
export PATH="$PATH:$BIN_DIR"

# 检测操作系统并设置正确的二进制文件名
if [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" || -n "$WINDIR" ]]; then
    # Windows 系统
    BIN_NAME="main.exe"
    echo "检测到 Windows 系统，使用 $BIN_NAME"
else
    # Linux/Mac 系统
    BIN_NAME="main"
    echo "检测到 Linux/Mac 系统，使用 $BIN_NAME"
fi

# 检查并更新 .air.toml 配置
if [ -f ".air.toml" ]; then
    # 检查配置是否需要更新
    if grep -q 'bin = "./tmp/main"' .air.toml && [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" || -n "$WINDIR" ]]; then
        echo "更新 .air.toml 配置以适配 Windows..."
        sed -i 's|bin = "./tmp/main"|bin = "./tmp/main.exe"|g' .air.toml
        sed -i 's|cmd = "go build -o ./tmp/main|cmd = "go build -o ./tmp/main.exe|g' .air.toml
    fi
fi

# 检查 air 是否可用
if command -v air > /dev/null 2>&1; then
    echo "启动 Air 热重载..."
    air
else
    echo "错误: air 命令未找到"
    echo ""
    echo "请先安装 Air:"
    echo "  go install github.com/air-verse/air@latest"
    echo ""
    echo "或使用完整路径:"
    echo "  $BIN_DIR/air"
    exit 1
fi

