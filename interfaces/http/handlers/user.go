package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (handler UserHandler) ShowUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "oke",
	})
}
