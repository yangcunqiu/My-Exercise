package service

import (
	"My-Exercise/global"
	"My-Exercise/model"
	"My-Exercise/model/entity"
	"My-Exercise/model/query"
	"My-Exercise/utils"
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
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

func SendVerifyCode(c *gin.Context) {
	emailAddr := c.Query("emailAddr")
	if emailAddr == "" {
		utils.Fail(c, model.ErrorCodeOf(20006, "邮箱地址不能为空"))
	}
	// 生成验证码
	code := utils.GenerateRandomNumberToString(6)

	// 存redis
	global.RDB.Set(context.Background(), emailAddr, code, time.Minute*5)

	// 发送邮件
	htmlStr := "<b>您的验证码是: " + code + "<b>"
	err := utils.SendEmail("验证码", htmlStr, emailAddr)
	if err != nil {
		log.Printf("发送验证码失败, to: %v, err: %v", emailAddr, err)
		utils.Fail(c, model.ErrorCodeOf(20008, "发送验证码失败"))
		return
	}
	utils.Success(c, nil)
	return
}

func RegisterUser(c *gin.Context) {
	userRegister := new(query.UserRegister)
	_ = c.ShouldBindJSON(userRegister)

	if userRegister.Name == "" || userRegister.Password == "" || userRegister.Email == "" {
		utils.Fail(c, model.ErrorCodeOf(20009, "用户注册消息不完整"))
	}

	if userRegister.Code == "" {
		utils.Fail(c, model.ErrorCodeOf(20012, "验证码不能为空"))
		return
	}

	// 校验唯一
	userByEmail := new(entity.User)
	err := entity.GetUserByEmail(userRegister.Email).First(userByEmail).Error
	if err != gorm.ErrRecordNotFound {
		utils.Fail(c, model.ErrorCodeOf(20010, "邮箱已被注册"))
		return
	}

	// 校验验证码
	var ctx = context.Background()
	resultCode, _ := global.RDB.Get(ctx, userRegister.Email).Result()
	if userRegister.Code != resultCode {
		utils.Fail(c, model.ErrorCodeOf(20011, "验证码错误"))
		return
	}

	// 保存
	user := &entity.User{
		Name:     userRegister.Name,
		Password: utils.GenerateMD5(userRegister.Password),
		Phone:    userRegister.Phone,
		Email:    userRegister.Email,
	}
	entity.SaveUser(user)

	// 生成token
	token, _ := utils.GenerateToken(user.Id, user.Name)
	utils.Success(c, token)
}
