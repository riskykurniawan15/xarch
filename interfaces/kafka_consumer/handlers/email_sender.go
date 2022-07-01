package handlers

import (
	"context"
	"fmt"
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
	handler.log.Info(fmt.Sprint(ctx.Value("topic")))
	key := strings.Split(keys, "_")
	UserID, err := strconv.Atoi(key[1])
	if err != nil {
		return err
	}

	user := models.User{
		ID:    uint(UserID),
		Email: value,
	}

	_, err = handler.UserService.SendTokenForgot(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}

func (handler EmailSenderHandler) SendVerification(ctx context.Context, keys, value string) error {
	handler.log.Info(fmt.Sprint(ctx.Value("topic")))
	key := strings.Split(keys, "_")
	UserID, err := strconv.Atoi(key[1])
	if err != nil {
		return err
	}

	user := models.User{
		ID:    uint(UserID),
		Email: value,
	}

	_, err = handler.UserService.SendEmailVerification(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}
