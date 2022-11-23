package entity

import "gorm.io/gorm"

// User 用户表
type User struct {
	gorm.Model
	Name     string `gorm:"comment:用户名"`
	Password string `gorm:"comment:密码"`
	Phone    string `gorm:"comment:手机号"`
	Mail     string `gorm:"comment:邮箱"`
}

func (u User) TableName() string {
	return "user"
}
