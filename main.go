package main

import (
	test "dainxor/we/test"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	testRouter := router.Group("/test")
	{
		testRouter.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		testRouter.GET("/get", test.Get)
		testRouter.POST("/post", test.Post)
		testRouter.PUT("/put", test.Put)
		testRouter.DELETE("/del", test.Delete)
		testRouter.PATCH("/patch", test.Patch)
	}

	//v1 := router.Group("/v1")
	//{
	//
	//}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
