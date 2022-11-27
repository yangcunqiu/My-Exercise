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
				"code":    200,
				"message": "pong!",
			})
		})

		problemGroup := rootGroup.Group("/problem")
		{
			problemGroup.GET("/list", service.ListProblem)
			problemGroup.GET("/info/:id", service.ProblemInfo)

		}

		userGroup := rootGroup.Group("/user")
		{
			userGroup.GET("/info/:id", service.UserInfo)
		}

		submitGroup := rootGroup.Group("/submit")
		{
			submitGroup.POST("/list", service.SubmitList)
		}

	}
}
