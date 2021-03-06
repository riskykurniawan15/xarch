package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/domain"
	handlers "github.com/riskykurniawan15/xarch/interfaces/xarch/http/handlers"
	"github.com/riskykurniawan15/xarch/interfaces/xarch/http/middleware"
	"github.com/riskykurniawan15/xarch/logger"
)

type Handlers struct {
	MiddlewareHandler *middleware.MiddlewareHandler
	HealthHandler     *handlers.HealthHandler
	UserHandler       *handlers.UserHandler
	AlquranHandler    *handlers.AlquranHandler
}

func StartHandlers(svc *domain.Service, cfg config.Config, log logger.Logger) *Handlers {
	middleware := middleware.NewMiddlewareHandlers(svc.UserService, cfg, log)
	health := handlers.NewHealthHandlers(svc.HealthService)
	users := handlers.NewUserHandlers(svc.UserService)
	alquran := handlers.NewAlquranHandlers(svc.AlquranService)

	return &Handlers{
		middleware,
		health,
		users,
		alquran,
	}
}

func Routers(svc *domain.Service, cfg config.Config, log logger.Logger) *echo.Echo {
	handler := StartHandlers(svc, cfg, log)
	middleware := handler.MiddlewareHandler
	health_handler := handler.HealthHandler
	user_handler := handler.UserHandler
	alquran_handler := handler.AlquranHandler

	engine := echo.New()
	engine.Use(middleware.LoggerMiddleware)

	engine.GET("/health", health_handler.Metric)

	engine.POST("/register", user_handler.Register)
	engine.POST("/login", user_handler.Login)
	engine.POST("/forgot", user_handler.ForgotPass)
	engine.POST("/reset/pass", user_handler.ResetPass)

	engine.POST("/user/verif/resend", user_handler.ReSendVerification)
	engine.GET("/user/verif/:id/:token", user_handler.Verification)
	user_group := engine.Group("/user")
	{
		user_group.Use(middleware.UserMiddleware)
		user_group.GET("/profile", user_handler.GetProfile)
		user_group.PUT("/profile", user_handler.UpdateProfile)
		user_group.POST("/password", user_handler.UpdatePassword)
		user_group.GET("/profile_image", user_handler.GetProfileImage)
		user_group.PUT("/profile_image", user_handler.UploadProfileImage)
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
