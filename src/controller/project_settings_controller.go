package controller

import (
	"dainxor/we/configs"
	"dainxor/we/logger"
	"dainxor/we/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProjectSettings(c *gin.Context) {
	var projectSettings models.ProjectSettingsDB
	var body models.ProjectSettingsCreate

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to create project settings: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{id_project: int, id_settings: int, settings: string}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projectSettings.Public = body.Public
	projectSettings.CreatedAt = configs.DB.NowFunc()
	projectSettings.UpdatedAt = configs.DB.NowFunc()

	configs.DB.Create(&projectSettings)

	c.JSON(http.StatusCreated,
		models.ProjectSettingsResponse{
			Public: projectSettings.Public,
		},
	)

}