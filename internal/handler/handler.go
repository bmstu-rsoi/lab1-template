package handler

import (
	"net/http"
	"time"

	"github.com/Astemirdum/person-service/pkg/validate"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"

	_ "github.com/Astemirdum/person-service/swagger"
)

type Handler struct {
	personSvc PersonService
	client    *http.Client
	log       *zap.Logger
}

func New(personSrv PersonService, log *zap.Logger) *Handler {
	h := &Handler{
		personSvc: personSrv,
		log:       log,
		client: &http.Client{
			Timeout: time.Minute,
		},
	}
	return h
}

func (h *Handler) NewRouter() *echo.Echo {
	e := echo.New()
	const (
		baseRPS = 10
		apiRPS  = 100
	)
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 4 << 10, // 4 KB
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodOptions, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	base := e.Group("", newRateLimiterMW(baseRPS))
	base.GET("/health", h.Health)
	base.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Validator = validate.NewCustomValidator()
	api := e.Group("/api/v1",
		middleware.RequestLoggerWithConfig(requestLoggerConfig()),
		middleware.RequestID(),
		newRateLimiterMW(apiRPS),
	)

	api.GET("/persons", h.ListPerson)
	api.POST("/persons", h.CreatePerson)

	api.GET("/persons/:id", h.GetPerson)
	api.DELETE("/persons/:id", h.DeletePerson)
	api.PATCH("/persons/:id", h.UpdatePerson)

	return e
}

func (h *Handler) Health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
