package db

import (
	"dainxor/we/base/configs"
	"dainxor/we/models"
	"dainxor/we/types"
	"fmt"
	"strconv"
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
func (collaboratorType) GetByProjectIDAndPermissionID(idPermission string, idProject string) types.Result[[]models.ProjectCollaboratorDB, models.ErrorResponse] {
	int64Permission, err := strconv.ParseUint(idPermission, 10, 64)

	if err == nil {
		return types.ResultErr[[]models.ProjectCollaboratorDB](
			models.Error(
				types.Http.BadRequest(),
				"bad_request",
				"Invalid permission ID",
			),
		)
	}

	permissionID := uint(int64Permission)
	collaborators := Project.Collaborator.GetByProjectID(idProject)

	if collaborators.IsErr() {
		return types.ResultErr[[]models.ProjectCollaboratorDB](collaborators.Error())
	}

	var matchingCollaborators = make([]models.ProjectCollaboratorDB, 0)

	for _, collaborator := range collaborators.Value() {
		collaboratorPermissions := Project.Permission.GetByCollaboratorID(fmt.Sprint(collaborator.ID))

		if collaboratorPermissions.IsErr() {
			continue
		}

		for _, permission := range collaboratorPermissions.Value() {
			if permission.IDPermission == permissionID {
				matchingCollaborators = append(matchingCollaborators, collaborator)
				break
			}
		}
	}

	return types.ResultOk[[]models.ProjectCollaboratorDB, models.ErrorResponse](matchingCollaborators)
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

func (permissionType) Delete(id string) types.Result[models.CollaboratorPermissionDB, models.ErrorResponse] {
	var permission models.CollaboratorPermissionDB
	configs.DataBase.First(&permission, id)
	configs.DataBase.Delete(&permission)
	return types.ResultOk[models.CollaboratorPermissionDB, models.ErrorResponse](permission)
}

// > Resources

// Currently, the resources can only be strings, but this will be changed in the future
func (resourcesType) CreateText(resource models.ResourceTextDB) types.Result[models.ResourceDB[string], models.ErrorResponse] {

	configs.DataBase.Create(&resource)

	err := models.ErrorInternal("Failed to create project resource")

	return types.ResultOf(resource.ResourceDB, err, resource.ID != 0)
}

func (resourcesType) GetTextByID(id string) types.Result[models.ResourceDB[string], models.ErrorResponse] {
	var textResource models.ResourceTextDB
	configs.DataBase.First(&textResource, id)

	err := models.ErrorNotFound(
		"Project resource not found",
		"Project resource with ID "+id+" not found",
	)

	return types.ResultOf(textResource.ResourceDB, err, textResource.ID != 0)
}
func (resourcesType) GetTextByProjectID(id string) types.Result[[]models.ResourceTextDB, models.ErrorResponse] {
	var resources []models.ResourceDB[string]
	configs.DataBase.Where("id_project = ?", id).Find(&resources)

	err := models.ErrorNotFound(
		"Project resources not found",
		"Project resources with project ID "+id+" not found",
	)

	textResources := make([]models.ResourceTextDB, len(resources))
	for i, resource := range resources {
		textResources[i] = models.ResourceTextDB{
			ResourceDB: resource,
		}
	}

	return types.ResultOf(textResources, err, len(resources) > 0)
}

func (resourcesType) UpdateText(resource models.ResourceTextDB) types.Result[models.ResourceTextDB, models.ErrorResponse] {
	configs.DataBase.Save(&resource.ResourceDB)
	return types.ResultOk[models.ResourceTextDB, models.ErrorResponse](resource)
}

func (resourcesType) DeleteTextByID(id string) types.Result[models.ResourceTextDB, models.ErrorResponse] {
	var resource models.ResourceTextDB
	configs.DataBase.First(&resource.ResourceDB, id)
	configs.DataBase.Delete(&resource.ResourceDB)
	return types.ResultOk[models.ResourceTextDB, models.ErrorResponse](resource)
}

// > Character

func (characterType) Create(character models.ProjectCharacterDB) types.Result[models.ProjectCharacterDB, models.ErrorResponse] {
	configs.DataBase.Create(&character)
	return types.ResultOk[models.ProjectCharacterDB, models.ErrorResponse](character)
}

func (characterType) GetByID(id string) types.Result[models.ProjectCharacterDB, models.ErrorResponse] {
	var character models.ProjectCharacterDB
	configs.DataBase.First(&character, id)

	err := models.ErrorNotFound(
		"Project character not found",
		"Project character with ID "+id+" not found",
	)
	return types.ResultOf(character, err, character.ID != 0)
}
func (characterType) GetByName(name string) types.Result[models.ProjectCharacterDB, models.ErrorResponse] {
	var character models.ProjectCharacterDB
	configs.DataBase.Where("name = ?", name).First(&character)

	err := models.ErrorNotFound(
		"Project character not found",
		"Project character with name "+name+" not found",
	)
	return types.ResultOf(character, err, character.ID != 0)
}

func (characterType) GetByProjectID(id string) types.Result[[]models.ProjectCharacterDB, models.ErrorResponse] {
	var characters []models.ProjectCharacterDB
	configs.DataBase.Where("id_project = ?", id).Find(&characters)

	err := models.ErrorNotFound(
		"Project characters not found",
		"Project characters with project ID "+id+" not found",
	)
	return types.ResultOf(characters, err, len(characters) > 0)
}

func (characterType) Update(character models.ProjectCharacterDB) types.Result[models.ProjectCharacterDB, models.ErrorResponse] {
	configs.DataBase.Save(&character)
	return types.ResultOk[models.ProjectCharacterDB, models.ErrorResponse](character)
}

func (characterType) Delete(id string) types.Result[models.ProjectCharacterDB, models.ErrorResponse] {
	var character models.ProjectCharacterDB
	configs.DataBase.First(&character, id)
	configs.DataBase.Delete(&character)
	return types.ResultOk[models.ProjectCharacterDB, models.ErrorResponse](character)
}

// > Character Relation

func (characterType) CreateRelation(relation models.ProjectCharacterRelationDB) types.Result[models.ProjectCharacterRelationDB, models.ErrorResponse] {
	configs.DataBase.Create(&relation)
	return types.ResultOk[models.ProjectCharacterRelationDB, models.ErrorResponse](relation)
}

func (characterType) GetRelationByID(id string) types.Result[models.ProjectCharacterRelationDB, models.ErrorResponse] {
	var relation models.ProjectCharacterRelationDB
	configs.DataBase.First(&relation, id)

	err := models.ErrorNotFound(
		"Project character relation not found",
		"Project character relation with ID "+id+" not found",
	)
	return types.ResultOf(relation, err, relation.ID != 0)
}
func (characterType) GetRelationByCharacterIDs(idCharacterOne string, idCharacterTwo string) types.Result[models.ProjectCharacterRelationDB, models.ErrorResponse] {
	var relation models.ProjectCharacterRelationDB
	configs.DataBase.Where("id_character_one = ? AND id_character_two = ?", idCharacterOne, idCharacterTwo).First(&relation)

	err := models.ErrorNotFound(
		"Project character relation not found",
		"Project character relation with character IDs "+idCharacterOne+" and "+idCharacterTwo+" not found",
	)
	return types.ResultOf(relation, err, relation.ID != 0)
}

func (characterType) GetRelationsByProjectID(id string) types.Result[[]models.ProjectCharacterRelationDB, models.ErrorResponse] {
	var relations []models.ProjectCharacterRelationDB
	configs.DataBase.Where("id_project = ?", id).Find(&relations)

	err := models.ErrorNotFound(
		"No relations found for this project",
		"Project character relations with project ID "+id+" not found",
	)
	return types.ResultOf(relations, err, len(relations) > 0)
}
func (characterType) GetRelationsByCharacterOneID(id uint) types.Result[[]models.ProjectCharacterRelationDB, models.ErrorResponse] {
	var relations []models.ProjectCharacterRelationDB
	configs.DataBase.Where("id_character_one = ?", id).Find(&relations)

	err := models.ErrorNotFound(
		"No relations found for this character",
		"Project character relations with character ID "+fmt.Sprint(id)+" not found",
	)
	return types.ResultOf(relations, err, len(relations) > 0)
}
func (characterType) GetRelationsByCharacterTwoID(id uint) types.Result[[]models.ProjectCharacterRelationDB, models.ErrorResponse] {
	var relations []models.ProjectCharacterRelationDB
	configs.DataBase.Where("id_character_two = ?", id).Find(&relations)

	err := models.ErrorNotFound(
		"No relations found for this character",
		"Project character relations with character ID "+fmt.Sprint(id)+" not found",
	)
	return types.ResultOf(relations, err, len(relations) > 0)
}
func (characterType) GetRelationsByCharacterID(id string) types.Result[[]models.ProjectCharacterRelationDB, models.ErrorResponse] {
	var relations []models.ProjectCharacterRelationDB
	configs.DataBase.Where("id_character_one = ? OR id_character_two = ?", id, id).Find(&relations)

	err := models.ErrorNotFound(
		"No relations found for this character",
		"Project character relations with character ID "+id+" not found",
	)
	return types.ResultOf(relations, err, len(relations) > 0)
}
func (characterType) GetRelationsByType(id uint) types.Result[[]models.ProjectCharacterRelationDB, models.ErrorResponse] {
	var relations []models.ProjectCharacterRelationDB
	configs.DataBase.Where("id_type = ?", id).Find(&relations)

	err := models.ErrorNotFound(
		"No relations found for this type",
		"Project character relations with type ID "+fmt.Sprint(id)+" not found",
	)
	return types.ResultOf(relations, err, len(relations) > 0)
}

func (characterType) UpdateRelation(relation models.ProjectCharacterRelationDB) types.Result[models.ProjectCharacterRelationDB, models.ErrorResponse] {
	configs.DataBase.Save(&relation)
	return types.ResultOk[models.ProjectCharacterRelationDB, models.ErrorResponse](relation)
}

func (characterType) DeleteRelation(id string) types.Result[models.ProjectCharacterRelationDB, models.ErrorResponse] {
	var relation models.ProjectCharacterRelationDB
	configs.DataBase.First(&relation, id)
	configs.DataBase.Delete(&relation)
	return types.ResultOk[models.ProjectCharacterRelationDB, models.ErrorResponse](relation)
}

// > Character Relation Type

func (characterType) CreateRelationType(relationType models.CharacterRelationTypeDB) types.Result[models.CharacterRelationTypeDB, models.ErrorResponse] {
	configs.DataBase.Create(&relationType)
	return types.ResultOk[models.CharacterRelationTypeDB, models.ErrorResponse](relationType)
}

func (characterType) GetRelationTypeByID(id string) types.Result[models.CharacterRelationTypeDB, models.ErrorResponse] {
	var relationType models.CharacterRelationTypeDB
	configs.DataBase.First(&relationType, id)

	err := models.ErrorNotFound(
		"Character relation type not found",
		"Character relation type with ID "+id+" not found",
	)
	return types.ResultOf(relationType, err, relationType.ID != 0)
}

func (characterType) GetAllRelationTypes() types.Result[[]models.CharacterRelationTypeDB, models.ErrorResponse] {
	var relationTypes []models.CharacterRelationTypeDB
	configs.DataBase.Find(&relationTypes)

	err := models.ErrorNotFound(
		"Character relation types not found",
	)
	return types.ResultOf(relationTypes, err, len(relationTypes) > 0)
}

func (characterType) UpdateRelationType(relationType models.CharacterRelationTypeDB) types.Result[models.CharacterRelationTypeDB, models.ErrorResponse] {
	configs.DataBase.Save(&relationType)
	return types.ResultOk[models.CharacterRelationTypeDB, models.ErrorResponse](relationType)
}

func (characterType) DeleteRelationType(id string) types.Result[models.CharacterRelationTypeDB, models.ErrorResponse] {
	var relationType models.CharacterRelationTypeDB
	configs.DataBase.First(&relationType, id)
	configs.DataBase.Delete(&relationType)
	return types.ResultOk[models.CharacterRelationTypeDB, models.ErrorResponse](relationType)
}

// > Location

func (locationType) Create(location models.ProjectLocationDB) types.Result[models.ProjectLocationDB, models.ErrorResponse] {
	configs.DataBase.Create(&location)
	return types.ResultOk[models.ProjectLocationDB, models.ErrorResponse](location)
}

func (locationType) GetByID(id string) types.Result[models.ProjectLocationDB, models.ErrorResponse] {
	var location models.ProjectLocationDB
	configs.DataBase.First(&location, id)

	err := models.ErrorNotFound(
		"Project location not found",
		"Project location with ID "+id+" not found",
	)
	return types.ResultOf(location, err, location.ID != 0)
}
func (locationType) GetByName(name string) types.Result[models.ProjectLocationDB, models.ErrorResponse] {
	var location models.ProjectLocationDB
	configs.DataBase.Where("name = ?", name).First(&location)

	err := models.ErrorNotFound(
		"Project location not found",
		"Project location with name "+name+" not found",
	)
	return types.ResultOf(location, err, location.ID != 0)
}

func (locationType) GetByProjectID(id string) types.Result[[]models.ProjectLocationDB, models.ErrorResponse] {
	var locations []models.ProjectLocationDB
	configs.DataBase.Where("id_project = ?", id).Find(&locations)

	err := models.ErrorNotFound(
		"Project locations not found",
		"Project locations with project ID "+id+" not found",
	)
	return types.ResultOf(locations, err, len(locations) > 0)
}

func (locationType) Update(location models.ProjectLocationDB) types.Result[models.ProjectLocationDB, models.ErrorResponse] {
	configs.DataBase.Save(&location)
	return types.ResultOk[models.ProjectLocationDB, models.ErrorResponse](location)
}

func (locationType) Delete(id string) types.Result[models.ProjectLocationDB, models.ErrorResponse] {
	var location models.ProjectLocationDB
	configs.DataBase.First(&location, id)
	configs.DataBase.Delete(&location)
	return types.ResultOk[models.ProjectLocationDB, models.ErrorResponse](location)
}

// > Story Element

func (elementType) Create(element models.ProjectStoryElementDB) types.Result[models.ProjectStoryElementDB, models.ErrorResponse] {
	configs.DataBase.Create(&element)
	return types.ResultOk[models.ProjectStoryElementDB, models.ErrorResponse](element)
}

func (elementType) GetByID(id string) types.Result[models.ProjectStoryElementDB, models.ErrorResponse] {
	var element models.ProjectStoryElementDB
	configs.DataBase.First(&element, id)

	err := models.ErrorNotFound(
		"Project story element not found",
		"Project story element with ID "+id+" not found",
	)
	return types.ResultOf(element, err, element.ID != 0)
}
func (elementType) GetByName(name string) types.Result[models.ProjectStoryElementDB, models.ErrorResponse] {
	var element models.ProjectStoryElementDB
	configs.DataBase.Where("name = ?", name).First(&element)

	err := models.ErrorNotFound(
		"Project story element not found",
		"Project story element with name "+name+" not found",
	)
	return types.ResultOf(element, err, element.ID != 0)
}

func (elementType) GetByProjectID(id string) types.Result[[]models.ProjectStoryElementDB, models.ErrorResponse] {
	var elements []models.ProjectStoryElementDB
	configs.DataBase.Where("id_project = ?", id).Find(&elements)

	err := models.ErrorNotFound(
		"Project story elements not found",
		"Project story elements with project ID "+id+" not found",
	)
	return types.ResultOf(elements, err, len(elements) > 0)
}

func (elementType) Update(element models.ProjectStoryElementDB) types.Result[models.ProjectStoryElementDB, models.ErrorResponse] {
	configs.DataBase.Save(&element)
	return types.ResultOk[models.ProjectStoryElementDB, models.ErrorResponse](element)
}

func (elementType) Delete(id string) types.Result[models.ProjectStoryElementDB, models.ErrorResponse] {
	var element models.ProjectStoryElementDB
	configs.DataBase.First(&element, id)
	configs.DataBase.Delete(&element)
	return types.ResultOk[models.ProjectStoryElementDB, models.ErrorResponse](element)
}
