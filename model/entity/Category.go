package entity

import (
	"My-Exercise/global"
	"My-Exercise/model"
	"gorm.io/gorm"
)

// Category 分类表
type Category struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" gorm:"type:varchar(100);comment:分类名称"`
	ParentId uint   `json:"parentId" gorm:"comment:父分类id"`
	model.BaseInfo
}

func (c Category) TableName() string {
	return "category"
}

func CategoryList(name string) *gorm.DB {
	tx := global.DB.Model(new(Category)).Where("parent_id = 0")
	if name != "" {
		tx.Where("name like ?", "%"+name+"%")
	}
	return tx
}

func AddCategory(category *Category) {
	global.DB.Model(new(Category)).Create(category)
}

func UpdateCategory(category *Category) {
	global.DB.Model(new(Category)).Where("id = ?", category.Id).Updates(category)
}

func GetCategoryById(id int) *gorm.DB {
	return global.DB.Model(new(Category)).Where("id = ?", id)
}

func GetCategoryByParentId(parentId int) *gorm.DB {
	return global.DB.Model(new(Category)).Where("parent_id = ?", parentId)
}

func DeleteCategory(id int) {
	global.DB.Delete(new(Category), id)
}
