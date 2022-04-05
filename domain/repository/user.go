package repository

import (
	"github.com/riskykurniawan15/xarch/domain/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) *UserRepo {
	return &UserRepo{
		DB,
	}
}

func (repo UserRepo) SelectUser() ([]*models.User, error) {
	var model []*models.User

	result := repo.DB.
		Model(&models.User{}).
		Find(&model)

	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}
