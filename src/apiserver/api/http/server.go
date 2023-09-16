package http

import (
	"fmt"
	"net/http"

	v1 "github.com/migregal/bmstu-iu7-ds-lab1/apiserver/api/http/v1"
)

type Server struct {
	mx *http.ServeMux
}

func New() (*Server, error) {
	s := Server{}

	s.mx = http.NewServeMux()

	err := v1.InitListener(s.mx)
	if err != nil {
		return nil, fmt.Errorf("failed to init v1 apis: %w", err)
	}

	return &s, nil
}

func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s.mx)
}
