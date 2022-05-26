package services

import (
	"context"
	"fmt"
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

	_, err = svc.UserRepo.SelectUserDetailByEmail(ctx, user)
	if err == nil {
		return nil, fmt.Errorf("failed register user, user is registered")
	}

	result, err := svc.UserRepo.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (svc UserService) LoginUser(ctx context.Context, user *models.User) (*models.User, error) {
	user.Email = strings.ToLower(user.Email)

	result, err := svc.UserRepo.SelectUserDetailByEmail(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("user not registered")
	}

	err = bcrypt.CompareHash(result.Password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("your password dont match")
	}

	token, err := svc.UserRepo.GenerateTokenUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed generate token")
	}

	result.Token = token

	return result, nil
}

func (US UserService) GetListUser() ([]*models.User, error) {
	User, err := US.UserRepo.SelectUser()
	if err != nil {
		return nil, err
	}
	return User, nil
}
