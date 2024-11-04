package routes

import (
	"dainxor/we/controller"

	"github.com/gin-gonic/gin"
)

func UtilRoutes(router *gin.Engine) {
	utilRouter := router.Group("api/v0/util")
	{
		usernameRouter := utilRouter.Group("/username")
		{
			usernameRouter.GET("/create", controller.Util.CreateUsername)
			usernameRouter.GET("/check/:username", controller.Util.CheckUsername)
		}

		nameTagRouter := utilRouter.Group("/name-tag")
		{
			nameTagRouter.GET("/create/:username", controller.Util.CreateNameTag)
			nameTagRouter.GET("/check", controller.Util.CheckNameTag)
		}

	}
}
