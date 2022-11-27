package service

import (
	"My-Exercise/model"
	"My-Exercise/model/entity"
	"My-Exercise/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func UserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		utils.Fail(c, model.ErrorCodeOf(20001, "用户id不能为空"))
		return
	}
	user := new(entity.User)
	tx := entity.GetUserById(id)
	err := tx.Omit("password").First(user).Error
	if err == gorm.ErrRecordNotFound {
		utils.Fail(c, model.ErrorCodeOf(20002, "用户不存在"))
		return
	}
	utils.Success(c, user)
}
