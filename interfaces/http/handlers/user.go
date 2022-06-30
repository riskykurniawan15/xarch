package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
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

var (
	validate *validator.Validate = validator.New()
)

type UserRegisterForm struct {
	Name     string `form:"name"     validate:"required,max=100"                  json:"name"`
	Email    string `form:"email"    validate:"required,max=100,email"            json:"email"`
	Password string `form:"password" validate:"required,max=100"                  json:"-"`
	Confirm  string `form:"confirm"  validate:"required,max=100,eqfield=Password" json:"-"`
	Gender   string `form:"gender"   validate:"required,oneof=male female"        json:"gender"`
}

func (handler UserHandler) Register(ctx echo.Context) error {
	form := new(UserRegisterForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		}))
	}

	err := validate.Struct(form)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		}))
	}

	pyld := models.SetUserData(form.Name, form.Email, form.Password, form.Gender)

	data, err := handler.UserService.RegisterUser(ctx.Request().Context(), pyld)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err,
		}))
	}

	return ctx.JSON(http.StatusCreated, entities.ResponseFormater(http.StatusCreated, map[string]interface{}{
		"data": data,
	}))
}

type ReSendVerificationForm struct {
	Email string `form:"email"    validate:"required,max=100,email"            json:"email"`
}

func (handler UserHandler) ReSendVerification(ctx echo.Context) error {
	form := new(ReSendVerificationForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		}))
	}

	err := validate.Struct(form)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		}))
	}

	data, err := handler.UserService.ReSendEmailVerification(ctx.Request().Context(), form.Email)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err.Error(),
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}

type UserLoginForm struct {
	Email    string `form:"email"    validate:"required,max=100,email" json:"email"`
	Password string `form:"password" validate:"required,max=100"       json:"-"`
}

func (handler UserHandler) Login(ctx echo.Context) error {
	form := new(UserLoginForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		}))
	}

	err := validate.Struct(form)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		}))
	}

	pyld := models.User{
		Email:    form.Email,
		Password: form.Password,
	}

	data, err := handler.UserService.LoginUser(ctx.Request().Context(), &pyld)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err.Error(),
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}

type ForgotPassForm struct {
	Email string `form:"email"    validate:"required,max=100,email"            json:"email"`
}

func (handler UserHandler) ForgotPass(ctx echo.Context) error {
	form := new(ForgotPassForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		}))
	}

	err := validate.Struct(form)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		}))
	}

	data, err := handler.UserService.ForgotPassword(ctx.Request().Context(), form.Email)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err.Error(),
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}

type UserUpdateProfileForm struct {
	Name   string `form:"name"     validate:"required,max=100"                  json:"name"`
	Email  string `form:"email"    validate:"required,max=100,email" json:"email"`
	Gender string `form:"gender"   validate:"required,oneof=male female"        json:"gender"`
}

func (handler UserHandler) UpdateProfile(ctx echo.Context) error {
	ID, err := strconv.Atoi(fmt.Sprint(ctx.Get("ID")))
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err.Error(),
		}))
	}

	form := new(UserUpdateProfileForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		}))
	}

	err = validate.Struct(form)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		}))
	}

	pyld := models.User{
		Name:   form.Name,
		Gender: form.Gender,
		Email:  strings.ToLower(form.Email),
	}

	data, err := handler.UserService.UpdateProfileUser(ctx.Request().Context(), uint(ID), &pyld)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err.Error(),
		}))
	}

	return ctx.JSON(http.StatusOK, entities.ResponseFormater(http.StatusOK, map[string]interface{}{
		"data": data,
	}))
}

type UpdatePasswordForm struct {
	OldPassword string `form:"password_old" validate:"required,max=100"                  json:"-"`
	Password    string `form:"password"     validate:"required,max=100"                  json:"-"`
	Confirm     string `form:"confirm"      validate:"required,max=100,eqfield=Password" json:"-"`
}

func (handler UserHandler) UpdatePassword(ctx echo.Context) error {
	ID, err := strconv.Atoi(fmt.Sprint(ctx.Get("ID")))
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err.Error(),
		}))
	}

	form := new(UpdatePasswordForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		}))
	}

	err = validate.Struct(form)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, entities.ResponseFormater(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		}))
	}

	data, err := handler.UserService.UpdatePasswordUser(ctx.Request().Context(), uint(ID), form.OldPassword, form.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, entities.ResponseFormater(http.StatusBadGateway, map[string]interface{}{
			"error": err,
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
