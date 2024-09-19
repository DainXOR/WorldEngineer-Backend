package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get test",
	})
}

func Post(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "post test",
	})
}

func Put(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "put test",
	})
}

func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "delete test",
	})
}

func Patch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "patch test",
	})
}
