package routes

import (
	"dainxor/we/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userRouter := router.Group("api/v0/user")
	{
		userRouter.POST("/", controller.UserCreate)

		userRouter.GET("/all/", controller.UserGetAll)
		userRouter.GET("/id/:id", controller.UserGetByID)
		userRouter.GET("/id-status/:id", controller.UserGetByStatusID)

		userRouter.PUT("/id/:id", controller.UserUpdateByID)

		userRouter.DELETE("/id/:id", controller.UserDeleteByID)

	}
}
