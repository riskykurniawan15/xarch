package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/riskykurniawan15/xarch/domain"
	handlers "github.com/riskykurniawan15/xarch/interfaces/http/handlers"
)

type Handlers struct {
	UserHandler    *handlers.UserHandler
	AlquranHandler *handlers.AlquranHandler
}

func StartHandlers(svc *domain.Service) *Handlers {
	users := handlers.NewUserHandlers(svc.UserService)
	alquran := handlers.NewAlquranHandlers(svc.AlquranService)

	return &Handlers{
		users,
		alquran,
	}
}

func Routers(svc *domain.Service) *echo.Echo {
	handler := StartHandlers(svc)
	user_handler := handler.UserHandler
	alquran_handler := handler.AlquranHandler

	e := echo.New()
	user_group := e.Group("/user")
	{
		user_group.GET("/", user_handler.Index)
	}

	alquran_group := e.Group("/alquran")
	{
		alquran_group.GET("/chapter", alquran_handler.ListChapter)
		alquran_group.GET("/chapter/:ID", alquran_handler.DetailChapter)
	}

	return e
}
