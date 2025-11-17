#!/bin/bash

# 检查开发工具是否已安装
# 使用方法: bash scripts/check_tools.sh

echo "=== 检查开发工具安装状态 ==="
echo ""

# 获取 GOPATH
GOPATH=$(go env GOPATH)
BIN_DIR="$GOPATH/bin"

echo "GOPATH: $GOPATH"
echo "Bin 目录: $BIN_DIR"
echo ""

# 检查 PATH 是否包含 Go bin 目录
if [[ ":$PATH:" == *":$BIN_DIR:"* ]]; then
    echo "✓ Go bin 目录已在 PATH 中"
else
    echo "✗ Go bin 目录不在 PATH 中"
    echo "  请运行: export PATH=\$PATH:$BIN_DIR"
    echo "  或添加到 ~/.bashrc: echo 'export PATH=\$PATH:$BIN_DIR' >> ~/.bashrc"
fi
echo ""

# 检查 Air
if command -v air > /dev/null 2>&1; then
    echo "✓ Air 已安装: $(which air)"
    air -v 2>/dev/null || echo "  (版本信息不可用)"
else
    echo "✗ Air 未安装"
    echo "  安装命令: go install github.com/air-verse/air@latest"
    if [ -f "$BIN_DIR/air.exe" ] || [ -f "$BIN_DIR/air" ]; then
        echo "  注意: 找到 air 可执行文件，但不在 PATH 中"
        echo "  请运行: export PATH=\$PATH:$BIN_DIR"
    fi
fi
echo ""

# 检查 Swag
if command -v swag > /dev/null 2>&1; then
    echo "✓ Swag 已安装: $(which swag)"
    swag version 2>/dev/null || echo "  (版本信息不可用)"
else
    echo "✗ Swag 未安装"
    echo "  安装命令: go install github.com/swaggo/swag/cmd/swag@latest"
    if [ -f "$BIN_DIR/swag.exe" ] || [ -f "$BIN_DIR/swag" ]; then
        echo "  注意: 找到 swag 可执行文件，但不在 PATH 中"
        echo "  请运行: export PATH=\$PATH:$BIN_DIR"
    fi
fi
echo ""

echo "=== 检查完成 ==="

