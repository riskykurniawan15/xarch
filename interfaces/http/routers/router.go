package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/domain"
	handlers "github.com/riskykurniawan15/xarch/interfaces/http/handlers"
	"github.com/riskykurniawan15/xarch/interfaces/http/middleware"
)

type Handlers struct {
	MiddlewareHandler *middleware.MiddlewareHandler
	UserHandler       *handlers.UserHandler
	AlquranHandler    *handlers.AlquranHandler
}

func StartHandlers(svc *domain.Service, cfg config.Config) *Handlers {
	middleware := middleware.NewMiddlewareHandlers(cfg, svc.UserService)
	users := handlers.NewUserHandlers(svc.UserService)
	alquran := handlers.NewAlquranHandlers(svc.AlquranService)

	return &Handlers{
		middleware,
		users,
		alquran,
	}
}

func Routers(svc *domain.Service, cfg config.Config) *echo.Echo {
	handler := StartHandlers(svc, cfg)
	middleware := handler.MiddlewareHandler
	user_handler := handler.UserHandler
	alquran_handler := handler.AlquranHandler

	engine := echo.New()
	engine.POST("/register", user_handler.Register)
	engine.POST("/login", user_handler.Login)
	engine.POST("/forgot", user_handler.ForgotPass)

	engine.POST("/user/verif/resend", user_handler.ReSendVerification)
	engine.GET("/user/verif/:id/:token", user_handler.Verification)
	user_group := engine.Group("/user")
	{
		user_group.Use(middleware.UserMiddleware)
		user_group.GET("/profile", user_handler.GetProfile)
		user_group.PUT("/profile", user_handler.UpdateProfile)
		user_group.POST("/password", user_handler.UpdatePassword)
	}

	alquran_group := engine.Group("/alquran")
	{
		alquran_group.GET("/chapter", alquran_handler.ListChapter)
		alquran_group.GET("/chapter/:ID", alquran_handler.DetailChapter)
		alquran_group.GET("/chapter/:ID/verse", alquran_handler.ListVerse)
		alquran_group.GET("/chapter/:ID/verse/:verse_number", alquran_handler.DetailVerse)
	}

	return engine
}
