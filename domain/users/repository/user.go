package repository

import (
	"context"

	"github.com/riskykurniawan15/xarch/domain/users/models"
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

func (repo UserRepo) InsertUser(ctx context.Context, user *models.User) (*models.User, error) {
	if err := repo.DB.
		WithContext(ctx).
		Create(user).
		Error; err != nil {
		return nil, err
	}

	return user, nil
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
