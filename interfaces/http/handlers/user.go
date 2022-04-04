package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/riskykurniawan15/xarch/services"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandlers(US *services.UserService) *UserHandler {
	return &UserHandler{
		US,
	}
}

func (handler UserHandler) ShowUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "oke",
	})
}
