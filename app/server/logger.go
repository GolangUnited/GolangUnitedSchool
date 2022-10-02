package server

import (
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(logLevel, logEncoding string) *zap.SugaredLogger {
	logger, err := buildLoader(logLevel, logEncoding)
	if err != nil {
		log.Fatalf("can not build logger: %s", err)
	}

	return logger.Sugar()
}

func buildLoader(logLevel, logEncoding string) (*zap.Logger, error) {
	zapLevel := zap.NewAtomicLevel()
	if err := zapLevel.UnmarshalText([]byte(logLevel)); err != nil {
		return nil, fmt.Errorf("bad log level - '%s': %s", logLevel, err)
	}

	return zap.Config{
		Encoding:         logEncoding,
		Level:            zapLevel,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "message",
			LevelKey:      "level",
			TimeKey:       "time",
			CallerKey:     "caller",
			StacktraceKey: "stacktrace",
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}.Build()
}
