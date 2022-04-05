package models

const UserTable = "users"

type User struct {
	ID      uint   `gorm:"column:id;primaryKey"`
	Name    string `gorm:"column:name"`
	Email   string `gorm:"column:email"`
	Contact string `gorm:"column:contact"`
	Gender  string `gorm:"column:gender"`
}

func (User) TableName() string {
	return UserTable
}
