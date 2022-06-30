package repository

import (
	"context"

	"github.com/riskykurniawan15/xarch/domain/users/models"
)

func (repo UserRepo) InsertTokenEmailVerfied(ctx context.Context, token *models.UserToken) (*models.UserToken, error) {
	if err := repo.DB.
		WithContext(ctx).
		Create(token).
		Error; err != nil {
		return nil, err
	}

	return token, nil
}

func (repo UserRepo) GetTokenEmailVerfied(ctx context.Context, userToken *models.UserToken) (*[]models.UserToken, error) {
	var model *[]models.UserToken

	result := repo.DB.
		WithContext(ctx).
		Model(&models.UserToken{}).
		Where(userToken).
		Find(&model)

	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (repo UserRepo) DeleteTokenEmailVerfied(ctx context.Context, userToken *models.UserToken) error {
	if err := repo.DB.
		WithContext(ctx).
		Where(userToken).
		Delete(&models.UserToken{}).
		Error; err != nil {
		return err
	}

	return nil
}
