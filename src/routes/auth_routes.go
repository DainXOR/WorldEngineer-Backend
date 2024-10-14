package routes

import (
	"dainxor/we/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	authRouter := router.Group("api/v0/auth")
	{
		authRouter.GET("/register/:email", controller.UserTryRegister)
		authRouter.GET("/login/:email", controller.UserTryLogin)
		authRouter.GET("/auth/:email", controller.UserAuth)
	}
}
