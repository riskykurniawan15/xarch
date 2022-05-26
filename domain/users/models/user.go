package models

const UserTable = "users"

type User struct {
	ID       uint   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"column:name"                        json:"name"`
	Email    string `gorm:"column:email"                       json:"email"`
	Password string `gorm:"column:password"                    json:"-"`
	Gender   string `gorm:"column:gender"                      json:"gender"`
	Token    string `gorm:"-"                                  json:"token,omitempty"`
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
