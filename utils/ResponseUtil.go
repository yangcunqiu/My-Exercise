package utils

import (
	"My-Exercise/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, model.Result{
		Code:    200,
		Message: "成功",
		Data:    data,
	})
}

func Fail(c *gin.Context, errorCode model.ErrorCode) {
	c.JSON(http.StatusOK, model.Result{
		Code:    errorCode.Code,
		Message: errorCode.Message,
	})
}
