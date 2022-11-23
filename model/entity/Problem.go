package entity

import "gorm.io/gorm"

// Problem 题目表
type Problem struct {
	gorm.Model
	Title     string `gorm:"comment:题目标题"`
	Content   string `gorm:"comment:题目内容"`
	Timeout   int    `gorm:"comment:超时时间(ms)"`
	MaxMemory int    `gorm:"comment:最大内存(kb)"`
}

func (p Problem) TableName() string {
	return "problem"
}
