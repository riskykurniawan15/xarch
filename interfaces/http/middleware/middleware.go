package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/domain/users/models"
	"github.com/riskykurniawan15/xarch/domain/users/services"
	"github.com/riskykurniawan15/xarch/helpers/jwt"
	"github.com/riskykurniawan15/xarch/interfaces/http/entities"
)

type MiddlewareHandler struct {
	cfg         config.Config
	UserService *services.UserService
}

func NewMiddlewareHandlers(cfg config.Config, US *services.UserService) *MiddlewareHandler {
	return &MiddlewareHandler{
		cfg,
		US,
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

		user, err := MW.UserService.GetDetailUser(ctx.Request().Context(), &models.User{ID: claims.ID})
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, entities.ResponseFormater(http.StatusUnauthorized, map[string]interface{}{
				"error": err.Error(),
			}))
		}

		ctx.Set("ID", user.ID)

		return next(ctx)
	}
}
