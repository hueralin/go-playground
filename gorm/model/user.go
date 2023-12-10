package model

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Age      int    `gorm:"column:age"`
}

type CreateUserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type GetUserListDto struct {
	PageNum    int    `json:"page_num"`
	PageSize   int    `json:"page_size"`
	OrderField string `json:"order_field"`
	Order      string `json:"order"`
}

type UpdateUserAgeDto struct {
	Id  int `json:"id"`
	Age int `json:"age"`
}

func CreateUser(user *User) error {
	return Conn.Create(user).Error
}

func GetUserById(id int) (User, error) {
	var user User
	err := Conn.Where("ID = ?", id).First(&user).Error
	return user, err
}

func GetUserList(param GetUserListDto) ([]User, error) {
	var users []User
	if param.OrderField == "" {
		param.OrderField = "created_at"
	}
	if param.Order == "" {
		param.Order = "asc"
	}
	order := fmt.Sprintf("%s %s", param.OrderField, param.Order)
	err := Conn.Order(order).Limit(param.PageSize).Offset((param.PageNum - 1) * param.PageSize).Find(&users).Error
	return users, err
}

func UpdateUserAgeById(id, age int) error {
	return Conn.Model(&User{}).Where("id = ?", id).Update("age", age).Error
}

func DeleteUserById(id int) error {
	return Conn.Where("id = ?", id).Delete(&User{}).Error
}
