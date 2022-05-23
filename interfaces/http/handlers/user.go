package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/riskykurniawan15/xarch/domain/users/services"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandlers(US *services.UserService) *UserHandler {
	return &UserHandler{
		US,
	}
}

func (handler UserHandler) Index(c echo.Context) error {
	User, err := handler.UserService.GetListUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status": "504",
			"error":  fmt.Sprint(err),
		})
	}

	return c.JSON(http.StatusOK, User)
}

func (handler UserHandler) Show(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]string{
		"status": "404",
	})
}

func (handler UserHandler) Store(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]string{
		"status": "404",
	})
}

func (handler UserHandler) Update(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]string{
		"status": "404",
	})
}

func (handler UserHandler) Delete(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]string{
		"status": "404",
	})
}
