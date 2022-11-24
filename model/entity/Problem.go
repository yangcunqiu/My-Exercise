package entity

import (
	"My-Exercise/global"
	"My-Exercise/model"
	"My-Exercise/model/dto"
)

// Problem 题目表
type Problem struct {
	Id        uint
	Title     string `json:"title" gorm:"type:varchar(200);comment:题目标题"`
	Content   string `json:"content" gorm:"type:text;comment:题目内容"`
	Timeout   int    `json:"timeout" gorm:"comment:超时时间(ms)"`
	MaxMemory int    `json:"maxMemory" gorm:"comment:最大内存(kb)"`
	model.BaseInfo
}

func (p Problem) TableName() string {
	return "problem"
}

func ListProblem(title string) []dto.ProblemCategoryDTO {
	problems := make([]dto.ProblemCategoryDTO, 0)
	global.DB.Debug().
		Raw("select p.* as categoryList.id from problem p left join problem_category pc on p.id = pc.problem_id left join category c on c.id = pc.category_id where p.title like concat('%', ?, '%')", title).
		Scan(&problems)
	return problems
}
