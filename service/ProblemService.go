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

	var total int64
	problems := entity.ListProblem(title)

	c.JSON(http.StatusOK, model.Result{
		Code:    200,
		Message: "成功",
		Data: model.Page{
			PageNum:  pageNum,
			PageSize: pageSize,
			Total:    total,
			List:     problems,
		},
	})
}
