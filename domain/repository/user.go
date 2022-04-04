package repository

import (
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

func (DB UserRepo) SelectUser() {

}
