package entity

import "My-Exercise/model"

// Category 分类表
type Category struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" gorm:"type:varchar(100);comment:分类名称"`
	ParentId int    `json:"parentId" gorm:"comment:父分类id"`
	model.BaseInfo
	ProblemList []*Problem `json:"problemList" gorm:"many2many:problem_category"`
}

func (c Category) TableName() string {
	return "category"
}
