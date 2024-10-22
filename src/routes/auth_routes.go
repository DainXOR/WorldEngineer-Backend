package routes

import (
	"dainxor/we/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	authRouter := router.Group("api/v0/auth")
	{
		authRouter.GET("/register/:email", controller.Auth.Register)
		authRouter.GET("/login/:email", controller.Auth.Login)
		authRouter.GET("/auth/:email", controller.Auth.Authenticate)
	}
}
