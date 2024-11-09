package routes

import (
	"dainxor/we/controller"

	"github.com/gin-gonic/gin"
)

func ProjectRoutes(router *gin.Engine) {
	projectRouter := router.Group("api/v0/project")
	{
		projectRouter.POST("/", controller.Project.Create)

		projectRouter.GET("/id/:id", controller.Project.GetByID)
		projectRouter.GET("/all/", controller.Project.GetAll)
		projectRouter.GET("/id-creator/:id", controller.Project.GetByCreatorID)

		projectRouter.PUT("/id/:id", controller.Project.UpdateByID)

		projectRouter.DELETE("/id/:id", controller.Project.DeleteByID)

		settingsRouter := projectRouter.Group("/settings")
		{
			settingsRouter.POST("/", controller.Project.Settings.Create)

			settingsRouter.GET("/id/:id", controller.Project.Settings.GetByID)
			settingsRouter.GET("/all/", controller.Project.Settings.GetAll)

			settingsRouter.PUT("/id/:id", controller.Project.Settings.UpdateByID)

			settingsRouter.DELETE("/id/:id", controller.Project.Settings.DeleteByID)
		}

		collaboratorRouter := projectRouter.Group("/collaborator")
		{
			collaboratorRouter.POST("/", controller.Project.Collaborator.Create)

			collaboratorRouter.GET("/id/:id", controller.Project.Collaborator.GetByID)
			collaboratorRouter.GET("/id-user/:idUser/id-project/:idProject", controller.Project.Collaborator.GetByUserIDAndProjectID)
			collaboratorRouter.GET("/all/", controller.Project.Collaborator.GetAll)
			collaboratorRouter.GET("/id-project/:id", controller.Project.Collaborator.GetByProjectID)
			collaboratorRouter.GET("/id-user/:id", controller.Project.Collaborator.GetByUserID)

			collaboratorRouter.DELETE("/id/:id", controller.Project.Collaborator.DeleteByID)
		}

		permissionRouter := projectRouter.Group("/permission")
		{
			permissionRouter.POST("/", controller.Project.Permission.Create)

			permissionRouter.GET("/id/:id", controller.Project.Permission.GetByID)
			permissionRouter.GET("/id-collaborator/:id", controller.Project.Permission.GetByCollaboratorID)
			permissionRouter.GET("/id-project/:idProject/id-permission/:idPermission", controller.Project.Permission.GetByProjectIDAndPermissionID)

			permissionRouter.DELETE("/id/:id", controller.Project.Permission.DeleteByID)
		}
	}
}
