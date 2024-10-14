package main

import (
	"dainxor/we/configs"
	"dainxor/we/models"
)

func init() {
	configs.ConnectPostgresTest()
}

func main() {
	configs.DB.AutoMigrate(&models.UserDB{})
	configs.DB.AutoMigrate(&models.ProjectDB{})
}
