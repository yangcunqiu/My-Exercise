package service

import (
	"My-Exercise/model"
	"My-Exercise/model/dto"
	"My-Exercise/model/entity"
	"My-Exercise/model/query"
	"My-Exercise/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func ListCategory(c *gin.Context) {
	name := c.Query("name")

	var total int64
	pageNum, pageSize, offset := model.PageParams(c)
	tx := entity.CategoryList(name)
	categoryList := make([]entity.Category, 0)
	tx.Count(&total).Offset(offset).Limit(pageSize).Find(&categoryList)

	categoryTireDTOList := make([]*dto.CategoryTireDTO, 0)
	for _, category := range categoryList {
		categoryTireDTO := &dto.CategoryTireDTO{
			Id:       int(category.Id),
			Name:     category.Name,
			ParentId: int(category.ParentId),
		}
		getCategoryChildListByParentId(categoryTireDTO)
		categoryTireDTOList = append(categoryTireDTOList, categoryTireDTO)
	}
	utils.Success(c, model.PageOf(pageNum, pageSize, total, categoryTireDTOList))
}

func getCategoryChildListByParentId(categoryTire *dto.CategoryTireDTO) {
	childList := make([]*dto.CategoryTireDTO, 0)
	tx := entity.GetCategoryByParentId(categoryTire.Id)
	tx.First(&childList)

	if len(childList) == 0 {
		return
	}

	for _, tire := range childList {
		getCategoryChildListByParentId(tire)
	}
	categoryTire.ChildList = childList
}

func AddCategory(c *gin.Context) {
	categorySave := new(query.CategorySave)
	_ = c.ShouldBindJSON(categorySave)
	if categorySave.Name == "" {
		utils.Fail(c, model.ErrorCodeOf(60001, "分类名称不能为空"))
		return
	}

	category := &entity.Category{
		Name:     categorySave.Name,
		ParentId: categorySave.ParentId,
	}

	entity.AddCategory(category)
	utils.Success(c, nil)
}

func UpdateCategory(c *gin.Context) {
	categorySave := new(query.CategorySave)
	_ = c.ShouldBindJSON(categorySave)
	if categorySave.Id == 0 {
		utils.Fail(c, model.ErrorCodeOf(60002, "分类id不能为空"))
		return
	}
	if categorySave.Name == "" {
		utils.Fail(c, model.ErrorCodeOf(60003, "分类名称不能为空"))
		return
	}

	category := &entity.Category{
		Id:       categorySave.Id,
		Name:     categorySave.Name,
		ParentId: categorySave.ParentId,
	}
	entity.UpdateCategory(category)
	utils.Success(c, nil)
}

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	category := new(entity.Category)
	tx := entity.GetCategoryById(id)
	err := tx.First(category).Error
	if err == gorm.ErrRecordNotFound {
		utils.Fail(c, model.ErrorCodeOf(60004, "分类不存在"))
		return
	}

	problemCount := entity.CountProblemByCategoryId(id)
	if problemCount > 0 {
		utils.Fail(c, model.ErrorCodeOf(60005, "分类下已存在问题, 不能删除"))
		return
	}

	categoryChildList := make([]entity.Category, 0)
	tx = entity.GetCategoryByParentId(id)
	tx.Find(&categoryChildList)
	if len(categoryChildList) > 0 {
		utils.Fail(c, model.ErrorCodeOf(60006, "分类下已存在子分类, 不能删除"))
		return
	}

	entity.DeleteCategory(id)
	utils.Success(c, nil)
}
