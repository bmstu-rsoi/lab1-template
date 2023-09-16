package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InitListener(mx *echo.Echo) error {
	gr := mx.Group("/api/v1")

	gr.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	gr.POST("/persons", PostPerson)
	gr.GET("/persons", GetPersons)
	gr.GET("/persons/:id", GetPersons)
	gr.PATCH("/persons/:id", PatchPerson)
	gr.DELETE("/persons/:id", DeletePerson)

	return nil
}

type PersonRequset struct {
	Name    string `json:"name" validate:"required"`
	Age     int32  `json:"age"`
	Address string `json:"address"`
	Work    string `json:"work"`
}

type PersonResponse struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	Age     int32  `json:"age,omitempty"`
	Address string `json:"address,omitempty"`
	Work    string `json:"work,omitempty"`
}

type ValidationErrorResponse struct {
	Message string `json:"message"`
	Errors  string `json:"errors"`
}

func PostPerson(c echo.Context) error {
	var req PersonRequset
	err := c.Bind(&req); if err != nil {
		// return c.JSON(http.StatusBadRequest, ValidationErrorResponse{})
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	info := PersonResponse{}

	return c.JSON(http.StatusCreated, info)
}

func GetPersons(c echo.Context) error {
	infos := make([]PersonResponse, 1)

	return c.JSON(http.StatusOK, infos)
}

func GetPerson(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	id := int32(id64)

	info := PersonResponse{ID: id}

	return c.JSON(http.StatusOK, info)
}

func PatchPerson(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	id := int32(id64)

	var req PersonRequset
	if err = c.Bind(&req); err != nil {
		// return c.JSON(http.StatusBadRequest, ValidationErrorResponse{})
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	info := PersonResponse{ID: id}

	return c.JSON(http.StatusOK, info)
}

func DeletePerson(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	id := int32(id64)

	var req PersonRequset
	if err = c.Bind(&req); err != nil {
		// return c.JSON(http.StatusBadRequest, ValidationErrorResponse{})
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	info := PersonResponse{ID: id}

	return c.JSON(http.StatusOK, info)
}
