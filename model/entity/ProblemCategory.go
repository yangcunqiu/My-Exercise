package entity

import (
	"My-Exercise/global"
	"gorm.io/gorm"
)

type ProblemCategory struct {
	Id         uint
	ProblemId  uint
	CategoryId uint
	gorm.Model
}

func (pc ProblemCategory) TableName() string {
	return "problem_category"
}

func AddProblemCategory(problemCategory *ProblemCategory) {
	global.DB.Model(new(ProblemCategory)).Create(problemCategory)
}

func DeleteProblemCategoryByProblemId(problemId uint) {
	global.DB.Model(new(ProblemCategory)).Where("problem_id = ?", problemId).Delete(new(ProblemCategory))
}
