package entity

import (
	"My-Exercise/global"
	"gorm.io/gorm"
)

type TestCase struct {
	Id        uint   `json:"id"`
	ProblemId uint   `json:"problemId"`
	Input     string `json:"input" gorm:"type:text;comment:输入"`
	Output    string `json:"output" gorm:"type:text;comment:输出"`
	gorm.Model
}

func (t TestCase) TableName() string {
	return "test_case"
}

func AddTestCase(testCase *TestCase) {
	global.DB.Model(new(TestCase)).Create(testCase)
}

func DeleteTestCaseByProblemId(problemId uint) {
	global.DB.Model(new(TestCase)).Where("problem_id = ?", problemId).Delete(new(TestCase))
}
