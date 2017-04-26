package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a global that is used in logging middleware
var Logger *zap.Logger

// LogLevel defines current logging level: debug, info, warn, etc
var LogLevel zapcore.Level

// InitLogger creates logger and configures it
func InitLogger() {
	switch GetLogLevel() {
	case "debug":
		LogLevel = zap.DebugLevel
	case "info":
		LogLevel = zap.InfoLevel
	case "warn":
		LogLevel = zap.WarnLevel
	case "error":
		LogLevel = zap.ErrorLevel
	default:
		LogLevel = zap.InfoLevel
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(LogLevel)
	config := zap.NewProductionConfig()
	config.OutputPaths = GetLogOutput()
	config.Level = atomicLevel
	config.DisableCaller = !ShouldLogCaller()
	config.DisableStacktrace = !ShouldLogStacktrace()

	Logger, _ = config.Build()
}
