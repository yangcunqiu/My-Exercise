package model

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Page struct {
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
	List     any   `json:"list"`
}

func PageOf(pageNumber, pageSize int, total int64, list any) Page {
	return Page{
		PageNum:  pageNumber,
		PageSize: pageSize,
		Total:    total,
		List:     list,
	}
}

func PageParams(c *gin.Context) (pageNum, pageSize, offset int) {
	pageNum, _ = strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ = strconv.Atoi(c.DefaultQuery("pageNum", "10"))
	if pageSize > 1000 {
		pageSize = 1000
	}
	offset = (pageNum - 1) * pageSize
	return
}
