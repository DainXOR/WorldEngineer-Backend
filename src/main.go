package main

import (
	"dainxor/we/configs"
	"dainxor/we/logger"
	"dainxor/we/middleware"
	"dainxor/we/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	logger.Init()

	gin.SetMode(gin.DebugMode)
	configs.ConnectPostgresTest()
}

func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	routes.MainRoutes(router)
	routes.InfoRoutes(router)
	routes.TestRoutes(router)

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.ProjectRoutes(router)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
