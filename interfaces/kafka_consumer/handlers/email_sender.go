package handlers

import (
	"context"
	"strconv"
	"strings"

	"github.com/riskykurniawan15/xarch/domain/users/models"
	"github.com/riskykurniawan15/xarch/domain/users/services"
	"github.com/riskykurniawan15/xarch/logger"
)

type EmailSenderHandler struct {
	log         logger.Logger
	UserService *services.UserService
}

func NewEmailSenderHandler(log logger.Logger, US *services.UserService) *EmailSenderHandler {
	return &EmailSenderHandler{
		log,
		US,
	}
}

func (handler EmailSenderHandler) SendForgotToken(ctx context.Context, keys, value string) error {
	handler.log.InfoW("Incomming Process Send Forgot Password By Emal", map[string]interface{}{
		"topic": ctx.Value("topic"),
		"keys":  keys,
		"value": value,
	})
	key := strings.Split(keys, "_")
	UserID, err := strconv.Atoi(key[1])
	if err != nil {
		return err
	}

	user := models.User{
		ID:    uint(UserID),
		Email: value,
	}

	_, err_handler := handler.UserService.SendTokenForgot(ctx, &user)
	if err_handler != nil {
		handler.log.ErrorW("Failed Send Forgot Password By Emal", map[string]interface{}{
			"topic": ctx.Value("topic"),
			"data":  user,
			"error": err_handler.Errors.Error(),
		})
		return err_handler.Errors
	}

	handler.log.InfoW("Success Send Forgot Password By Emal", map[string]interface{}{
		"topic": ctx.Value("topic"),
		"data":  user,
	})
	return nil
}

func (handler EmailSenderHandler) SendVerification(ctx context.Context, keys, value string) error {
	handler.log.InfoW("Incomming Process Send Email Verification", map[string]interface{}{
		"topic": ctx.Value("topic"),
		"keys":  keys,
		"value": value,
	})
	key := strings.Split(keys, "_")
	UserID, err := strconv.Atoi(key[1])
	if err != nil {
		handler.log.ErrorW("Failed Send Email Verification", map[string]interface{}{
			"topic": ctx.Value("topic"),
			"error": err.Error(),
		})
		return err
	}

	user := models.User{
		ID:    uint(UserID),
		Email: value,
	}

	_, err_handler := handler.UserService.SendEmailVerification(ctx, &user)
	if err_handler != nil {
		handler.log.ErrorW("Failed Send Email Verification", map[string]interface{}{
			"topic": ctx.Value("topic"),
			"data":  user,
			"error": err_handler.Errors.Error(),
		})
		return err_handler.Errors
	}

	handler.log.InfoW("Success Send Email Verification", map[string]interface{}{
		"topic": ctx.Value("topic"),
		"data":  user,
	})
	return nil
}
