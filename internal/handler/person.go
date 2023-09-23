package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Astemirdum/person-service/internal/errs"

	"github.com/labstack/echo/v4"

	personModel "github.com/Astemirdum/person-service/internal/model"
)

type UpdatePersonReq struct {
	Name    string `db:"name"`
	Surname string `db:"surname"`
}

func (h *Handler) ListPerson(c echo.Context) error {
	ctx := c.Request().Context()

	persons, err := h.personSvc.List(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set("Content-Type", echo.MIMEApplicationJSON)
	// h.log.Info("h.personSvc.Get", zap.Any("person", person))

	return c.JSON(http.StatusOK, persons)
}

func (h *Handler) CreatePerson(c echo.Context) error {
	var pers personModel.Person
	if err := c.Bind(&pers); err != nil {
		return httpValidationError(c, err)
	}
	if err := c.Validate(pers); err != nil {
		return httpValidationError(c, err)
	}
	ctx := c.Request().Context()
	if err := h.personSvc.Create(ctx, pers); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set("Location", "")

	return c.String(http.StatusOK, "OK")
}

func (h *Handler) GetPerson(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	person, err := h.personSvc.Get(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, person)
}

func (h *Handler) DeletePerson(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.personSvc.Delete(ctx, id); err != nil && !errors.Is(err, errs.ErrNotFound) {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusNoContent, "ok")
}

func (h *Handler) UpdatePerson(c echo.Context) error {

	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var pers personModel.Person
	if err := c.Bind(&pers); err != nil {
		return httpValidationError(c, err)
	}
	if err := c.Validate(pers); err != nil {
		return httpValidationError(c, err)
	}
	pers.ID = id
	resp, err := h.personSvc.Update(ctx, pers)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func httpValidationError(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, errs.ValidationErrorResponse{
		Message: err.Error(),
		Errors: struct {
			AdditionalProperties string `json:"additionalProperties"`
		}{
			AdditionalProperties: "",
		},
	})
}
