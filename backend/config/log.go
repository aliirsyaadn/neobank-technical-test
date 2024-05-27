package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	config.DisableStacktrace = true
	logger, _ := config.Build()
	defer logger.Sync()

	sugar := logger.Sugar()
	return sugar
}
