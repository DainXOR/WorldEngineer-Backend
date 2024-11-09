package db

import (
	"dainxor/we/base/configs"
	"dainxor/we/models"
	"dainxor/we/types"
	"fmt"
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

func (projectType) Create(project models.ProjectDB) types.Result[models.ProjectDB, models.ErrorResponse] {
	configs.DataBase.Create(&project)
	return types.ResultOk[models.ProjectDB, models.ErrorResponse](project)
}

func (projectType) GetByID(id string) types.Result[models.ProjectDB, models.ErrorResponse] {
	var project models.ProjectDB
	configs.DataBase.First(&project, id)

	err := models.ErrorNotFound(
		"Project not found",
		"Project with ID "+id+" not found",
	)
	return types.ResultOf(project, err, project.ID != 0)
}

func (projectType) GetAll() types.Result[[]models.ProjectDB, models.ErrorResponse] {
	var projects []models.ProjectDB
	configs.DataBase.Find(&projects)

	err := models.ErrorNotFound(
		"Projects not found",
	)
	return types.ResultOf(projects, err, len(projects) > 0)
}
func (projectType) GetByUserID(id string) types.Result[[]models.ProjectDB, models.ErrorResponse] {
	var projects []models.ProjectDB
	configs.DataBase.Where("id_creator = ?", id).Find(&projects)

	err := models.ErrorNotFound(
		"Projects not found",
		"Projects with user ID "+id+" not found",
	)
	return types.ResultOf(projects, err, len(projects) > 0)
}

func (projectType) Update(project models.ProjectDB) types.Result[models.ProjectDB, models.ErrorResponse] {
	configs.DataBase.Save(&project)
	return types.ResultOk[models.ProjectDB, models.ErrorResponse](project)
}

func (projectType) Delete(id string) types.Result[models.ProjectDB, models.ErrorResponse] {
	var project models.ProjectDB
	configs.DataBase.First(&project, id)
	configs.DataBase.Delete(&project)
	return types.ResultOk[models.ProjectDB, models.ErrorResponse](project)
}

// > Settings

func (settingsType) Create(settings models.ProjectSettingsDB) types.Result[models.ProjectSettingsDB, models.ErrorResponse] {
	configs.DataBase.Create(&settings)
	return types.ResultOf(
		settings,
		models.ErrorInternal("Failed to create project settings"),
		settings.ID != 0,
	)
}

func (settingsType) GetByID(id string) types.Result[models.ProjectSettingsDB, models.ErrorResponse] {
	var settings models.ProjectSettingsDB
	configs.DataBase.First(&settings, id)

	err := models.ErrorNotFound(
		"Project settings not found",
		"Project settings with ID "+id+" not found",
	)
	return types.ResultOf(settings, err, settings.ID != 0)
}

func (settingsType) GetAll() types.Result[[]models.ProjectSettingsDB, models.ErrorResponse] {
	var settings []models.ProjectSettingsDB
	configs.DataBase.Find(&settings)

	err := models.ErrorNotFound(
		"Project settings not found",
	)
	return types.ResultOf(settings, err, len(settings) > 0)
}

func (settingsType) Update(settings models.ProjectSettingsDB) types.Result[models.ProjectSettingsDB, models.ErrorResponse] {
	configs.DataBase.Save(&settings)
	return types.ResultOk[models.ProjectSettingsDB, models.ErrorResponse](settings)
}

func (settingsType) Delete(id string) types.Result[models.ProjectSettingsDB, models.ErrorResponse] {
	var settings models.ProjectSettingsDB
	configs.DataBase.First(&settings, id)
	configs.DataBase.Delete(&settings)
	return types.ResultOk[models.ProjectSettingsDB, models.ErrorResponse](settings)
}

// > Collaborator

func (collaboratorType) Create(collaborator models.ProjectCollaboratorDB) types.Result[models.ProjectCollaboratorDB, models.ErrorResponse] {
	configs.DataBase.Create(&collaborator)
	return types.ResultOk[models.ProjectCollaboratorDB, models.ErrorResponse](collaborator)
}

func (collaboratorType) GetByID(id string) types.Result[models.ProjectCollaboratorDB, models.ErrorResponse] {
	var collaborator models.ProjectCollaboratorDB
	configs.DataBase.First(&collaborator, id)

	err := models.ErrorNotFound(
		"Project collaborator not found",
		"Project collaborator with ID "+id+" not found",
	)
	return types.ResultOf(collaborator, err, collaborator.ID != 0)
}
func (collaboratorType) GetByUserIDAndProjectID(userID string, projectID string) types.Result[models.ProjectCollaboratorDB, models.ErrorResponse] {
	var collaborator models.ProjectCollaboratorDB
	configs.DataBase.Where("id_user = ? AND id_project = ?", userID, projectID).First(&collaborator)

	err := models.ErrorNotFound(
		"Project collaborator not found",
		"Project collaborator with user ID "+userID+" and project ID "+projectID+" not found",
	)
	return types.ResultOf(collaborator, err, collaborator.ID != 0)
}

func (collaboratorType) GetAll() types.Result[[]models.ProjectCollaboratorDB, models.ErrorResponse] {
	var collaborators []models.ProjectCollaboratorDB
	configs.DataBase.Find(&collaborators)

	err := models.ErrorNotFound(
		"Project collaborators not found",
	)
	return types.ResultOf(collaborators, err, len(collaborators) > 0)
}
func (collaboratorType) GetByProjectID(id string) types.Result[[]models.ProjectCollaboratorDB, models.ErrorResponse] {
	var collaborators []models.ProjectCollaboratorDB
	configs.DataBase.Where("id_project = ?", id).Find(&collaborators)

	err := models.ErrorNotFound(
		"Project collaborators not found",
		"Project collaborators with project ID "+id+" not found",
	)
	return types.ResultOf(collaborators, err, len(collaborators) > 0)
}
func (collaboratorType) GetByUserID(id string) types.Result[[]models.ProjectCollaboratorDB, models.ErrorResponse] {
	var collaborators []models.ProjectCollaboratorDB
	configs.DataBase.Where("id_user = ?", id).Find(&collaborators)

	err := models.ErrorNotFound(
		"Project collaborators not found",
		"Project collaborators with user ID "+id+" not found",
	)
	return types.ResultOf(collaborators, err, len(collaborators) > 0)
}

func (collaboratorType) Delete(id string) types.Result[models.ProjectCollaboratorDB, models.ErrorResponse] {
	var collaborator models.ProjectCollaboratorDB
	configs.DataBase.First(&collaborator, id)
	configs.DataBase.Delete(&collaborator)
	return types.ResultOk[models.ProjectCollaboratorDB, models.ErrorResponse](collaborator)
}

// > Permission

func (permissionType) Create(permission models.CollaboratorPermissionDB) types.Result[models.CollaboratorPermissionDB, models.ErrorResponse] {
	configs.DataBase.Create(&permission)
	return types.ResultOk[models.CollaboratorPermissionDB, models.ErrorResponse](permission)
}

func (permissionType) GetByID(id string) types.Result[models.CollaboratorPermissionDB, models.ErrorResponse] {
	var permission models.CollaboratorPermissionDB
	configs.DataBase.First(&permission, id)

	err := models.ErrorNotFound(
		"Collaborator permission not found",
		"Collaborator permission with ID "+id+" not found",
	)
	return types.ResultOf(permission, err, permission.ID != 0)
}

func (permissionType) GetByCollaboratorID(id string) types.Result[[]models.CollaboratorPermissionDB, models.ErrorResponse] {
	var permission []models.CollaboratorPermissionDB
	configs.DataBase.Where("id_collaborator = ?", id).Find(&permission)

	err := models.ErrorNotFound(
		"Collaborator permissions not found",
		"Collaborator permissions with collaborator ID "+id+" not found",
	)
	return types.ResultOf(permission, err, len(permission) > 0)
}
func (permissionType) GetByProjectIDAndPermissionID(idPermission string, idProject string) types.Result[map[uint][]models.CollaboratorPermissionDB, models.ErrorResponse] {
	collaborators := Project.Collaborator.GetByProjectID(idProject)

	if collaborators.IsErr() {
		return types.ResultErr[map[uint][]models.CollaboratorPermissionDB](collaborators.Error())
	}

	var permission = make(map[uint][]models.CollaboratorPermissionDB, len(collaborators.Value()))

	for _, collaborator := range collaborators.Value() {
		collaboratorPermissions := Project.Permission.GetByCollaboratorID(fmt.Sprint(collaborator.ID))
		permissions, err := collaboratorPermissions.Get()

		if err.IsPresent() {
			permissions = types.OptionalOf(make([]models.CollaboratorPermissionDB, 0))
		}

		permission[uint(collaborator.ID)] = permissions.Get()
	}

	return types.ResultOk[map[uint][]models.CollaboratorPermissionDB, models.ErrorResponse](permission)
}

func (permissionType) Delete(id string) types.Result[models.CollaboratorPermissionDB, models.ErrorResponse] {
	var permission models.CollaboratorPermissionDB
	configs.DataBase.First(&permission, id)
	configs.DataBase.Delete(&permission)
	return types.ResultOk[models.CollaboratorPermissionDB, models.ErrorResponse](permission)
}
