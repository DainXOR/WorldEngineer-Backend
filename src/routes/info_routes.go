package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InfoRoutes(router *gin.Engine) {
	routes := gin.H{
		"info":               "/api/info/",
		"info ping":          "/api/info/ping",
		"info api version":   "/api/info/api-version",
		"info route version": "/api/info/route-version",

		"info test get":    "/api/test/get",
		"info test post":   "/api/test/post",
		"info test put":    "/api/test/put",
		"info test patch":  "/api/test/patch",
		"info test delete": "/api/test/del",

		"register email":        "/api/v0/user/register/:email",
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
	}
}
