package models

import (
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

type User struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100)" json:"name"`
	Age     int64  `gorm:"type:bigint(20)" json:"age"`
	Email   string `gorm:"type:varchar(100);unique_index" json:"email"`
	Address string `gorm:"index:addr" json:"address"`
}

func (user User) Validate(v *revel.Validation) {
	v.Check(user.Name, revel.ValidRequired(), revel.ValidMinSize(3), revel.ValidMaxSize(100)).Message("Name must be between 3-100 characters long")
	v.Check(user.Email, revel.ValidRequired(), revel.ValidEmail()).Message("Sorry that is not a valid email address")
	v.Check(user.Age, revel.ValidRequired()).Message("Age cannot be empty")
	v.Check(user.Address, revel.ValidRequired()).Message("Address cannot be empty")
}

func (user User) AddUser() error {
	response := DB.Create(&user)
	return response.Error

}

func (user User) GetUser(id int64) (User, error) {
	response := DB.First(&user, id)
	return user, response.Error
}

func (user User) UpdateUser(id int64) error {
	user.ID = uint(id)
	response := DB.Save(&user)
	return response.Error
}

func (user User) DeleteUser(id int64) error {
	user.ID = uint(id)
	response := DB.Delete(&user)
	return response.Error
}

func (user User) ListUsers() ([]User, error) {
	users := make([]User, 0, 0)
	response := DB.Order("id desc").Find(&users)
	return users, response.Error
}
