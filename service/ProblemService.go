package service

import (
	"My-Exercise/model"
	"My-Exercise/model/entity"
	"github.com/gin-gonic/gin"
	"net/http"
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
		if len(item.CategoryList) != 0 {
			problemsResult = append(problemsResult, item)
		}
	}

	c.JSON(http.StatusOK, model.Result{
		Code:    200,
		Message: "成功",
		Data: model.Page{
			PageNum:  pageNum,
			PageSize: pageSize,
			Total:    total,
			List:     problemsResult,
		},
	})
}
