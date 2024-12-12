package routes

import (
	"dainxor/we/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userRouter := router.Group("api/v0/user")
	{
		userRouter.POST("/", controller.User.Create)

		userRouter.GET("/all/", controller.User.GetAll)
		userRouter.GET("/id/:id", controller.User.GetByID)
		userRouter.GET("/id-status/:id", controller.User.GetAllByStatusID)

		userRouter.PUT("/id/:id", controller.User.UpdateByID)

		userRouter.DELETE("/id/:id", controller.User.DeleteByID)

	}
}
