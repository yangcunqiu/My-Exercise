package entity

import "My-Exercise/model"

// ProblemCategory 题目-分类中间表
type ProblemCategory struct {
	Id         uint
	ProblemId  int `gorm:"comment:题目id"`
	CategoryId int `gorm:"comment:分类id"`
	model.BaseInfo
}

func (pc ProblemCategory) TableName() string {
	return "problem_category"
}
