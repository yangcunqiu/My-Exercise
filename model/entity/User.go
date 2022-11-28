package entity

import (
	"My-Exercise/global"
	"My-Exercise/model"
	"gorm.io/gorm"
)

// User 用户表
type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" gorm:"type:varchar(100);comment:用户名"`
	Password string `json:"password" gorm:"type:varchar(255);comment:密码"`
	Phone    string `json:"phone" gorm:"type:varchar(20);comment:手机号"`
	Email    string `json:"email" gorm:"type:varchar(50);comment:邮箱"`
	model.BaseInfo
}

func (u User) TableName() string {
	return "user"
}

func GetUserById(id int) *gorm.DB {
	return global.DB.Model(new(User)).Where("id = ?", id)
}

func GetUserByName(name string) *gorm.DB {
	return global.DB.Where("name = ?", name)
}

func GetUserByEmail(email string) *gorm.DB {
	return global.DB.Where("email = ?", email)
}
