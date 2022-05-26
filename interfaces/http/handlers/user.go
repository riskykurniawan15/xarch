package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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

func (handler UserHandler) Index(ctx echo.Context) error {
	User, err := handler.UserService.GetListUser()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"status": "504",
			"error":  fmt.Sprint(err),
		})
	}

	return ctx.JSON(http.StatusOK, User)
}

func (handler UserHandler) Show(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotFound, entities.ResponseFormater(http.StatusNotFound, map[string]interface{}{
		"error": "error",
	}))
}

func (handler UserHandler) Store(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotFound, entities.ResponseFormater(http.StatusNotFound, map[string]interface{}{
		"error": "error",
	}))
}

func (handler UserHandler) Update(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotFound, entities.ResponseFormater(http.StatusNotFound, map[string]interface{}{
		"error": "error",
	}))
}

func (handler UserHandler) Delete(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotFound, entities.ResponseFormater(http.StatusNotFound, map[string]interface{}{
		"error": "error",
	}))
}
