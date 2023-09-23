package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	LogLevel zapcore.Level `yaml:"level" envconfig:"LOG_LEVEL"`
	Sink     string        `yaml:"sink" envconfig:"LOG_SINK"` // TODO: stdOut or file
}

func NewLogger(c Log, ns string) *zap.Logger {
	conf := zap.NewDevelopmentEncoderConfig()
	conf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(conf),
		zapcore.AddSync(os.Stdout),
		c.LogLevel,
	))
	return logger.Named(ns)
}
