package service

import (
	"My-Exercise/model"
	"My-Exercise/model/entity"
	"My-Exercise/model/query"
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

func Login(c *gin.Context) {
	userLogin := new(query.UserLogin)
	_ = c.ShouldBindJSON(userLogin)
	if userLogin.Name == "" || userLogin.Password == "" {
		utils.Fail(c, model.ErrorCodeOf(20003, "用户名或密码不能为空"))
		return
	}
	// 校验是否存在
	tx := entity.GetUserByName(userLogin.Name)
	user := new(entity.User)
	err := tx.First(user).Error
	if err == gorm.ErrRecordNotFound {
		utils.Fail(c, model.ErrorCodeOf(20004, "用户名不存在"))
		return
	}
	// 校验密码
	passwordMD5 := utils.GenerateMD5(userLogin.Password)
	if user.Password != passwordMD5 {
		utils.Fail(c, model.ErrorCodeOf(20005, "用户名或密码不正确"))
		return
	}
	// token
	token, _ := utils.GenerateToken(user.Id, user.Name)
	utils.Success(c, token)
}
