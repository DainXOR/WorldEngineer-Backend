package controller

import (
	"dainxor/we/configs"
	"dainxor/we/logger"
	"dainxor/we/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func ProjectCreate(c *gin.Context) {
	var project models.ProjectDB
	var body models.ProjectCreate

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to create project: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{name: string, description: string, id_creator: int, public: bool}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var projectSettings models.ProjectSettingsDB
	newProjectSettings := models.ProjectSettingsCreate{
		Public: body.Public,
	}

	// Manually create project settings on the database
	projectSettings.Public = newProjectSettings.Public
	projectSettings.CreatedAt = configs.DB.NowFunc()
	projectSettings.UpdatedAt = configs.DB.NowFunc()

	configs.DB.Create(&projectSettings)

	// Send a temporary redirect to the project settings POST endpoint
	//c.JSON(http.StatusTemporaryRedirect, newProjectSettings)
	//c.Redirect(http.StatusTemporaryRedirect, "/project_settings")
	//
	//if c.Request.Response.StatusCode != http.StatusCreated {
	//	logger.Error("Failed to create project settings")
	//	return
	//}
	//
	//if err := c.ShouldBindJSON(&projectSettings); err != nil {
	//	logger.Error(err.Error())
	//	logger.Error("Failed to create project: JSON request body is invalid")
	//	logger.Error("Request body: ", c.Request.Body)
	//	logger.Error("Expected body: ", "{name: string, description: string, id_creator: int, public: bool}")
	//
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	project.Name = body.Name
	project.Description = body.Description
	project.IDCreator = body.IDCreator
	project.CreatedAt = configs.DB.NowFunc()
	project.UpdatedAt = configs.DB.NowFunc()
	project.IDStatus = 1
	project.StatusTimeStamp = configs.DB.NowFunc()
	project.IDSettings = projectSettings.ID

	configs.DB.Create(&project)

	c.JSON(http.StatusCreated,
		models.ProjectResponse{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
			CreatedAt:   project.CreatedAt.String(),
			IDCreator:   project.IDCreator,
			IDStatus:    project.IDStatus,
			IDSettings:  project.IDSettings,
		},
	)
}
