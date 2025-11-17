#!/bin/bash

# 检查 Swagger 配置
# 使用方法: bash scripts/check_swagger.sh

echo "=== 检查 Swagger 配置 ==="
echo ""

# 检查 docs 目录是否存在
if [ -d "docs" ]; then
    echo "✓ docs 目录存在"
    
    # 检查关键文件
    if [ -f "docs/docs.go" ]; then
        echo "✓ docs/docs.go 存在"
    else
        echo "✗ docs/docs.go 不存在"
        echo "  请运行: swag init -g cmd/server/main.go -o docs"
    fi
    
    if [ -f "docs/swagger.json" ]; then
        echo "✓ docs/swagger.json 存在"
    else
        echo "✗ docs/swagger.json 不存在"
    fi
    
    if [ -f "docs/swagger.yaml" ]; then
        echo "✓ docs/swagger.yaml 存在"
    else
        echo "✗ docs/swagger.yaml 不存在"
    fi
else
    echo "✗ docs 目录不存在"
    echo "  请运行: swag init -g cmd/server/main.go -o docs"
fi

echo ""

# 检查 main.go 中是否有 Swagger 注释
if grep -q '@title' cmd/server/main.go; then
    echo "✓ cmd/server/main.go 包含 Swagger 注释"
else
    echo "✗ cmd/server/main.go 缺少 Swagger 注释"
fi

# 检查 router.go 中是否导入了 docs
if grep -q '"bgame/docs"' internal/router/router.go; then
    echo "✓ internal/router/router.go 已导入 docs 包"
else
    echo "✗ internal/router/router.go 未导入 docs 包"
    echo "  需要在 import 中添加: \"bgame/docs\""
fi

# 检查 Swagger 路由是否配置
if grep -q '/swagger/\*any' internal/router/router.go; then
    echo "✓ Swagger 路由已配置"
else
    echo "✗ Swagger 路由未配置"
fi

echo ""
echo "=== 检查完成 ==="
echo ""
echo "如果所有检查都通过，Swagger 文档应该可以通过以下地址访问："
echo "  http://localhost:8080/swagger/index.html"
echo ""
echo "如果服务未运行，请先启动服务："
echo "  go run cmd/server/main.go"
echo "  或"
echo "  air"

