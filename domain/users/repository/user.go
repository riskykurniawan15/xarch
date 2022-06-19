package repository

import (
	"github.com/riskykurniawan15/xarch/config"
	"gorm.io/gorm"
)

type UserRepo struct {
	cfg config.Config
	DB  *gorm.DB
}

func NewUserRepo(cfg config.Config, DB *gorm.DB) *UserRepo {
	return &UserRepo{
		cfg,
		DB,
	}
}
