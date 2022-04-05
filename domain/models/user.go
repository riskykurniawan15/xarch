package models

const UserTable = "users"

type User struct {
	ID      uint   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name    string `gorm:"column:name"                        json:"name"`
	Email   string `gorm:"column:email"                       json:"email"`
	Contact string `gorm:"column:contact"                     json:"contact"`
	Gender  string `gorm:"column:gender"                      json:"gender"`
}

func (User) TableName() string {
	return UserTable
}
