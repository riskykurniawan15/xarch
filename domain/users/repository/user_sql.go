package repository

import (
	"context"

	"github.com/riskykurniawan15/xarch/domain/users/models"
	"github.com/riskykurniawan15/xarch/helpers/jwt"
)

func (repo UserRepo) GenerateTokenUser(ctx context.Context, user *models.User) (string, error) {
	NewJwt := jwt.NewJwtToken(repo.cfg)

	pyld := &jwt.JwtCustomClaims{
		ID:    user.ID,
		Email: user.Email,
	}

	token, err := NewJwt.GenerateToken(pyld)
	if err != nil {
		return "", err
	}

	return token, nil
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

func (repo UserRepo) UpdateUser(ctx context.Context, ID uint, user *models.User) (*models.User, error) {
	if err := repo.DB.
		WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", ID).
		Updates(user).
		Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo UserRepo) SelectUserDetailByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	var model *models.User

	result := repo.DB.
		WithContext(ctx).
		Model(&models.User{}).
		Where("email = ?", user.Email).
		First(&model)

	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (repo UserRepo) SelectUserDetail(ctx context.Context, user *models.User) (*models.User, error) {
	var model *models.User

	result := repo.DB.
		WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", user.ID).
		First(&model)

	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}
