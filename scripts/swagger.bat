@echo off
REM Swagger 文档生成脚本（Windows）
REM 使用方法: scripts\swagger.bat

echo 正在生成 Swagger 文档...

swag init -g cmd/server/main.go -o docs

if %errorlevel% equ 0 (
    echo.
    echo ✓ Swagger 文档生成成功！
    echo.
    echo 文档位置: docs\
    echo 访问地址: http://localhost:8080/swagger/index.html
    echo.
    echo 请确保服务已启动才能访问 Swagger UI
) else (
    echo.
    echo ✗ Swagger 文档生成失败
    echo.
    echo 请检查：
    echo   1. Swag 是否已安装: swag version
    echo   2. 入口文件是否存在: cmd\server\main.go
    echo   3. 是否包含 Swagger 注释
    exit /b 1
)

pause

