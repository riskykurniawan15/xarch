package models

import "time"

const UserTable = "users"

type User struct {
	ID         uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"column:name"                        json:"name"`
	Email      string    `gorm:"column:email"                       json:"email"`
	Password   string    `gorm:"column:password"                    json:"-"`
	Gender     string    `gorm:"column:gender"                      json:"gender"`
	Image      string    `gorm:"column:image"                       json:"image"`
	VerifiedAt time.Time `gorm:"column:verified_at;default:null"    json:"-"`
	CreatedAt  time.Time `gorm:"column:created_at"                  json:"-"`
	UpdatedAt  time.Time `gorm:"column:updated_at"                  json:"-"`
	Token      string    `gorm:"-"                                  json:"token,omitempty"`
}

func (User) TableName() string {
	return UserTable
}

func SetUserData(Name, Email, Password, Gender string) *User {
	var user User
	user.Name = Name
	user.Email = Email
	user.Password = Password
	user.Gender = Gender
	return &user
}
