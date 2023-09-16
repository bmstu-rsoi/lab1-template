package http

import (
	"fmt"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/migregal/bmstu-iu7-ds-lab1/apiserver/api/http/common"
	v1 "github.com/migregal/bmstu-iu7-ds-lab1/apiserver/api/http/v1"
	"github.com/migregal/bmstu-iu7-ds-lab1/pkg/readiness"
)

type Server struct {
	mx *echo.Echo
}

func New(lg *slog.Logger, probe *readiness.Probe) (*Server, error) {
	mx := echo.New()
	mx.Use(
		middleware.Recover(),
		middleware.Logger(),
		middleware.RequestID(),
	)
	mx.Debug = false
	// mx.HideBanner = true
	// mx.HidePort = true
	mx.HTTPErrorHandler = func(err error, c echo.Context) {
		// Take required information from error and context and send it to a service like New Relic
		// fmt.Println(c.Path(), c.QueryParams(), err.Error())

		mx.DefaultHTTPErrorHandler(err, c)
	}

	s := Server{mx : mx}

	err := common.InitListener(s.mx, probe)
	if err != nil {
		return nil, fmt.Errorf("failed to init common apis: %w", err)
	}
	err = v1.InitListener(s.mx)
	if err != nil {
		return nil, fmt.Errorf("failed to init v1 apis: %w", err)
	}

	return &s, nil
}

func (s *Server) ListenAndServe(addr string) error {
	return s.mx.Start(addr)
}
