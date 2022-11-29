package router

import (
	"My-Exercise/middlewares"
	"My-Exercise/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRouter(r *gin.Engine) {
	rootGroup := r.Group("/exercise")
	authGroup := rootGroup.Group("/auth")
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
		authProblemGroup := authGroup.Group("/problem", middlewares.Authentication(true))
		{
			authProblemGroup.POST("/add", service.AddProblem)
		}

		userGroup := rootGroup.Group("/user")
		{
			userGroup.GET("/info/:id", service.UserInfo)
			userGroup.POST("/login", service.Login)
			userGroup.POST("/register", service.RegisterUser)
			userGroup.GET("/sendVerifyCode", service.SendVerifyCode)
			userGroup.GET("/rank", service.GetUserRankList)
		}

		submitGroup := rootGroup.Group("/submit")
		{
			submitGroup.POST("/list", service.SubmitList)
		}

	}
}
