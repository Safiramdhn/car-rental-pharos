package log

import (
	"car-rental/config"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(cfg config.Config) (*zap.Logger, error) {
	logLevel := zap.InfoLevel
	if cfg.AppDebug {
		logLevel = zap.DebugLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		MessageKey:     "M",
		CallerKey:      "C",
		StacktraceKey:  "S",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	file, err := os.Create("app.log")
	if err != nil {
		log.Panicf("Failed to create log file: %v", err)
		return nil, err
	}

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(os.Stdout), logLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(file), logLevel),
	)

	logger := zap.New(core, zap.AddCaller())
	return logger, nil
}
