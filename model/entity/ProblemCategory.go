package entity

import "gorm.io/gorm"

// ProblemCategory 题目-分类中间表
type ProblemCategory struct {
	gorm.Model
	ProblemId  int `gorm:"comment:题目id"`
	CategoryId int `gorm:"comment:分类id"`
}

func (pc ProblemCategory) TableName() string {
	return "problem_category"
}
