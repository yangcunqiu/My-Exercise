package router

import (
	"My-Exercise/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRouter(r *gin.Engine) {
	rootGroup := r.Group("/exercise")
	{
		rootGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "pong!",
			})
		})

		problemGroup := rootGroup.Group("/problem")
		{
			problemGroup.GET("/list", service.ListProblem)

		}

	}
}
