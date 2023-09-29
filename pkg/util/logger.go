package util

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log zap.SugaredLogger
var RichLog zap.Logger

func InitLogger() {
	var cfg zap.Config
	if os.Getenv("ENVIRONMENT") == "prod" {
		cfg = zap.NewProductionConfig()
		cfg.Level.SetLevel(zap.InfoLevel)

	} else {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.DisableCaller = true
		cfg.Level.SetLevel(zap.DebugLevel)
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	Log = *logger.Sugar()
	RichLog = *logger
}
