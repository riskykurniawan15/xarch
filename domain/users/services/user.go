package services

import (
	"github.com/riskykurniawan15/xarch/domain/users/models"
	"github.com/riskykurniawan15/xarch/domain/users/repository"
)

type UserService struct {
	UserRepo *repository.UserRepo
}

func NewUserService(UR *repository.UserRepo) *UserService {
	return &UserService{
		UR,
	}
}

func (US UserService) GetListUser() ([]*models.User, error) {
	User, err := US.UserRepo.SelectUser()
	if err != nil {
		return nil, err
	}
	return User, nil
}
