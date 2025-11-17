@echo off
REM Windows 快速启动 Air 热重载
REM 使用方法: scripts\run_air.bat

REM 获取 GOPATH
for /f "tokens=*" %%i in ('go env GOPATH') do set GOPATH=%%i
set BIN_DIR=%GOPATH%\bin

REM 检查 air.exe 是否存在
if exist "%BIN_DIR%\air.exe" (
    echo 启动 Air 热重载...
    "%BIN_DIR%\air.exe"
) else (
    echo 错误: air.exe 未找到
    echo.
    echo 请先安装 Air:
    echo   go install github.com/air-verse/air@latest
    echo.
    echo 或使用完整路径:
    echo   %BIN_DIR%\air.exe
    pause
    exit /b 1
)

