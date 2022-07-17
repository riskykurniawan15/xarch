package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/riskykurniawan15/xarch/domain/alquran/services"
	"github.com/riskykurniawan15/xarch/interfaces/xarch/http/entities"
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
		return ctx.JSON(err.HttpCode, entities.ResponseFormater(err.HttpCode, map[string]interface{}{
			"error": err.Errors.Error(),
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}

func (handler AlquranHandler) DetailChapter(ctx echo.Context) error {
	ID, err_param := strconv.Atoi(ctx.Param("ID"))
	if err_param != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err_param,
		}))
	}

	data, err := handler.AlquranService.GetDetailChapter(ctx.Request().Context(), ID)
	if err != nil {
		return ctx.JSON(err.HttpCode, entities.ResponseFormater(err.HttpCode, map[string]interface{}{
			"error": err.Errors.Error(),
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}

func (handler AlquranHandler) ListVerse(ctx echo.Context) error {
	ID, err_param := strconv.Atoi(ctx.Param("ID"))
	if err_param != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err_param,
		}))
	}

	data, err := handler.AlquranService.GetListVerse(ctx.Request().Context(), ID)
	if err != nil {
		return ctx.JSON(err.HttpCode, entities.ResponseFormater(err.HttpCode, map[string]interface{}{
			"error": err.Errors.Error(),
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}

func (handler AlquranHandler) DetailVerse(ctx echo.Context) error {
	ID, err_param := strconv.Atoi(ctx.Param("ID"))
	if err_param != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err_param,
		}))
	}

	VerseNumber, err_param := strconv.Atoi(ctx.Param("verse_number"))
	if err_param != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err_param,
		}))
	}

	data, err := handler.AlquranService.GetDetailVerse(ctx.Request().Context(), ID, VerseNumber)
	if err != nil {
		return ctx.JSON(err.HttpCode, entities.ResponseFormater(err.HttpCode, map[string]interface{}{
			"error": err.Errors.Error(),
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}
