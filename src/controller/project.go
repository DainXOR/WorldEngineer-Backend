package controller

import (
	"dainxor/we/base/logger"
	"dainxor/we/db"
	"dainxor/we/models"
	"dainxor/we/types"
	"dainxor/we/utils"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type settingsType struct{}
type collaboratorType struct{}
type permissionType struct{}

type projectType struct {
	Settings     settingsType
	Collaborator collaboratorType
	Permission   permissionType
}

var Project projectType

// > Project

func (projectType) Create(c *gin.Context) {
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

	settingsResult := db.Project.Settings.Create(body.Settings().ToDB())

	if settingsResult.IsErr() {
		c.JSON(settingsResult.Error().Code.AsInt(), settingsResult.Error())
		return
	}

	project = body.ToDB()
	project.IDSettings = settingsResult.Value().ID

	projectResult := db.Project.Create(project)

	if projectResult.IsErr() {
		c.JSON(projectResult.Error().Code.AsInt(), projectResult.Error())
		return
	}

	c.JSON(http.StatusCreated, projectResult.Value().ToResponse())
}

func (projectType) GetByID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.GetByID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}

func (projectType) GetAll(c *gin.Context) {
	result := db.Project.GetAll()

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	projects := result.Value()
	response := utils.Map(projects, models.ProjectDB.ToResponse)

	c.JSON(http.StatusOK, response)
}
func (projectType) GetByCreatorID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.GetByUserID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	projects := result.Value()
	response := utils.Map(projects, models.ProjectDB.ToResponse)

	c.JSON(http.StatusOK, response)
}

func (projectType) UpdateByID(c *gin.Context) {
	var project models.ProjectDB
	var body models.ProjectUpdate
	id := c.Param("id")
	idUint, err := strconv.Atoi(id)

	if err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to update project: ID is invalid")
		logger.Error("Expected ID: ", "uint")

		c.JSON(http.StatusBadRequest, models.Error(
			types.Http.BadRequest(),
			"bad_request",
			"ID is invalid",
		))
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to update project: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{name: string, description: string, id_status: int}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Project.GetByID(id); result.IsErr() {
		c.JSON(result.Error().Code.AsInt(), result.Error())
		return
	}

	project = body.ToDB()
	project.ID = uint(idUint)

	updateResult := db.Project.Update(project)

	if updateResult.IsErr() {
		c.JSON(updateResult.Error().Code.AsInt(), updateResult.Error())
		return
	}

	c.JSON(http.StatusOK, updateResult.Value().ToResponse())
}

func (projectType) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.GetByID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	result = db.Project.Delete(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}

// > Settings

func (settingsType) Create(c *gin.Context) {
	var body models.ProjectSettingsCreate

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to create project settings: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{public: bool}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	settingsResult := db.Project.Settings.Create(body.ToDB())

	if settingsResult.IsErr() {
		c.JSON(settingsResult.Error().Code.AsInt(), settingsResult.Error())
		return
	}

	c.JSON(http.StatusCreated, settingsResult.Value().ToResponse())
}

func (settingsType) GetByID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Settings.GetByID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}

func (settingsType) GetAll(c *gin.Context) {
	result := db.Project.Settings.GetAll()

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	settings := result.Value()
	response := utils.Map(settings, models.ProjectSettingsDB.ToResponse)

	c.JSON(http.StatusOK, response)
}

func (settingsType) UpdateByID(c *gin.Context) {
	var settings models.ProjectSettingsDB
	var body models.ProjectSettingsUpdate
	id := c.Param("id")
	idUint, err := strconv.Atoi(id)

	if err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to update project settings: ID is invalid")
		logger.Error("Expected ID: ", "uint")

		c.JSON(http.StatusBadRequest, models.Error(
			types.Http.BadRequest(),
			"bad_request",
			"ID is invalid",
		))
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to update project settings: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{public: bool}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Project.Settings.GetByID(id); result.IsErr() {
		c.JSON(result.Error().Code.AsInt(), result.Error())
		return
	}

	settings = body.ToDB()
	settings.ID = uint(idUint)

	updateResult := db.Project.Settings.Update(settings)

	if updateResult.IsErr() {
		c.JSON(updateResult.Error().Code.AsInt(), updateResult.Error())
		return
	}

	c.JSON(http.StatusOK, updateResult.Value().ToResponse())
}

func (settingsType) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Settings.GetByID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	result = db.Project.Settings.Delete(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}

// > Collaborator

func (collaboratorType) Create(c *gin.Context) {
	var body models.ProjectCollaboratorCreate

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to create project collaborator: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{id_project: int, id_user: int}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Project.Collaborator.Create(body.ToDB())

	if result.IsErr() {
		c.JSON(result.Error().Code.AsInt(), result.Error())
		return
	}

	c.JSON(http.StatusCreated, result.Value().ToResponse())
}

func (collaboratorType) GetByID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Collaborator.GetByID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}
func (collaboratorType) GetByUserIDAndProjectID(c *gin.Context) {
	idUser := c.Param("idUser")
	idProject := c.Param("idProject")
	result := db.Project.Collaborator.GetByUserIDAndProjectID(idUser, idProject)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}

func (collaboratorType) GetAll(c *gin.Context) {
	result := db.Project.Collaborator.GetAll()

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	collaborators := result.Value()
	response := utils.Map(collaborators, models.ProjectCollaboratorDB.ToResponse)

	c.JSON(http.StatusOK, response)
}
func (collaboratorType) GetByProjectID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Collaborator.GetByProjectID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	collaborators := result.Value()
	response := utils.Map(collaborators, models.ProjectCollaboratorDB.ToResponse)

	c.JSON(http.StatusOK, response)
}
func (collaboratorType) GetByUserID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Collaborator.GetByUserID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	collaborators := result.Value()
	response := utils.Map(collaborators, models.ProjectCollaboratorDB.ToResponse)

	c.JSON(http.StatusOK, response)
}

func (collaboratorType) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Collaborator.GetByID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	result = db.Project.Collaborator.Delete(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}

// > Permission

func (permissionType) Create(c *gin.Context) {
	var body models.CollaboratorPermissionCreate

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to create project permission: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{id_project: int, id_user: int, id_permission: int}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Project.Permission.Create(body.ToDB())

	if result.IsErr() {
		c.JSON(result.Error().Code.AsInt(), result.Error())
		return
	}

	c.JSON(http.StatusCreated, result.Value().ToResponse())
}

func (permissionType) GetByID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Permission.GetByID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}

func (permissionType) GetByCollaboratorID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Permission.GetByCollaboratorID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	permissions := result.Value()
	response := utils.Map(permissions, models.CollaboratorPermissionDB.ToResponse)

	c.JSON(http.StatusOK, response)
}
func (permissionType) GetByProjectIDAndPermissionID(c *gin.Context) {
	idProject := c.Param("idProject")
	idPermission := c.Param("idPermission")
	result := db.Project.Permission.GetByProjectIDAndPermissionID(idPermission, idProject)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	permissions := result.Value()

	response := utils.MMap(permissions, func(_ uint, v []models.CollaboratorPermissionDB) []models.CollaboratorPermissionResponse {
		return utils.Map(v, models.CollaboratorPermissionDB.ToResponse)
	})

	c.JSON(http.StatusOK, response)
}

func (permissionType) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Permission.GetByID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	result = db.Project.Permission.Delete(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}
