package service

import (
	"My-Exercise/global"
	"My-Exercise/model/entity"
)

func GetTestCaseByProblemId(problemId int) []entity.TestCase {
	testCaseList := make([]entity.TestCase, 0)
	global.DB.Model(new(entity.TestCase)).Where("problem_id = ?", problemId).Find(&testCaseList)
	return testCaseList
}
