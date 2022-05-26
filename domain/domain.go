package domain

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"github.com/riskykurniawan15/xarch/config"

	AlquranRepo "github.com/riskykurniawan15/xarch/domain/alquran/repository"
	AlquranService "github.com/riskykurniawan15/xarch/domain/alquran/services"
	UserRepo "github.com/riskykurniawan15/xarch/domain/users/repository"
	UserService "github.com/riskykurniawan15/xarch/domain/users/services"
)

type Repo struct {
	UserRepo    *UserRepo.UserRepo
	AlquranRepo *AlquranRepo.AlquranRepo
}

type Service struct {
	UserService    *UserService.UserService
	AlquranService *AlquranService.AlquranService
}

func StartRepo(cfg config.Config, DB *gorm.DB, RDB *redis.Client) *Repo {
	UserRepo := UserRepo.NewUserRepo(cfg, DB)
	AlquranRepo := AlquranRepo.NewAlquranRepo(cfg, RDB)

	return &Repo{
		UserRepo,
		AlquranRepo,
	}
}

func StartService(cfg config.Config, DB *gorm.DB, RDB *redis.Client) *Service {
	Repo := StartRepo(cfg, DB, RDB)
	UserService := UserService.NewUserService(Repo.UserRepo)
	AlquranService := AlquranService.NewAlquranService(Repo.AlquranRepo)

	return &Service{
		UserService,
		AlquranService,
	}
}
