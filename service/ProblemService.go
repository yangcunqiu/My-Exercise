package service

import (
	"My-Exercise/model"
	"My-Exercise/model/entity"
	"My-Exercise/model/query"
	"My-Exercise/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func ListProblem(c *gin.Context) {
	title := c.Query("title")
	categoryId, _ := strconv.Atoi(c.Query("categoryId"))
	pageNum, pageSize, _ := model.PageParams(c)

	var total int64
	//problems := make([]entity.Problem, 0)
	problemCategoryDTO := entity.ListProblem(title, categoryId)
	//tx.Debug().Omit("content").Count(&total).Offset(offset).Limit(pageSize).Find(&problems)
	//problemsResult := make([]entity.Problem, 0)
	//for _, item := range problems {
	//	if categoryId != 0 {
	//		if len(item.CategoryList) != 0 {
	//			problemsResult = append(problemsResult, item)
	//		}
	//	} else {
	//		problemsResult = append(problemsResult, item)
	//	}
	//
	//}
	utils.Success(c, model.PageOf(pageNum, pageSize, total, problemCategoryDTO))
}

func ProblemInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		utils.Fail(c, model.ErrorCodeOf(10001, "问题id不能为空"))
		return
	}
	problem := new(entity.Problem)
	tx := entity.GetProblemById(id)
	err := tx.First(&problem).Error
	if err == gorm.ErrRecordNotFound {
		utils.Fail(c, model.ErrorCodeOf(10002, "问题不存在"))
		return
	}
	utils.Success(c, problem)
}

func AddProblem(c *gin.Context) {
	problemSave := new(query.ProblemSave)
	_ = c.ShouldBindJSON(problemSave)
	if len(problemSave.CategoryIdList) <= 0 || len(problemSave.TestCaseList) <= 0 || problemSave.Title == "" {
		utils.Fail(c, model.ErrorCodeOf(10003, "参数不完整"))
		return
	}

	// 问题
	problem := &entity.Problem{
		Title:     problemSave.Title,
		Content:   problemSave.Content,
		Timeout:   problemSave.Timeout,
		MaxMemory: problemSave.MaxMemory,
	}
	entity.AddProblem(problem)

	// 分类
	for _, id := range problemSave.CategoryIdList {
		problemCategory := &entity.ProblemCategory{
			ProblemId:  problem.Id,
			CategoryId: id,
		}
		entity.AddProblemCategory(problemCategory)
	}

	// 测试用例
	for _, testCaseSave := range problemSave.TestCaseList {
		testCase := &entity.TestCase{
			ProblemId: problem.Id,
			Input:     testCaseSave.Input,
			Output:    testCaseSave.Output,
		}
		entity.AddTestCase(testCase)
	}
	utils.Success(c, nil)
}
