package service

import (
	"My-Exercise/model"
	"My-Exercise/model/entity"
	"My-Exercise/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func ListProblem(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageNum", "10"))
	title := c.Query("title")
	categoryId, _ := strconv.Atoi(c.Query("categoryId"))

	var total int64
	problems := make([]entity.Problem, 0)
	tx := entity.ListProblem(title, categoryId)
	tx.Debug().Omit("content").Count(&total).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&problems)
	problemsResult := make([]entity.Problem, 0)
	for _, item := range problems {
		if categoryId != 0 {
			if len(item.CategoryList) != 0 {
				problemsResult = append(problemsResult, item)
			}
		} else {
			problemsResult = append(problemsResult, item)
		}

	}
	utils.Success(c, model.PageOf(pageNum, pageSize, total, problemsResult))
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
