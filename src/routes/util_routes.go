package routes

import (
	"dainxor/we/controller"

	"github.com/gin-gonic/gin"
)

func UtilRoutes(router *gin.Engine) {
	utilRouter := router.Group("api/v0/util")
	{
		nameTagRouter := utilRouter.Group("/name-tag")
		{
			nameTagRouter.GET("/create/:username", controller.Util.CreateNameTag)
			nameTagRouter.GET("/available", controller.Util.AvailableNameTag)
		}

	}
}
