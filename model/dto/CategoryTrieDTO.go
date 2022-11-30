package dto

type CategoryTireDTO struct {
	Id        int
	Name      string
	ParentId  int
	ChildList []*CategoryTireDTO
}
