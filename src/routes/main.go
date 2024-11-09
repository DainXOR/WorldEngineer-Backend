package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Go to /api/info/ to see available routes",
		})
	})
}
