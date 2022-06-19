package handlers

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/riskykurniawan15/xarch/domain/users/models"
	"github.com/riskykurniawan15/xarch/domain/users/services"
)

type EmailSenderHandler struct {
	UserService *services.UserService
}

func NewEmailSenderHandler(US *services.UserService) *EmailSenderHandler {
	return &EmailSenderHandler{
		US,
	}
}

func (handler EmailSenderHandler) SendVerification(ctx context.Context, keys, value string) error {
	key := strings.Split(keys, "_")
	UserID, err := strconv.Atoi(key[1])
	if err != nil {
		return err
	}

	user := models.User{
		ID:    uint(UserID),
		Email: value,
	}

	result, err := handler.UserService.SendEmailVerification(ctx, &user)
	if err != nil {
		return err
	}

	fmt.Print(result)

	return nil
}
