package models

const PeopleTable = "peoples"

type People struct {
	ID      uint   `gorm:"column:id;primaryKey"`
	Name    string `gorm:"column:name"`
	Email   string `gorm:"column:email"`
	Contact string `gorm:"column:contact"`
	Gender  string `gorm:"column:gender"`
}

func (People) TableName() string {
	return PeopleTable
}
