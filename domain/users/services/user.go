package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/riskykurniawan15/xarch/domain/users/models"
	"github.com/riskykurniawan15/xarch/domain/users/repository"
	"github.com/riskykurniawan15/xarch/helpers/bcrypt"
	"github.com/riskykurniawan15/xarch/helpers/md5"
)

type UserService struct {
	UserRepo *repository.UserRepo
}

func NewUserService(UR *repository.UserRepo) *UserService {
	return &UserService{
		UR,
	}
}

func (svc UserService) RegisterUser(ctx context.Context, user *models.User) (*models.User, error) {
	Password, err := bcrypt.Hash(user.Password)
	if err != nil {
		return nil, err
	}

	user.Email = strings.ToLower(user.Email)
	user.Password = Password

	_, err = svc.UserRepo.SelectUserDetailByEmail(ctx, user)
	if err == nil {
		return nil, fmt.Errorf("failed register user, user is registered")
	}

	result, err := svc.UserRepo.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}

	err = svc.UserRepo.VerifiedEmailPublish(ctx, result)
	if err != nil {
		fmt.Println(err.Error())
	}

	return result, nil
}

func (svc UserService) LoginUser(ctx context.Context, user *models.User) (*models.User, error) {
	user.Email = strings.ToLower(user.Email)

	result, err := svc.UserRepo.SelectUserDetailByEmail(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("user not registered")
	}

	err = bcrypt.CompareHash(result.Password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("your password dont match")
	}

	token, err := svc.UserRepo.GenerateTokenUser(ctx, result)
	if err != nil {
		return nil, fmt.Errorf("failed generate token")
	}

	result.Token = token

	return result, nil
}

func (svc UserService) GetDetailUser(ctx context.Context, user *models.User) (*models.User, error) {
	user, err := svc.UserRepo.SelectUserDetail(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (svc UserService) SendEmailVerification(ctx context.Context, user *models.User) (*models.User, error) {
	exp := time.Now().Add(time.Minute * 10)

	userData, err := svc.UserRepo.SelectUserDetail(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	user.Name = userData.Name

	enc := md5.Encrypt(exp.Format("2006-01-02 15:04:05"))

	TokenHash, err := bcrypt.Hash(enc)
	if err != nil {
		return nil, err
	}

	token := &models.UserToken{
		UserID:  user.ID,
		Method:  "verified",
		Token:   TokenHash,
		Expired: exp,
	}

	_, err = svc.UserRepo.InsertTokenEmailVerfied(ctx, token)
	if err != nil {
		return nil, err
	}

	token.Token = enc

	err = svc.UserRepo.EmailVerfiedSender(ctx, user, token)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
