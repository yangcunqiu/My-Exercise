package dto

import "My-Exercise/model"

type ProblemCategory struct {
	Id           uint
	Title        string `json:"title"`
	Content      string `json:"content"`
	CategoryId   uint   `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	Timeout      int    `json:"timeout"`
	MaxMemory    int    `json:"maxMemory"`
	model.BaseInfo
}
