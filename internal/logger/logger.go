package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger

// Init...
func Init(core zapcore.Core, options ...zap.Option) {
	globalLogger = zap.New(core, options...)
}

// Init...
func Debug(msg string, fields ...zap.Field) {
	globalLogger.Debug(msg, fields...)
}

// Info...
func Info(msg string, fields ...zap.Field) {
	globalLogger.Info(msg, fields...)
}

// Warn...
func Warn(msg string, fields ...zap.Field) {
	globalLogger.Warn(msg, fields...)
}

// Error...
func Error(msg string, fields ...zap.Field) {
	globalLogger.Error(msg, fields...)
}

// Fatal...
func Fatal(msg string, fields ...zap.Field) {
	globalLogger.Fatal(msg, fields...)
}

// WithOptions...
func WithOptions(opts ...zap.Option) *zap.Logger {
	return globalLogger.WithOptions(opts...)
}
