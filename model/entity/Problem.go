package entity

import (
	"My-Exercise/global"
	"My-Exercise/model"
	"gorm.io/gorm"
)

// Problem 题目表
type Problem struct {
	Id        uint
	Title     string `json:"title" gorm:"type:varchar(200);comment:题目标题"`
	Content   string `json:"content" gorm:"type:text;comment:题目内容"`
	Timeout   int    `json:"timeout" gorm:"comment:超时时间(ms)"`
	MaxMemory int    `json:"maxMemory" gorm:"comment:最大内存(kb)"`
	model.BaseInfo
	CategoryList []*Category `json:"categoryList" gorm:"many2many:problem_category"`
}

func (p Problem) TableName() string {
	return "problem"
}

func ListProblem(title string, categoryId int) *gorm.DB {
	tx := global.DB.Model(new(Problem))
	if title != "" {
		tx.Where("title like ?", "%"+title+"%")
	}
	if categoryId != 0 {
		tx.Preload("CategoryList", "id = ?", categoryId)
	} else {
		tx.Preload("CategoryList")
	}
	return tx
}
