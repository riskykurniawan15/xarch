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
	var responseData entities.Response

	data, err := handler.AlquranService.GetListChapter(ctx.Request().Context())
	if err != nil {
		responseData = entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err,
		})
	} else {
		responseData = entities.ResponseFormater(http.StatusOK, map[string]interface{}{
			"data": data,
		})
	}

	return ctx.JSON(responseData.Status, responseData)
}

func (handler AlquranHandler) DetailChapter(ctx echo.Context) error {
	var responseData entities.Response

	ID, err := strconv.Atoi(ctx.Param("ID"))

	if err != nil {
		responseData = entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	if ID < 1 || ID > 114 {
		responseData = entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": "Chapter min 1 and max 114",
		})
	}

	data, err := handler.AlquranService.GetDetailChapter(ctx.Request().Context(), ID)
	if err != nil {
		responseData = entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err,
		})
	} else {
		responseData = entities.ResponseFormater(http.StatusOK, map[string]interface{}{
			"data": data,
		})
	}

	return ctx.JSON(responseData.Status, responseData)
}
