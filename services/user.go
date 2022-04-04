package services

import (
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

func (US UserService) GetListUser() {

}
