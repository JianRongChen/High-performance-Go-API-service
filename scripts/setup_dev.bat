@echo off
REM Windows 开发环境设置脚本
REM 使用方法: scripts\setup_dev.bat

echo === 设置开发环境 ===

echo 1. 安装项目依赖...
go mod download
go mod tidy

echo 2. 安装开发工具...
echo    安装 Air (热重载工具)...
go install github.com/air-verse/air@latest

echo    安装 Swag (Swagger 文档生成工具)...
go install github.com/swaggo/swag/cmd/swag@latest

echo 3. 生成 Swagger 文档...
swag init -g cmd/server/main.go -o docs

echo.
echo === 设置完成 ===
echo.
echo 接下来可以：
echo   - 运行 'air' 启动热重载开发
echo   - 运行 'swag init -g cmd/server/main.go -o docs' 重新生成文档
echo   - 或使用 'make.bat air' 和 'make.bat swagger'
echo   - 访问 http://localhost:8080/swagger/index.html 查看 API 文档

pause

