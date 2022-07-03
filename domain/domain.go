package domain

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"github.com/riskykurniawan15/xarch/config"

	AlquranRepo "github.com/riskykurniawan15/xarch/domain/alquran/repository"
	AlquranService "github.com/riskykurniawan15/xarch/domain/alquran/services"
	HealthRepo "github.com/riskykurniawan15/xarch/domain/health/repository"
	HealthService "github.com/riskykurniawan15/xarch/domain/health/services"
	UserRepo "github.com/riskykurniawan15/xarch/domain/users/repository"
	UserService "github.com/riskykurniawan15/xarch/domain/users/services"
)

type Repo struct {
	HealthRepo  *HealthRepo.HealthRepo
	UserRepo    *UserRepo.UserRepo
	AlquranRepo *AlquranRepo.AlquranRepo
}

type Service struct {
	HealthService  *HealthService.HealthService
	UserService    *UserService.UserService
	AlquranService *AlquranService.AlquranService
}

func StartRepo(cfg config.Config, DB *gorm.DB, RDB *redis.Client) *Repo {
	HealthRepo := HealthRepo.NewHealthRepo(cfg, DB, RDB)
	UserRepo := UserRepo.NewUserRepo(cfg, DB)
	AlquranRepo := AlquranRepo.NewAlquranRepo(cfg, RDB)

	return &Repo{
		HealthRepo,
		UserRepo,
		AlquranRepo,
	}
}

func StartService(cfg config.Config, DB *gorm.DB, RDB *redis.Client) *Service {
	Repo := StartRepo(cfg, DB, RDB)
	HealthService := HealthService.NewHealthService(Repo.HealthRepo)
	UserService := UserService.NewUserService(Repo.UserRepo)
	AlquranService := AlquranService.NewAlquranService(Repo.AlquranRepo)

	return &Service{
		HealthService,
		UserService,
		AlquranService,
	}
}
