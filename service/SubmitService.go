package service

import (
	"My-Exercise/model"
	"My-Exercise/model/entity"
	"My-Exercise/model/query"
	"My-Exercise/utils"
	"github.com/gin-gonic/gin"
)

func SubmitList(c *gin.Context) {
	submitListQuery := new(query.SubmitListQuery)
	_ = c.ShouldBindJSON(submitListQuery)

	var total int64
	submitList := make([]entity.Submit, 0)
	pageNum, pageSize, offset := model.PageParams(c)
	tx := entity.ListSubmit(*submitListQuery)
	tx.Count(&total).Offset(offset).Limit(pageSize).Find(&submitList)

	utils.Success(c, model.PageOf(pageNum, pageSize, total, submitList))
}
