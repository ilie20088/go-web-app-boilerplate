package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var LogLevel zapcore.Level

func InitLogger() {

	switch ConfigManager.GetString("log.level") {
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
	config.OutputPaths = ConfigManager.GetStringSlice("log.output")
	config.Level = atomicLevel

	config.DisableCaller = !ConfigManager.GetBool("log.caller")
	config.DisableStacktrace = !ConfigManager.GetBool("log.stacktrace")

	Logger, _ = config.Build()
}
