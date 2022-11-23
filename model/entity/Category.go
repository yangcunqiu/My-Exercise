package entity

import "gorm.io/gorm"

// Category 分类表
type Category struct {
	gorm.Model
	Name     string `gorm:"comment:分类名称"`
	ParentId int    `gorm:"comment:父分类id"`
}

func (c Category) TableName() string {
	return "category"
}
