package config

import (
	"time"

	"go.uber.org/zap/zapcore"
)

type Option func(config *Config)

func WithWriteTimeout(dur time.Duration) Option {
	return func(cfg *Config) {
		cfg.Server.WriteTimeout = dur
	}
}

func WithLogLevel(level zapcore.Level) Option {
	return func(cfg *Config) {
		cfg.Log.LogLevel = level
	}
}
