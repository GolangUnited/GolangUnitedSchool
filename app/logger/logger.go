package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Newlogger(level, encoding string) (*zap.SugaredLogger, error) {
	zapLevel := zap.NewAtomicLevel()
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		return nil, fmt.Errorf("zap logger unmarshal log level: %v", err)
	}

	zapCfg := zap.Config{
		Encoding:         encoding,
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
		}}

	logger, err := zapCfg.Build()
	if err != nil {
		return nil, fmt.Errorf("zap logger build: %v", err)
	}

	return logger.Sugar(), nil
}
