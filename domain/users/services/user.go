package services

import (
	"context"
	"fmt"
	"math/rand"
	"mime/multipart"
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

func (svc UserService) GetDetailUser(ctx context.Context, user *models.User) (*models.User, *errors.ErrorResponse) {
	user, err := svc.UserRepo.SelectUserDetail(ctx, user)
	if err != nil {
		return nil, errors.InternalError.NewError(fmt.Errorf("user not found"))
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

func (svc UserService) SendTokenForgot(ctx context.Context, user *models.User) (*models.User, *errors.ErrorResponse) {
	exp := time.Now().Add(time.Minute * 10)

	userData, err := svc.UserRepo.SelectUserDetail(ctx, user)
	if err != nil {
		return nil, errors.BadRequest.NewError(fmt.Errorf("user not found"))
	}

	user.Name = userData.Name

	rand.Seed(time.Now().UnixNano())
	enc := fmt.Sprintf("%05d", rand.Intn(99999))

	TokenHash, err := bcrypt.Hash(enc)
	if err != nil {
		return nil, errors.InternalError.NewError(err)
	}

	token := &models.UserToken{
		UserID:  user.ID,
		Method:  models.MethodForgot,
		Token:   TokenHash,
		Expired: exp,
	}

	_, err = svc.UserRepo.InsertTokenUser(ctx, token)
	if err != nil {
		return nil, errors.InternalError.NewError(err)
	}

	token.Token = enc

	err = svc.UserRepo.EmailForgotSender(ctx, user, token)
	if err != nil {
		return nil, errors.InternalError.NewError(err)
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

func (svc UserService) SendEmailVerification(ctx context.Context, user *models.User) (*models.User, *errors.ErrorResponse) {
	exp := time.Now().Add(time.Minute * 10)

	userData, err := svc.UserRepo.SelectUserDetail(ctx, user)
	if err != nil {
		return nil, errors.BadRequest.NewError(fmt.Errorf("user not found"))
	}

	user.Name = userData.Name

	enc := md5.Encrypt(exp.Format("2006-01-02 15:04:05"))

	TokenHash, err := bcrypt.Hash(enc)
	if err != nil {
		return nil, errors.InternalError.NewError(err)
	}

	token := &models.UserToken{
		UserID:  user.ID,
		Method:  models.MethodVerified,
		Token:   TokenHash,
		Expired: exp,
	}

	_, err = svc.UserRepo.InsertTokenUser(ctx, token)
	if err != nil {
		return nil, errors.InternalError.NewError(err)
	}

	token.Token = enc

	err = svc.UserRepo.EmailVerfiedSender(ctx, user, token)
	if err != nil {
		return nil, errors.InternalError.NewError(err)
	}
	return nil, nil
}

func (svc UserService) EmailVerification(ctx context.Context, ID uint, token string) (string, *errors.ErrorResponse) {
	userData, err := svc.UserRepo.SelectUserDetail(ctx, &models.User{ID: ID})
	if err != nil {
		return "", errors.BadRequest.NewError(fmt.Errorf("user not found"))
	}

	if fmt.Sprint(userData.VerifiedAt) != fmt.Sprint(time.Time{}) {
		return "", errors.BadRequest.NewError(fmt.Errorf("akun telah terverifikasi"))
	}

	tokenQuery := &models.UserToken{
		UserID: ID,
		Method: models.MethodVerified,
	}
	tokenList, err := svc.UserRepo.GetTokenUser(ctx, tokenQuery)
	if err != nil {
		return "", errors.InternalError.NewError(err)
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
		return "", errors.BadRequest.NewError(fmt.Errorf("verifikasi gagal token tidak ditemukan atau sudah expired"))
	}

	userData.VerifiedAt = today
	_, err = svc.UserRepo.UpdateUser(ctx, ID, userData)
	if err != nil {
		return "", errors.InternalError.NewError(fmt.Errorf("failed update user"))
	}

	svc.UserRepo.DeleteTokenUser(ctx, tokenQuery)

	return "varifikasi success", nil
}

func (svc UserService) UploadProfileImage(ctx context.Context, ID uint, file *multipart.FileHeader) (*models.User, *errors.ErrorResponse) {
	userData, err := svc.UserRepo.SelectUserDetail(ctx, &models.User{ID: ID})
	if err != nil {
		return nil, errors.BadRequest.NewError(fmt.Errorf("user not found"))
	}

	src, err := file.Open()
	if err != nil {
		return nil, errors.InternalError.NewError(fmt.Errorf("error on uploading file"))
	}
	defer src.Close()

	image, err := svc.UserRepo.StoreFile(ctx, src)
	if err != nil {
		return nil, errors.InternalError.NewError(fmt.Errorf("error on storing file to cloud"))
	}

	_, err = svc.UserRepo.UpdateUser(ctx, ID, &models.User{
		Image: image,
	})
	if err != nil {
		return nil, errors.InternalError.NewError(fmt.Errorf("failed to update password"))
	}

	if userData.Image != "" {
		svc.UserRepo.DestroyFile(ctx, userData.Image)
	}

	result, err := svc.UserRepo.SelectUserDetail(ctx, &models.User{ID: ID})
	if err != nil {
		return nil, errors.BadRequest.NewError(fmt.Errorf("user not found"))
	}

	return result, nil
}

func (svc UserService) GetProfileImage(ctx context.Context, ID uint) ([]byte, *errors.ErrorResponse) {
	userData, err := svc.UserRepo.SelectUserDetail(ctx, &models.User{ID: ID})
	if err != nil {
		return nil, errors.BadRequest.NewError(fmt.Errorf("user not found"))
	}

	if userData.Image == "" {
		return nil, errors.NotFound.NewError(fmt.Errorf("no image"))
	}

	meta, err := svc.UserRepo.RetrieveFile(ctx, userData.Image)
	if err != nil {
		return nil, errors.InternalError.NewError(fmt.Errorf("error on retrieve meta file to cloud"))
	}

	result, err := svc.UserRepo.GetFile(ctx, meta.URL)
	if err != nil {
		return nil, errors.InternalError.NewError(fmt.Errorf("error on storing file to cloud"))
	}

	return result, nil
}
