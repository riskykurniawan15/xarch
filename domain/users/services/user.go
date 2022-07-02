package services

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/riskykurniawan15/xarch/domain/users/models"
	"github.com/riskykurniawan15/xarch/domain/users/repository"
	"github.com/riskykurniawan15/xarch/helpers/bcrypt"
	"github.com/riskykurniawan15/xarch/helpers/errors"
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

var (
	validate *validator.Validate = validator.New()
)

func (svc UserService) RegisterUser(ctx context.Context, req *models.UserRegisterForm) (*models.User, *errors.ErrorResponse) {
	err := validate.Struct(req)
	if err != nil {
		return nil, errors.BadRequest.NewError(err)
	}

	user := models.SetUserData(req.Name, req.Email, req.Password, req.Gender)

	_, err = svc.UserRepo.SelectUserDetailByEmail(ctx, user)
	if err == nil {
		return nil, errors.BadRequest.NewError(fmt.Errorf("failed register user, user is registered"))
	}

	Password, err := bcrypt.Hash(user.Password)
	if err != nil {
		return nil, errors.InternalError.NewError(err)
	}

	user.Email = strings.ToLower(user.Email)
	user.Password = Password

	result, err := svc.UserRepo.InsertUser(ctx, user)
	if err != nil {
		return nil, errors.InternalError.NewError(err)
	}

	err = svc.UserRepo.VerifiedEmailPublish(ctx, result)
	if err != nil {
		return nil, errors.InternalError.NewError(fmt.Errorf("failed to send email"))
	}

	return result, nil
}

func (svc UserService) LoginUser(ctx context.Context, req *models.UserLoginForm) (*models.User, *errors.ErrorResponse) {
	err := validate.Struct(req)
	if err != nil {
		return nil, errors.BadRequest.NewError(err)
	}

	user := &models.User{
		Email:    strings.ToLower(req.Email),
		Password: req.Password,
	}

	result, err := svc.UserRepo.SelectUserDetailByEmail(ctx, user)
	if err != nil {
		return nil, errors.BadRequest.NewError(fmt.Errorf("user not registered"))
	}

	err = bcrypt.CompareHash(result.Password, user.Password)
	if err != nil {
		return nil, errors.BadRequest.NewError(fmt.Errorf("your password dont match"))
	}

	token, err := svc.UserRepo.GenerateTokenUser(ctx, result)
	if err != nil {
		return nil, errors.InternalError.NewError(fmt.Errorf("failed generate token"))
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

func (svc UserService) UpdateProfileUser(ctx context.Context, ID uint, req *models.UserUpdateProfileForm) (*models.User, *errors.ErrorResponse) {
	err := validate.Struct(req)
	if err != nil {
		return nil, errors.BadRequest.NewError(err)
	}

	userData, err := svc.UserRepo.SelectUserDetail(ctx, &models.User{ID: ID})
	if err != nil {
		return nil, errors.InternalError.NewError(fmt.Errorf("user not found"))
	}

	pyld := &models.User{
		Name:   req.Name,
		Gender: req.Gender,
		Email:  strings.ToLower(req.Email),
	}

	reset_verif := false

	if userData.Email != pyld.Email {
		_, err = svc.UserRepo.SelectUserDetailByEmail(ctx, pyld)
		if err == nil {
			return nil, errors.BadRequest.NewError(fmt.Errorf("failed update your profile, email is registered"))
		}

		pyld.VerifiedAt = time.Time{}
		reset_verif = true
	}

	_, err = svc.UserRepo.UpdateUser(ctx, ID, pyld)
	if err != nil {
		return nil, errors.InternalError.NewError(fmt.Errorf("failed to update data"))
	}

	userData, err = svc.UserRepo.SelectUserDetail(ctx, &models.User{ID: ID})
	if err != nil {
		return nil, errors.InternalError.NewError(fmt.Errorf("user not found"))
	}

	if reset_verif {
		svc.UserRepo.VerifiedEmailPublish(ctx, userData)
	}

	return userData, nil
}

func (svc UserService) UpdatePasswordUser(ctx context.Context, ID uint, req *models.UpdatePasswordForm) (string, *errors.ErrorResponse) {
	err := validate.Struct(req)
	if err != nil {
		return "", errors.BadRequest.NewError(err)
	}

	result, err := svc.UserRepo.SelectUserDetail(ctx, &models.User{ID: ID})
	if err != nil {
		return "", errors.InternalError.NewError(fmt.Errorf("user not found"))
	}

	err = bcrypt.CompareHash(result.Password, req.OldPassword)
	if err != nil {
		return "", errors.InternalError.NewError(fmt.Errorf("your password dont match"))
	}

	Password, _ := bcrypt.Hash(req.Password)

	_, err = svc.UserRepo.UpdateUser(ctx, ID, &models.User{
		Password: Password,
	})
	if err != nil {
		return "", errors.InternalError.NewError(fmt.Errorf("failed to update password"))
	}

	return "success update password", nil
}

func (svc UserService) ForgotPassword(ctx context.Context, req *models.ForgotPassForm) (string, *errors.ErrorResponse) {
	err := validate.Struct(req)
	if err != nil {
		return "", errors.BadRequest.NewError(err)
	}

	user, err := svc.UserRepo.SelectUserDetailByEmail(ctx, &models.User{Email: strings.ToLower(req.Email)})
	if err != nil {
		return "", errors.BadRequest.NewError(fmt.Errorf("user not found"))
	}

	err = svc.UserRepo.ForgotPassByEmailPublish(ctx, user)
	if err != nil {
		return "", errors.InternalError.NewError(fmt.Errorf("failed send email"))
	}

	return "success send token reset password by email", nil
}

func (svc UserService) SendTokenForgot(ctx context.Context, user *models.User) (*models.User, error) {
	exp := time.Now().Add(time.Minute * 10)

	userData, err := svc.UserRepo.SelectUserDetail(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	user.Name = userData.Name

	rand.Seed(time.Now().UnixNano())
	enc := fmt.Sprintf("%05d", rand.Intn(99999))

	TokenHash, err := bcrypt.Hash(enc)
	if err != nil {
		return nil, err
	}

	token := &models.UserToken{
		UserID:  user.ID,
		Method:  models.MethodForgot,
		Token:   TokenHash,
		Expired: exp,
	}

	_, err = svc.UserRepo.InsertTokenUser(ctx, token)
	if err != nil {
		return nil, err
	}

	token.Token = enc

	err = svc.UserRepo.EmailForgotSender(ctx, user, token)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (svc UserService) ResetPass(ctx context.Context, req *models.ResetPassForm) (string, *errors.ErrorResponse) {
	err := validate.Struct(req)
	if err != nil {
		return "", errors.BadRequest.NewError(err)
	}

	userData, err := svc.UserRepo.SelectUserDetailByEmail(ctx, &models.User{Email: strings.ToLower(req.Email)})
	if err != nil {
		return "", errors.BadRequest.NewError(fmt.Errorf("user not found"))
	}

	tokenQuery := &models.UserToken{
		UserID: userData.ID,
		Method: models.MethodForgot,
	}
	tokenList, err := svc.UserRepo.GetTokenUser(ctx, tokenQuery)
	if err != nil {
		return "", errors.InternalError.NewError(err)
	}

	exist := false
	today := time.Now()
	for _, val := range *tokenList {
		if bcrypt.CompareHash(val.Token, req.Token) == nil {
			if today.Before(val.Expired) {
				exist = true
				break
			}
		}
	}

	if !exist {
		return "", errors.BadRequest.NewError(fmt.Errorf("token tidak ditemukan atau sudah expired"))
	}

	userData.Password, _ = bcrypt.Hash(req.Password)
	_, err = svc.UserRepo.UpdateUser(ctx, userData.ID, userData)
	if err != nil {
		return "", errors.InternalError.NewError(fmt.Errorf("failed update password"))
	}

	svc.UserRepo.DeleteTokenUser(ctx, tokenQuery)

	return "success update password", nil
}

func (svc UserService) ReSendEmailVerification(ctx context.Context, req *models.ReSendVerificationForm) (string, *errors.ErrorResponse) {
	err := validate.Struct(req)
	if err != nil {
		return "", errors.BadRequest.NewError(err)
	}

	user, err := svc.UserRepo.SelectUserDetailByEmail(ctx, &models.User{Email: req.Email})
	if err != nil {
		return "", errors.BadRequest.NewError(fmt.Errorf("user not found"))
	}

	if fmt.Sprint(user.VerifiedAt) != fmt.Sprint(time.Time{}) {
		return "", errors.BadRequest.NewError(fmt.Errorf("failed send verification email, your account has been verified"))
	}

	err = svc.UserRepo.VerifiedEmailPublish(ctx, user)
	if err != nil {
		return "", errors.InternalError.NewError(fmt.Errorf("failed send verification email"))
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

	_, err = svc.UserRepo.InsertTokenUser(ctx, token)
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
	tokenList, err := svc.UserRepo.GetTokenUser(ctx, tokenQuery)
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

	svc.UserRepo.DeleteTokenUser(ctx, tokenQuery)

	return "varifikasi success", nil
}
