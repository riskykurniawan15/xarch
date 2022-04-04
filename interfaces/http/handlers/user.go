package handlers

import (
	"fmt"
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
	User, err := handler.UserService.GetListUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status": "504",
			"error":  fmt.Sprint(err),
		})
	}

	return c.JSON(http.StatusOK, User)
}
