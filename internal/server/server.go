package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/Astemirdum/person-service/config"
)

type Server struct {
	s *http.Server
}

func NewServer(cfg config.HTTPServer, router *echo.Echo) *Server {
	s := &http.Server{
		Addr:              net.JoinHostPort(cfg.Host, cfg.Port),
		Handler:           router,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		ReadHeaderTimeout: time.Second * 5,
		MaxHeaderBytes:    8 * 1024,
	}
	return &Server{s: s}
}

func (s *Server) Run() error {
	return s.s.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.s.Shutdown(ctx)
}
