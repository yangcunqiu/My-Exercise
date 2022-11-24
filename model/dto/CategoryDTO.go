package dto

import "My-Exercise/model"

type CategoryDTO struct {
	Id       uint
	Name     string
	ParentId int
	model.BaseInfo
}
