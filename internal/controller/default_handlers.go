package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func indexHandler(c echo.Context) error {
	return c.JSON(http.StatusOK,
		map[string]string{
			"service_name": "template",
		})
}

func checkHealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
