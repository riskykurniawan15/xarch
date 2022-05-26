package services

import (
	"context"
	"strings"

	"github.com/riskykurniawan15/xarch/domain/users/models"
	"github.com/riskykurniawan15/xarch/domain/users/repository"
	"github.com/riskykurniawan15/xarch/helpers/bcrypt"
)

type UserService struct {
	UserRepo *repository.UserRepo
}

func NewUserService(UR *repository.UserRepo) *UserService {
	return &UserService{
		UR,
	}
}

func (svc UserService) RegisterUser(ctx context.Context, user *models.User) (*models.User, error) {
	Password, err := bcrypt.Hash(user.Password)
	if err != nil {
		return nil, err
	}

	user.Email = strings.ToLower(user.Email)
	user.Password = Password

	result, err := svc.UserRepo.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (US UserService) GetListUser() ([]*models.User, error) {
	User, err := US.UserRepo.SelectUser()
	if err != nil {
		return nil, err
	}
	return User, nil
}
