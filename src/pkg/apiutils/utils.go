package apiutils

import (
	"log/slog"
	"os"
)


type Serveable interface {
	ListenAndServe(adr string) error
}

type Callable func(errs chan<- error)

func NewCallable(addr string, api Serveable) Callable {
	return func(errs chan<- error) {
		errs <- api.ListenAndServe(addr)
	}
}

func Serve(lg *slog.Logger, apis ...Callable) {
	errs := make(chan error, len(apis))

	for _, api := range apis {
		go api(errs)
	}

	select {
	case err := <-errs:
		lg.Error("[shutdown] terminating application: %w", err)
		os.Exit(1)
	}
}
