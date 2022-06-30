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

func (svc UserService) UpdateProfileUser(ctx context.Context, ID uint, user *models.User) (*models.User, error) {
	userData, err := svc.UserRepo.SelectUserDetail(ctx, &models.User{ID: ID})
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	pyld := &models.User{
		Name:   user.Name,
		Gender: user.Gender,
		Email:  strings.ToLower(user.Email),
	}

	reset_verif := false

	if userData.Email != pyld.Email {
		_, err = svc.UserRepo.SelectUserDetailByEmail(ctx, pyld)
		if err == nil {
			return nil, fmt.Errorf("failed update your profile, email is registered")
		}

		pyld.VerifiedAt = time.Time{}
		reset_verif = true
	}

	_, err = svc.UserRepo.UpdateUser(ctx, ID, pyld)
	if err != nil {
		return nil, fmt.Errorf("failed to update data")
	}

	userData, err = svc.UserRepo.SelectUserDetail(ctx, &models.User{ID: ID})
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	if reset_verif {
		svc.UserRepo.VerifiedEmailPublish(ctx, userData)
	}

	return userData, nil
}

func (svc UserService) UpdatePasswordUser(ctx context.Context, ID uint, old_password, new_password string) (string, error) {
	result, err := svc.UserRepo.SelectUserDetail(ctx, &models.User{ID: ID})
	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHash(result.Password, old_password)
	if err != nil {
		return "", fmt.Errorf("your password dont match")
	}

	Password, _ := bcrypt.Hash(new_password)

	_, err = svc.UserRepo.UpdateUser(ctx, ID, &models.User{
		Password: Password,
	})
	if err != nil {
		return "", fmt.Errorf("failed to update password")
	}

	return "success update password", nil
}

func (svc UserService) ReSendEmailVerification(ctx context.Context, email string) (string, error) {
	user, err := svc.UserRepo.SelectUserDetailByEmail(ctx, &models.User{Email: email})
	if err != nil {
		return "user not found", nil
	}

	if fmt.Sprint(user.VerifiedAt) != fmt.Sprint(time.Time{}) {
		return "failed send verification email, your account has been verified", err
	}

	err = svc.UserRepo.VerifiedEmailPublish(ctx, user)
	if err != nil {
		return "failed send verification email", err
	}

	return "success send verification email", nil
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
		Method:  models.MethodVerified,
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

func (svc UserService) EmailVerification(ctx context.Context, ID uint, token string) (string, error) {
	userData, err := svc.UserRepo.SelectUserDetail(ctx, &models.User{ID: ID})
	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	if fmt.Sprint(userData.VerifiedAt) != fmt.Sprint(time.Time{}) {
		return "akun telah terverifikasi", err
	}

	tokenQuery := &models.UserToken{
		UserID: ID,
		Method: models.MethodVerified,
	}
	tokenList, err := svc.UserRepo.GetTokenEmailVerfied(ctx, tokenQuery)
	if err != nil {
		return "", err
	}

	exist := false
	today := time.Now()
	for _, val := range *tokenList {
		if bcrypt.CompareHash(val.Token, token) == nil {
			if today.Before(val.Expired) {
				exist = true
				break
			}
		}
	}

	if !exist {
		return "verifikasi gagal token tidak ditemukan atau sudah expired", nil
	}

	userData.VerifiedAt = today
	_, err = svc.UserRepo.UpdateUser(ctx, ID, userData)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	svc.UserRepo.DeleteTokenEmailVerfied(ctx, tokenQuery)

	return "varifikasi success", nil
}
