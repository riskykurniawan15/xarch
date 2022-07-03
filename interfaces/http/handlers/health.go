package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/riskykurniawan15/xarch/domain/health/services"
	"github.com/riskykurniawan15/xarch/interfaces/http/entities"
)

type HealthHandler struct {
	HealthService *services.HealthService
}

func NewHealthHandlers(US *services.HealthService) *HealthHandler {
	return &HealthHandler{
		US,
	}
}

func (handler HealthHandler) Metric(ctx echo.Context) error {
	metric := handler.HealthService.HealthMetric(ctx.Request().Context())
	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": metric,
	}))
}
