package routes

import (
	"dainxor/we/controller"

	"github.com/gin-gonic/gin"
)

func ProjectRoutes(router *gin.Engine) {
	projectRouter := router.Group("api/v0/project")
	{
		projectRouter.GET("/all/", controller.User.GetAll)
		projectRouter.GET("/id/:id", controller.User.GetByID)
		projectRouter.GET("/id-status/:id", controller.User.GetAllByStatusID)

		projectRouter.PUT("/id/:id", controller.User.UpdateByID)

		projectRouter.DELETE("/id/:id", controller.User.DeleteByID)

	}
}
