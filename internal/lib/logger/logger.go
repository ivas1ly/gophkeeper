package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(level string, cfg zap.Config) *zap.Logger {
	var logLevel zapcore.Level

	switch strings.ToLower(level) {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	case "fatal":
		logLevel = zap.FatalLevel
	default:
		logLevel = zap.InfoLevel
	}

	zapConfig := cfg
	zapConfig.Level = zap.NewAtomicLevelAt(logLevel)

	logger := zap.Must(zapConfig.Build())

	return logger
}

func NewDefaultLoggerConfig() zap.Config {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	loggerCfg := zap.Config{
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stdout"},
		InitialFields:     nil,
	}

	return loggerCfg
}

func SetGlobalLogger(log *zap.Logger) {
	zap.ReplaceGlobals(log)
}
