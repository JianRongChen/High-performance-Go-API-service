package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"bgame/internal/config"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
	logMutex    sync.Mutex
	currentDate string
	logDir      string
)

// InitLogger 初始化日志系统
func InitLogger() error {
	logMutex.Lock()
	defer logMutex.Unlock()

	// 设置日志目录
	logDir = "logs"
	if config.Cfg != nil && config.Cfg.Log.Dir != "" {
		logDir = config.Cfg.Log.Dir
	}

	// 创建日志目录
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("创建日志目录失败: %w", err)
	}

	// 初始化日志
	return updateLoggers()
}

// updateLoggers 更新日志文件（按日期）
func updateLoggers() error {
	today := time.Now().Format("2006-01-02")

	// 如果日期没变，不需要更新
	if currentDate == today && infoLogger != nil && errorLogger != nil {
		return nil
	}

	currentDate = today

	// Info 日志文件
	infoFile, err := openLogFile(filepath.Join(logDir, fmt.Sprintf("%s.info.log", today)))
	if err != nil {
		return fmt.Errorf("打开 info 日志文件失败: %w", err)
	}

	// Error 日志文件
	errorFile, err := openLogFile(filepath.Join(logDir, fmt.Sprintf("%s.error.log", today)))
	if err != nil {
		return fmt.Errorf("打开 error 日志文件失败: %w", err)
	}

	// 创建日志记录器
	infoLogger = log.New(infoFile, "[INFO] ", log.LstdFlags|log.Lshortfile)
	errorLogger = log.New(errorFile, "[ERROR] ", log.LstdFlags|log.Lshortfile)

	return nil
}

// openLogFile 打开日志文件（追加模式）
func openLogFile(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// getLogger 获取当前日期的日志记录器（自动切换日期）
func getLogger(level string) *log.Logger {
	logMutex.Lock()
	defer logMutex.Unlock()

	// 检查日期是否变化
	today := time.Now().Format("2006-01-02")
	if currentDate != today {
		updateLoggers()
	}

	if level == "error" {
		return errorLogger
	}
	return infoLogger
}

// Info 记录 Info 级别日志
func Info(format string, v ...interface{}) {
	logger := getLogger("info")
	if logger != nil {
		logger.Printf(format, v...)
	}
	// 同时输出到控制台
	log.Printf("[INFO] "+format, v...)
}

// LogError 记录 Error 级别日志
func LogError(format string, v ...interface{}) {
	logger := getLogger("error")
	if logger != nil {
		logger.Printf(format, v...)
	}
	// 同时输出到控制台
	log.Printf("[ERROR] "+format, v...)
}

// Warn 记录 Warn 级别日志（写入 info 文件）
func Warn(format string, v ...interface{}) {
	logger := getLogger("info")
	if logger != nil {
		logger.Printf("[WARN] "+format, v...)
	}
	// 同时输出到控制台
	log.Printf("[WARN] "+format, v...)
}

// Debug 记录 Debug 级别日志（根据配置决定是否记录）
func Debug(format string, v ...interface{}) {
	if config.Cfg != nil && config.Cfg.Log.Level == "debug" {
		logger := getLogger("info")
		if logger != nil {
			logger.Printf("[DEBUG] "+format, v...)
		}
		// 同时输出到控制台
		log.Printf("[DEBUG] "+format, v...)
	}
}

// GetLogWriter 获取日志写入器（用于 Gin）
func GetLogWriter() io.Writer {
	return &logWriter{}
}

// logWriter 实现 io.Writer 接口
type logWriter struct{}

func (w *logWriter) Write(p []byte) (n int, err error) {
	Info(string(p))
	return len(p), nil
}

// GetErrorLogWriter 获取错误日志写入器
func GetErrorLogWriter() io.Writer {
	return &errorLogWriter{}
}

// errorLogWriter 实现 io.Writer 接口
type errorLogWriter struct{}

func (w *errorLogWriter) Write(p []byte) (n int, err error) {
	LogError(string(p))
	return len(p), nil
}

