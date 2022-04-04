package services

import (
	"github.com/riskykurniawan15/xarch/domain/models"
	"github.com/riskykurniawan15/xarch/domain/repository"
)

type UserService struct {
	UserRepo *repository.UserRepo
}

func NewUserService(UR *repository.UserRepo) *UserService {
	return &UserService{
		UR,
	}
}

func (US UserService) GetListUser() ([]*models.People, error) {
	User, err := US.UserRepo.SelectUser()
	if err != nil {
		return nil, err
	}
	return User, nil
}
