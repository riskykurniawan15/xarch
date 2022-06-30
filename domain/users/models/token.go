package models

import "time"

const (
	UserTokenTable = "user_tokens"
	MethodVerified = "verified"
	MethodForgot   = "forgot"
)

type UserToken struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"column:user_id"                     json:"name"`
	Method    string    `gorm:"column:method"                      json:"method"`
	Token     string    `gorm:"column:token"                       json:"-"`
	Expired   time.Time `gorm:"column:expired"                     json:"expired"`
	CreatedAt time.Time `gorm:"column:created_at"                  json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at"                  json:"-"`
}

func (UserToken) TableName() string {
	return UserTokenTable
}
