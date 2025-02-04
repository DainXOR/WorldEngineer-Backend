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
type resourcesType struct{}
type characterType struct{}
type locationType struct{}
type elementType struct{}

type projectType struct {
	Settings     settingsType
	Collaborator collaboratorType
	Permission   permissionType
	Resources    resourcesType
	Character    characterType
	Location     locationType
	StoryElement elementType
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
func (collaboratorType) GetByProjectIDAndPermissionID(c *gin.Context) {
	idProject := c.Param("idProject")
	idPermission := c.Param("idPermission")
	result := db.Project.Collaborator.GetByProjectIDAndPermissionID(idPermission, idProject)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	collaborators := result.Value()

	response := utils.Map(collaborators, models.ProjectCollaboratorDB.ToResponse)

	c.JSON(http.StatusOK, response)
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
	id := c.Param("idUser")
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

// > Resources

func (resourcesType) CreateText(c *gin.Context) {
	var body models.ResourceTextCreate

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to create text resource: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{id_project: int, id_creator: int, content: string}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	textResource := models.ResourceTextDB{
		ResourceDB: body.ToDB(),
	}

	result := db.Project.Resources.CreateText(textResource)

	if result.IsErr() {
		c.JSON(result.Error().Code.AsInt(), result.Error())
		return
	}

	c.JSON(http.StatusCreated, result.Value().ToResponse())
}

func (resourcesType) GetTextByID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Resources.GetTextByID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}

func (resourcesType) GetTextByProjectID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Resources.GetTextByProjectID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	resources := result.Value()
	response := utils.Map(resources, models.ResourceTextDB.ToResponse)

	c.JSON(http.StatusOK, response)
}

func (resourcesType) UpdateTextByID(c *gin.Context) {
	var resource models.ResourceTextDB
	var body models.ResourceTextUpdate
	id := c.Param("id")
	idUint, err := strconv.Atoi(id)

	if err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to update text resource: ID is invalid")
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
		logger.Error("Failed to update text resource: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{content: string}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Project.Resources.GetTextByID(id); result.IsErr() {
		c.JSON(result.Error().Code.AsInt(), result.Error())
		return
	}

	resource = models.ResourceTextDB{ResourceDB: body.ToDB()}
	resource.ID = uint(idUint)

	updateResult := db.Project.Resources.UpdateText(resource)

	if updateResult.IsErr() {
		c.JSON(updateResult.Error().Code.AsInt(), updateResult.Error())
		return
	}

	c.JSON(http.StatusOK, updateResult.Value().ToResponse())
}

// > Character

func (characterType) Create(c *gin.Context) {
	var body models.ProjectCharacterCreate

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to create project character: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{id_project: int, name: string, description: string}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Project.Character.Create(body.ToDB())

	if result.IsErr() {
		c.JSON(result.Error().Code.AsInt(), result.Error())
		return
	}

	c.JSON(http.StatusCreated, result.Value().ToResponse())
}

func (characterType) GetByID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Character.GetByID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}

func (characterType) GetByProjectID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Character.GetByProjectID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	characters := result.Value()
	response := utils.Map(characters, models.ProjectCharacterDB.ToResponse)

	c.JSON(http.StatusOK, response)
}

func (characterType) UpdateByID(c *gin.Context) {
	var character models.ProjectCharacterDB
	var body models.ProjectCharacterUpdate
	id := c.Param("id")
	idUint, err := strconv.Atoi(id)

	if err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to update project character: ID is invalid")
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
		logger.Error("Failed to update project character: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{name: string, description: string}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Project.Character.GetByID(id); result.IsErr() {
		c.JSON(result.Error().Code.AsInt(), result.Error())
		return
	}

	character = body.ToDB()
	character.ID = uint(idUint)

	updateResult := db.Project.Character.Update(character)

	if updateResult.IsErr() {
		c.JSON(updateResult.Error().Code.AsInt(), updateResult.Error())
		return
	}

	c.JSON(http.StatusOK, updateResult.Value().ToResponse())
}

func (characterType) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Character.GetByID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	result = db.Project.Character.Delete(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}

// > Character Relation

func (characterType) CreateRelation(c *gin.Context) {
	var body models.ProjectCharacterRelationCreate

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to create project character relation: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{id_project: int, id_character_one: int, id_character_two: int, id_type: int, name: string}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Project.Character.CreateRelation(body.ToDB())

	if result.IsErr() {
		c.JSON(result.Error().Code.AsInt(), result.Error())
		return
	}

	c.JSON(http.StatusCreated, result.Value().ToResponse())
}

func (characterType) GetRelationByID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Character.GetRelationByID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}
func (characterType) GetRelationByCharacterIDs(c *gin.Context) {
	idCharacterOne := c.Param("idCharacterOne")
	idCharacterTwo := c.Param("idCharacterTwo")
	result := db.Project.Character.GetRelationByCharacterIDs(idCharacterOne, idCharacterTwo)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}

func (characterType) GetRelationsByProjectID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Character.GetRelationsByProjectID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	relations := result.Value()
	response := utils.Map(relations, models.ProjectCharacterRelationDB.ToResponse)

	c.JSON(http.StatusOK, response)
}
func (characterType) GetRelationsByCharacterID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Character.GetRelationsByCharacterID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	relations := result.Value()
	response := utils.Map(relations, models.ProjectCharacterRelationDB.ToResponse)

	c.JSON(http.StatusOK, response)
}

func (characterType) UpdateRelationByID(c *gin.Context) {
	var relation models.ProjectCharacterRelationDB
	var body models.ProjectCharacterRelationUpdate
	id := c.Param("id")
	idUint, err := strconv.Atoi(id)

	if err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to update project character relation: ID is invalid")
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
		logger.Error("Failed to update project character relation: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{name: string}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Project.Character.GetRelationByID(id); result.IsErr() {
		logger.Error(result.Error().Code.AsInt())
		logger.Error(result.Error().Message)
		c.JSON(result.Error().Code.AsInt(), result.Error())
		return
	}

	relation = body.ToDB()
	relation.ID = uint(idUint)

	updateResult := db.Project.Character.UpdateRelation(relation)

	if updateResult.IsErr() {
		c.JSON(updateResult.Error().Code.AsInt(), updateResult.Error())
		return
	}

	c.JSON(http.StatusOK, updateResult.Value().ToResponse())
}

func (characterType) DeleteRelationByID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Character.GetRelationByID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	result = db.Project.Character.DeleteRelation(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}

// > Location

func (locationType) Create(c *gin.Context) {
	var body models.ProjectLocationCreate

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to create project location: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{id_project: int, name: string, description: string}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Project.Location.Create(body.ToDB())

	if result.IsErr() {
		c.JSON(result.Error().Code.AsInt(), result.Error())
		return
	}

	c.JSON(http.StatusCreated, result.Value().ToResponse())
}

func (locationType) GetByID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Location.GetByID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}

func (locationType) GetByProjectID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Location.GetByProjectID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	locations := result.Value()
	response := utils.Map(locations, models.ProjectLocationDB.ToResponse)

	c.JSON(http.StatusOK, response)
}

func (locationType) UpdateByID(c *gin.Context) {
	var location models.ProjectLocationDB
	var body models.ProjectLocationUpdate
	id := c.Param("id")
	idUint, err := strconv.Atoi(id)

	if err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to update project location: ID is invalid")
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
		logger.Error("Failed to update project location: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{name: string, description: string}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Project.Location.GetByID(id); result.IsErr() {
		c.JSON(result.Error().Code.AsInt(), result.Error())
		return
	}

	location = body.ToDB()
	location.ID = uint(idUint)

	updateResult := db.Project.Location.Update(location)

	if updateResult.IsErr() {
		c.JSON(updateResult.Error().Code.AsInt(), updateResult.Error())
		return
	}

	c.JSON(http.StatusOK, updateResult.Value().ToResponse())
}

func (locationType) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.Location.GetByID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	result = db.Project.Location.Delete(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}

// > Story Element

func (elementType) Create(c *gin.Context) {
	var body models.ProjectStoryElementCreate

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to create project story element: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{id_project: int, name: string, description: string}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Project.StoryElement.Create(body.ToDB())

	if result.IsErr() {
		c.JSON(result.Error().Code.AsInt(), result.Error())
		return
	}

	c.JSON(http.StatusCreated, result.Value().ToResponse())
}

func (elementType) GetByID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.StoryElement.GetByID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}
func (elementType) GetByName(c *gin.Context) {
	name := c.Param("name")
	result := db.Project.StoryElement.GetByName(name)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}

func (elementType) GetByProjectID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.StoryElement.GetByProjectID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	elements := result.Value()
	response := utils.Map(elements, models.ProjectStoryElementDB.ToResponse)

	c.JSON(http.StatusOK, response)
}

func (elementType) UpdateByID(c *gin.Context) {
	var element models.ProjectStoryElementDB
	var body models.ProjectStoryElementUpdate
	id := c.Param("id")
	idUint, err := strconv.Atoi(id)

	if err != nil {
		logger.Error(err.Error())
		logger.Error("Failed to update project story element: ID is invalid")
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
		logger.Error("Failed to update project story element: JSON request body is invalid")
		logger.Error("Request body: ", c.Request.Body)
		logger.Error("Expected body: ", "{name: string, description: string}")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Project.StoryElement.GetByID(id); result.IsErr() {
		c.JSON(result.Error().Code.AsInt(), result.Error())
		return
	}

	element = body.ToDB()
	element.ID = uint(idUint)

	updateResult := db.Project.StoryElement.Update(element)

	if updateResult.IsErr() {
		c.JSON(updateResult.Error().Code.AsInt(), updateResult.Error())
		return
	}

	c.JSON(http.StatusOK, updateResult.Value().ToResponse())
}

func (elementType) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	result := db.Project.StoryElement.GetByID(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	result = db.Project.StoryElement.Delete(id)

	if result.IsErr() {
		err := result.Error()
		c.JSON(err.Code.AsInt(), err)
		return
	}

	c.JSON(http.StatusOK, result.Value().ToResponse())
}
