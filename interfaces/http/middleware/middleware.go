package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/domain/users/models"
	"github.com/riskykurniawan15/xarch/domain/users/services"
	"github.com/riskykurniawan15/xarch/helpers/jwt"
	"github.com/riskykurniawan15/xarch/interfaces/http/entities"
	"github.com/riskykurniawan15/xarch/logger"
)

type MiddlewareHandler struct {
	UserService *services.UserService
	cfg         config.Config
	log         logger.Logger
}

func NewMiddlewareHandlers(US *services.UserService, cfg config.Config, log logger.Logger) *MiddlewareHandler {
	return &MiddlewareHandler{
		US,
		cfg,
		log,
	}
}

func (MW MiddlewareHandler) LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		MW.log.InfoW("incomming request", map[string]interface{}{
			"at":     time.Now().Format("2006-01-02 15:04:05"),
			"method": ctx.Request().Method,
			"ip":     ctx.Request().RemoteAddr,
			"uri":    ctx.Request().URL.String(),
		})
		return next(ctx)
	}
}

func (MW MiddlewareHandler) UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		Auth := ctx.Request().Header.Values("Authorization")
		if len(Auth) != 1 {
			return ctx.JSON(http.StatusUnauthorized, entities.ResponseFormater(http.StatusUnauthorized, map[string]interface{}{
				"error": "No Authorization",
			}))
		}

		Token := strings.Split(Auth[0], "Bearer ")
		if len(Token) != 2 {
			return ctx.JSON(http.StatusUnauthorized, entities.ResponseFormater(http.StatusUnauthorized, map[string]interface{}{
				"error": "Malformed Format Token Bearer",
			}))
		}

		NewJwt := jwt.NewJwtToken(MW.cfg)
		claims, err := NewJwt.ClaimsToken(Token[1])
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, entities.ResponseFormater(http.StatusUnauthorized, map[string]interface{}{
				"error": err.Error(),
			}))
		}

		user, err_get := MW.UserService.GetDetailUser(ctx.Request().Context(), &models.User{ID: claims.ID})
		if err_get != nil {
			return ctx.JSON(http.StatusUnauthorized, entities.ResponseFormater(http.StatusUnauthorized, map[string]interface{}{
				"error": err_get.Errors.Error(),
			}))
		}

		ctx.Set("ID", user.ID)

		return next(ctx)
	}
}
