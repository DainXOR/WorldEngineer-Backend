package routes

import (
	"dainxor/we/mail"

	"github.com/gin-gonic/gin"
)

func TestRoutes(router *gin.Engine) {
	testRouter := router.Group("api/test")
	{
		testRouter.GET("/mail/1/", mail.SendTestEmail1)
		testRouter.GET("/mail/2/", mail.SendTestEmail2)

	}
}
