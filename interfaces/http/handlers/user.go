package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/riskykurniawan15/xarch/domain/users/models"
	"github.com/riskykurniawan15/xarch/domain/users/services"
	"github.com/riskykurniawan15/xarch/interfaces/http/entities"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandlers(US *services.UserService) *UserHandler {
	return &UserHandler{
		US,
	}
}

func (handler UserHandler) Register(ctx echo.Context) error {
	form := new(models.UserRegisterForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		}))
	}

	data, err := handler.UserService.RegisterUser(ctx.Request().Context(), form)
	if err != nil {
		return ctx.JSON(err.HttpCode, entities.ResponseFormater(err.HttpCode, map[string]interface{}{
			"error": err.Errors.Error(),
		}))
	}

	return ctx.JSON(http.StatusCreated, entities.ResponseFormater(http.StatusCreated, map[string]interface{}{
		"data": data,
	}))
}

func (handler UserHandler) ReSendVerification(ctx echo.Context) error {
	form := new(models.ReSendVerificationForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		}))
	}

	data, err := handler.UserService.ReSendEmailVerification(ctx.Request().Context(), form)
	if err != nil {
		return ctx.JSON(err.HttpCode, entities.ResponseFormater(err.HttpCode, map[string]interface{}{
			"error": err.Errors.Error(),
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}

func (handler UserHandler) Login(ctx echo.Context) error {
	form := new(models.UserLoginForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		}))
	}

	data, err := handler.UserService.LoginUser(ctx.Request().Context(), form)
	if err != nil {
		return ctx.JSON(err.HttpCode, entities.ResponseFormater(err.HttpCode, map[string]interface{}{
			"error": err.Errors.Error(),
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}

func (handler UserHandler) ForgotPass(ctx echo.Context) error {
	form := new(models.ForgotPassForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		}))
	}

	data, err := handler.UserService.ForgotPassword(ctx.Request().Context(), form)
	if err != nil {
		return ctx.JSON(err.HttpCode, entities.ResponseFormater(err.HttpCode, map[string]interface{}{
			"error": err.Errors.Error(),
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}

func (handler UserHandler) ResetPass(ctx echo.Context) error {
	form := new(models.ResetPassForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		}))
	}

	data, err := handler.UserService.ResetPass(ctx.Request().Context(), form)
	if err != nil {
		return ctx.JSON(err.HttpCode, entities.ResponseFormater(err.HttpCode, map[string]interface{}{
			"error": err.Errors.Error(),
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}

func (handler UserHandler) UpdateProfile(ctx echo.Context) error {
	ID, err_id := strconv.Atoi(fmt.Sprint(ctx.Get("ID")))
	if err_id != nil {
		return ctx.JSON(http.StatusBadGateway, entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err_id.Error(),
		}))
	}

	form := new(models.UserUpdateProfileForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		}))
	}

	data, err := handler.UserService.UpdateProfileUser(ctx.Request().Context(), uint(ID), form)
	if err != nil {
		return ctx.JSON(err.HttpCode, entities.ResponseFormater(err.HttpCode, map[string]interface{}{
			"error": err.Errors.Error(),
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}

func (handler UserHandler) UpdatePassword(ctx echo.Context) error {
	ID, err_id := strconv.Atoi(fmt.Sprint(ctx.Get("ID")))
	if err_id != nil {
		return ctx.JSON(http.StatusBadGateway, entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err_id.Error(),
		}))
	}

	form := new(models.UpdatePasswordForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		}))
	}

	data, err := handler.UserService.UpdatePasswordUser(ctx.Request().Context(), uint(ID), form)
	if err != nil {
		return ctx.JSON(err.HttpCode, entities.ResponseFormater(err.HttpCode, map[string]interface{}{
			"error": err.Errors.Error(),
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}

func (handler UserHandler) GetProfile(ctx echo.Context) error {
	ID, err := strconv.Atoi(fmt.Sprint(ctx.Get("ID")))
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err.Error(),
		}))
	}

	data, err := handler.UserService.GetDetailUser(ctx.Request().Context(), &models.User{ID: uint(ID)})
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err.Error(),
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}

func (handler UserHandler) Verification(ctx echo.Context) error {
	ID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err.Error(),
		}))
	}
	token := ctx.Param("token")

	data, err := handler.UserService.EmailVerification(ctx.Request().Context(), uint(ID), token)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err.Error(),
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}
