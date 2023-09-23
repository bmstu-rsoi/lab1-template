package main

import (
	"context"
	stdLog "log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/Astemirdum/person-service/internal/handler"
	"github.com/Astemirdum/person-service/internal/repository"
	"github.com/Astemirdum/person-service/internal/server"
	"github.com/Astemirdum/person-service/internal/service"
	"github.com/Astemirdum/person-service/pkg/logger"
	"github.com/Astemirdum/person-service/pkg/postgres"

	"github.com/Astemirdum/person-service/config"
	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
)

func main() {
	if err := godotenv.Load(); err != nil {
		stdLog.Fatal("load envs from .env ", zap.Error(err))
	}
	cfg := config.NewConfig(
		config.WithLogLevel(zapcore.DebugLevel),
		config.WithWriteTimeout(time.Minute),
	)

	run(cfg)
}

func run(cfg *config.Config) {
	log := logger.NewLogger(cfg.Log, "person")
	db, err := postgres.NewPostgresDB(&cfg.Database)
	if err != nil {
		log.Fatal("db init", zap.Error(err))
	}
	repo, err := repository.NewRepository(db, log)
	if err != nil {
		log.Fatal("repo users", zap.Error(err))
	}
	userService := service.NewService(repo, log)

	h := handler.New(userService, log)

	srv := server.NewServer(cfg.Server, h.NewRouter())
	log.Info("http server start ON: ",
		zap.String("addr",
			net.JoinHostPort(cfg.Server.Host, cfg.Server.Port)))
	go func() {
		if err := srv.Run(); err != nil {
			log.Error("server run", zap.Error(err))
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	termSig := <-sig

	log.Debug("Graceful shutdown", zap.Any("signal", termSig))

	closeCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err = srv.Stop(closeCtx); err != nil {
		log.DPanic("srv.Stop", zap.Error(err))
	}
	if err = db.Close(); err != nil {
		log.DPanic(" db.Close()", zap.Error(err))
	}
	log.Info("Graceful shutdown finished")
}
