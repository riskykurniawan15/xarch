package routers

import (
	"github.com/labstack/echo/v4"
	user "github.com/riskykurniawan15/xarch/interfaces/http/handlers"
)

func Routers() *echo.Echo {
	users := user.UserHandler{}

	e := echo.New()
	e.GET("/", users.ShowUser)

	return e
}
