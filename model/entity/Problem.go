package entity

import (
	"My-Exercise/global"
	"My-Exercise/model"
	"My-Exercise/model/dto"
	"gorm.io/gorm"
)

// Problem 题目表
type Problem struct {
	Id        uint   `json:"id"`
	Title     string `json:"title" gorm:"type:varchar(200);comment:题目标题"`
	Content   string `json:"content" gorm:"type:text;comment:题目内容"`
	Timeout   int    `json:"timeout" gorm:"comment:超时时间(ms)"`
	MaxMemory int    `json:"maxMemory" gorm:"comment:最大内存(kb)"`
	model.BaseInfo
}

func (p Problem) TableName() string {
	return "problem"
}

func ListProblem(title string, categoryId int) []dto.ProblemCategoryDTO {
	// TODO 原生sql
	problemCategoryDTOList := make([]dto.ProblemCategoryDTO, 0)
	problemCategoryDTO := new(dto.ProblemCategoryDTO)
	rows, _ := global.DB.Raw("select * from problem p left join problem_category pc on p.id = pc.problem_id left join category c on c.id = pc.category_id where p.deleted_at is null").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(problemCategoryDTO)
	}
	return problemCategoryDTOList
}

func GetProblemById(id int) *gorm.DB {
	return global.DB.Model(new(Problem)).Where("id = ?", id).Preload("CategoryList")
}

func AddProblem(problem *Problem) {
	global.DB.Model(new(Problem)).Create(problem)
}

func CountProblemByCategoryId(categoryId int) int64 {
	var count int64
	global.DB.Debug().Model(new(Problem)).Joins("left join problem_category pc on pc.category_id = problem.id").Where("pc.category_id = ?", categoryId).Where("pc.deleted_at is null").Count(&count)
	return count
}
