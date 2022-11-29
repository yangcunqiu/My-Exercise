package middlewares

import (
	"My-Exercise/model"
	"My-Exercise/model/entity"
	"My-Exercise/utils"
	"github.com/gin-gonic/gin"
)

func Authentication(isAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			utils.Fail(c, model.ErrorCodeOf(40001, "token不能为空"))
			c.Abort()
			return
		}
		userClaims, ok := utils.ParseToken(auth)
		if !ok {
			utils.Fail(c, model.ErrorCodeOf(40002, "token验证失败"))
			c.Abort()
			return
		}
		user := new(entity.User)
		entity.GetUserById(int(userClaims.Id)).First(user)
		c.Set("user", user)
		if isAdmin && !user.Admin {
			utils.Fail(c, model.ErrorCodeOf(40003, "暂无权限"))
			c.Abort()
			return
		}
		c.Next()
	}
}
