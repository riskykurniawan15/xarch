package domain

import (
	"github.com/riskykurniawan15/xarch/domain/users/repository"
	"github.com/riskykurniawan15/xarch/domain/users/services"
	"gorm.io/gorm"
)

type Repo struct {
	UserRepo *repository.UserRepo
}

type Service struct {
	UserService *services.UserService
}

func StartRepo(DB *gorm.DB) *Repo {
	UserRepo := repository.NewUserRepo(DB)

	return &Repo{
		UserRepo,
	}
}

func StartService(DB *gorm.DB) *Service {
	Repo := StartRepo(DB)
	UserService := services.NewUserService(Repo.UserRepo)

	return &Service{
		UserService,
	}
}
