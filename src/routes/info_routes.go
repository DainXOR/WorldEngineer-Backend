package routes

import (
	"dainxor/we/test"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InfoRoutes(router *gin.Engine) {
	routes := gin.H{
		"info":               "/api/info/",
		"info ping":          "/api/info/ping",
		"info api version":   "/api/info/api-version",
		"info route version": "/api/info/route-version",
		"info test get":      "/api/info/get",
		"info test post":     "/api/info/post",
		"info test put":      "/api/info/put",
		"info test patch":    "/api/info/patch",
		"info test delete":   "/api/info/del",

		"user create":           "/api/v0/user/",
		"user get all":          "/api/v0/user/all/",
		"user get by id":        "/api/v0/user/id/:id",
		"user get by status id": "/api/v0/user/id-status/:id",
		"user update by id":     "/api/v0/user/id/:id",
		"user delete by id":     "/api/v0/user/id/:id",
	}

	infoRouter := router.Group("api/info")
	{
		infoRouter.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Available routes",
				"routes":  routes,
			})
		})
		infoRouter.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		infoRouter.GET("/api-version", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"version": "0.1.0",
			})
		})
		infoRouter.GET("/route-version", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"version": "v0",
			})
		})

		infoRouter.GET("/get", test.Get)
		infoRouter.POST("/post", test.Post)
		infoRouter.PUT("/put", test.Put)
		infoRouter.PATCH("/patch", test.Patch)
		infoRouter.DELETE("/del", test.Delete)
	}
}
