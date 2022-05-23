package routers

import (
	"github.com/labstack/echo/v4"
	gate "github.com/riskykurniawan15/xarch/domain"
	user "github.com/riskykurniawan15/xarch/interfaces/http/handlers"
)

type Handlers struct {
	UserHandler *user.UserHandler
}

func StartHandlers(svc *gate.Service) *Handlers {
	users := user.NewUserHandlers(svc.UserService)

	return &Handlers{
		users,
	}
}

func Routers(svc *gate.Service) *echo.Echo {
	handler := StartHandlers(svc)

	e := echo.New()
	user_group := e.Group("/user")
	{
		user_handler := handler.UserHandler
		user_group.GET("/", user_handler.Index)
	}

	return e
}
