package main

import (
	"dainxor/we/base/configs"
	"dainxor/we/base/logger"
	"dainxor/we/db"
	"dainxor/we/middleware"
	"dainxor/we/routes"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var address = "localhost:8080"

func Init() {
	logger.Init()
	logger.Info("Loading configurations")

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")

	}

	logger.EnvInit()
	configs.DB.EnvInit()
	address = os.Getenv("ADDRESS")
	db.Mail.LoadCredentials()

	logger.Info("Starting server")
}

func main() {
	Init()

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	routes.MainRoutes(router)
	routes.InfoRoutes(router)
	routes.TestRoutes(router)

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.ProjectRoutes(router)

	router.Run(address) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
