package models

type UserRegisterForm struct {
	Name     string `form:"name"     validate:"required,max=100"                  json:"name"`
	Email    string `form:"email"    validate:"required,max=100,email"            json:"email"`
	Password string `form:"password" validate:"required,max=100"                  json:"-"`
	Confirm  string `form:"confirm"  validate:"required,max=100,eqfield=Password" json:"-"`
	Gender   string `form:"gender"   validate:"required,oneof=male female"        json:"gender"`
}

type ReSendVerificationForm struct {
	Email string `form:"email"    validate:"required,max=100,email"            json:"email"`
}

type UserLoginForm struct {
	Email    string `form:"email"    validate:"required,max=100,email" json:"email"`
	Password string `form:"password" validate:"required,max=100"       json:"-"`
}

type ForgotPassForm struct {
	Email string `form:"email"    validate:"required,max=100,email"            json:"email"`
}

type ResetPassForm struct {
	Token    string `form:"token"        validate:"required,min=5,max=5"              json:"-"`
	Email    string `form:"email"        validate:"required,max=100,email"            json:"email"`
	Password string `form:"password"     validate:"required,max=100"                  json:"-"`
	Confirm  string `form:"confirm"      validate:"required,max=100,eqfield=Password" json:"-"`
}

type UserUpdateProfileForm struct {
	Name   string `form:"name"     validate:"required,max=100"                  json:"name"`
	Email  string `form:"email"    validate:"required,max=100,email" json:"email"`
	Gender string `form:"gender"   validate:"required,oneof=male female"        json:"gender"`
}

type UpdatePasswordForm struct {
	OldPassword string `form:"password_old" validate:"required,max=100"                  json:"-"`
	Password    string `form:"password"     validate:"required,max=100"                  json:"-"`
	Confirm     string `form:"confirm"      validate:"required,max=100,eqfield=Password" json:"-"`
}
