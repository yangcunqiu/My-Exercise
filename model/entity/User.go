package entity

import "My-Exercise/model"

// User 用户表
type User struct {
	Id       uint
	Name     string `gorm:"type:varchar(100);comment:用户名"`
	Password string `gorm:"type:varchar(255);comment:密码"`
	Phone    string `gorm:"type:varchar(20);comment:手机号"`
	Mail     string `gorm:"type:varchar(50);comment:邮箱"`
	model.BaseInfo
}

func (u User) TableName() string {
	return "user"
}
