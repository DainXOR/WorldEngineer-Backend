package routes

import (
	"dainxor/we/test"

	"github.com/gin-gonic/gin"
)

func TestRoutes(router *gin.Engine) {
	testRouter := router.Group("api/test")
	{
		testRouter.GET("/mail/1/", test.SendTestEmail1)
		testRouter.GET("/mail/2/", test.SendTestEmail2)

		testRouter.GET("/get", test.Get)
		testRouter.POST("/post", test.Post)
		testRouter.PUT("/put", test.Put)
		testRouter.PATCH("/patch", test.Patch)
		testRouter.DELETE("/del", test.Delete)
	}
}
