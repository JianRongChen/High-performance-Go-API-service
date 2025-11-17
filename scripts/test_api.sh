#!/bin/bash

# API 测试脚本
# 使用方法: bash scripts/test_api.sh

BASE_URL="http://localhost:8080"

echo "=== 测试健康检查 ==="
curl -X GET "${BASE_URL}/health"
echo -e "\n\n"

echo "=== 测试用户注册 ==="
curl -X POST "${BASE_URL}/api/user/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "123456",
    "email": "test@example.com",
    "nickname": "测试用户"
  }'
echo -e "\n\n"

echo "=== 测试用户登录 ==="
TOKEN=$(curl -s -X POST "${BASE_URL}/api/user/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "123456"
  }' | grep -o '"token":"[^"]*' | cut -d'"' -f4)

echo "Token: $TOKEN"
echo -e "\n\n"

echo "=== 测试获取用户信息 ==="
curl -X GET "${BASE_URL}/api/user/info" \
  -H "Authorization: Bearer ${TOKEN}"
echo -e "\n\n"

echo "=== 测试管理员登录 ==="
ADMIN_TOKEN=$(curl -s -X POST "${BASE_URL}/api/admin/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123"
  }' | grep -o '"token":"[^"]*' | cut -d'"' -f4)

echo "Admin Token: $ADMIN_TOKEN"
echo -e "\n\n"

echo "=== 测试获取管理员信息 ==="
curl -X GET "${BASE_URL}/api/admin/info" \
  -H "Authorization: Bearer ${ADMIN_TOKEN}"
echo -e "\n\n"

echo "=== 测试获取角色列表 ==="
curl -X GET "${BASE_URL}/api/admin/roles"
echo -e "\n"

