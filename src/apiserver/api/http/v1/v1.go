package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitListener(mx *echo.Echo) error {
	gr := mx.Group("/api/v1")

	gr.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	return nil
}
