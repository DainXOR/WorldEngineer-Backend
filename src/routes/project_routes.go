package routes

import (
	"dainxor/we/controller"

	"github.com/gin-gonic/gin"
)

func ProjectRoutes(router *gin.Engine) {
	projectRouter := router.Group("api/v0/user")
	{
		projectRouter.POST("/register/:email", controller.UserTryRegister)
		projectRouter.POST("/login/:email", controller.UserTryLogin)
		projectRouter.POST("/", controller.UserCreate)

		projectRouter.GET("/all/", controller.UserGetAll)
		projectRouter.GET("/id/:id", controller.UserGetByID)
		projectRouter.GET("/id-status/:id", controller.UserGetByStatusID)

		projectRouter.PUT("/id/:id", controller.UserUpdateByID)

		projectRouter.DELETE("/id/:id", controller.UserDeleteByID)

	}
}
