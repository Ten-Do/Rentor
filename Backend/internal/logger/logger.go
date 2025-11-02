package logger

import (
	"go.uber.org/zap"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

// logger is a global zap.Logger instance
var logger *zap.Logger

// InitLogger initializes the global logger based on the environment
func InitLogger(env string) error {
	switch env {
	case EnvLocal, EnvDev:
		var err error
		logger, err = zap.NewDevelopment()
		if err != nil {
			return err
		}
	case EnvProd:
		var err error
		logger, err = zap.NewProduction()
		if err != nil {
			return err
		}
	}
	return nil
}

// Logging functions
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
func Sync() error {
	return logger.Sync()
}
func Field(key string, value any) zap.Field {
	return zap.Any(key, value)
}
