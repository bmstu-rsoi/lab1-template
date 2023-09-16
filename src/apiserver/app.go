package apiserver

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/migregal/bmstu-iu7-ds-lab1/apiserver/api/http"
	"github.com/migregal/bmstu-iu7-ds-lab1/apiserver/config"
	"github.com/migregal/bmstu-iu7-ds-lab1/pkg/apiutils"
)

type App struct {
	cfg *config.Config

	http *http.Server
}

func New(cfg *config.Config) (*App, error) {
	a := App{cfg: cfg}

	var err error
	a.http, err = http.New()
	if err != nil {
		return nil, fmt.Errorf("failed to init http server: %w", err)
	}

	return &a, nil
}

func (s *App) Run() {lg := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	apiutils.Serve(lg,
		apiutils.NewCallable(s.cfg.HttpAddr, s.http),
	)
}
