package handler

import (
	"github.com/Astemirdum/person-service/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/time/rate"
)

func newRateLimiterMW(rps rate.Limit) echo.MiddlewareFunc {
	return middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rps))
}

func requestLoggerConfig() middleware.RequestLoggerConfig {
	cfg := logger.Log{LogLevel: zapcore.DebugLevel, Sink: ""}
	log := logger.NewLogger(cfg, "echo")
	c := middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		HandleError:  true,
		LogError:     true,
		LogLatency:   true,
		LogRequestID: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			level := zapcore.InfoLevel
			if v.Error != nil {
				level = zapcore.ErrorLevel
			}
			log.Log(level, "request",
				zap.String("URI", v.URI),
				zap.String("Method", v.Method),
				zap.Int("status", v.Status),
				zap.Duration("latency", v.Latency),
				zap.Error(v.Error),
				zap.String("request_id", v.RequestID),
			)
			return nil
		},
	}
	return c
}
