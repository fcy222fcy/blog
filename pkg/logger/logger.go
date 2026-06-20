package logger

import (
	"os"

	"blog/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger 日志
type Logger struct {
	*zap.SugaredLogger
}

var globalLogger *Logger

// NewLogger 创建日志实例
func NewLogger(cfg config.LogConfig) (*Logger, error) {
	// 日志级别
	level := zap.InfoLevel
	switch cfg.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	}

	// 日志轮转配置
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,    // MB
		MaxBackups: cfg.MaxBackups, // 个
		MaxAge:     cfg.MaxAge,     // 天
		Compress:   true,
	})

	// 编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// 创建 Core
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writeSyncer, level),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level),
	)

	// 创建 Logger
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &Logger{zapLogger.Sugar()}, nil
}

// SetGlobalLogger 设置全局日志
func SetGlobalLogger(l *Logger) {
	globalLogger = l
}

// GetGlobalLogger 获取全局日志
func GetGlobalLogger() *Logger {
	return globalLogger
}

// Debug 调试日志
func Debug(args ...interface{}) {
	globalLogger.Debug(args...)
}

// Debugf 调试日志（格式化）
func Debugf(template string, args ...interface{}) {
	globalLogger.Debugf(template, args...)
}

// Info 信息日志
func Info(args ...interface{}) {
	globalLogger.Info(args...)
}

// Infof 信息日志（格式化）
func Infof(template string, args ...interface{}) {
	globalLogger.Infof(template, args...)
}

// Warn 警告日志
func Warn(args ...interface{}) {
	globalLogger.Warn(args...)
}

// Warnf 警告日志（格式化）
func Warnf(template string, args ...interface{}) {
	globalLogger.Warnf(template, args...)
}

// Error 错误日志
func Error(args ...interface{}) {
	globalLogger.Error(args...)
}

// Errorf 错误日志（格式化）
func Errorf(template string, args ...interface{}) {
	globalLogger.Errorf(template, args...)
}

// Fatal 致命错误日志
func Fatal(args ...interface{}) {
	globalLogger.Fatal(args...)
}

// Fatalf 致命错误日志（格式化）
func Fatalf(template string, args ...interface{}) {
	globalLogger.Fatalf(template, args...)
}

// Sync 同步日志
func (l *Logger) Sync() error {
	return l.Sync()
}
