package main

import (
	"log/slog"
	"os"

	"github.com/migregal/bmstu-iu7-ds-lab1/apiserver"
	"github.com/migregal/bmstu-iu7-ds-lab1/apiserver/config"
)

func main() {
	cfg := config.ReadConfig()

	lg := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app, err := apiserver.New(lg, cfg)
	if err != nil {
		lg.Error("[startup] failed to init app: %w", err)
		os.Exit(1)
	}

	app.Run()
}
