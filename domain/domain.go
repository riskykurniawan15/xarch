package domain

import (
	UserRepo "github.com/riskykurniawan15/xarch/domain/users/repository"
	UserService "github.com/riskykurniawan15/xarch/domain/users/services"
	"gorm.io/gorm"
)

type Repo struct {
	UserRepo *UserRepo.UserRepo
}

type Service struct {
	UserService *UserService.UserService
}

func StartRepo(DB *gorm.DB) *Repo {
	UserRepo := UserRepo.NewUserRepo(DB)

	return &Repo{
		UserRepo,
	}
}

func StartService(DB *gorm.DB) *Service {
	Repo := StartRepo(DB)
	UserService := UserService.NewUserService(Repo.UserRepo)

	return &Service{
		UserService,
	}
}
