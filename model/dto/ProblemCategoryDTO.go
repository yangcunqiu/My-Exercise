package dto

import (
	"My-Exercise/model"
)

type ProblemCategoryDTO struct {
	Id        uint
	Title     string `json:"title"`
	Content   string `json:"content"`
	Timeout   int    `json:"timeout"`
	MaxMemory int    `json:"maxMemory"`
	model.BaseInfo
	CategoryList []CategoryDTO `json:"categoryList"`
}
