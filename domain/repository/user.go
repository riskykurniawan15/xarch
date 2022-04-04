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

func (repo UserRepo) SelectUser() ([]*models.People, error) {
	var model []*models.People

	result := repo.DB.
		Model(&models.People{}).
		Find(&model)

	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}
