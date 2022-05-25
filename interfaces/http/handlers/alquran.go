package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/riskykurniawan15/xarch/domain/alquran/services"
	"github.com/riskykurniawan15/xarch/interfaces/http/entities"
)

type AlquranHandler struct {
	AlquranService *services.AlquranService
}

func NewAlquranHandlers(AS *services.AlquranService) *AlquranHandler {
	return &AlquranHandler{
		AS,
	}
}

func (handler AlquranHandler) ListChapter(ctx echo.Context) error {
	data, err := handler.AlquranService.GetListChapter(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err,
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}

func (handler AlquranHandler) DetailChapter(ctx echo.Context) error {
	ID, err := strconv.Atoi(ctx.Param("ID"))

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		}))
	}

	if ID < 1 || ID > 114 {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": "Chapter min 1 and max 114",
		}))
	}

	data, err := handler.AlquranService.GetDetailChapter(ctx.Request().Context(), ID)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err,
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}
