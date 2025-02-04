package main

import (
	"dainxor/we/base/configs"
	"dainxor/we/base/logger"
	"dainxor/we/models"

	"github.com/joho/godotenv"
)

func init() {
	logger.Init()
	logger.Info("Loading configurations")

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")

	}

	logger.EnvInit()
	configs.DB.EnvInit()
}

func main() {
	configs.DataBase.AutoMigrate(&models.AuthCodeDB{})

	configs.DataBase.AutoMigrate(&models.UserDB{})

	configs.DataBase.AutoMigrate(&models.ProjectDB{})
	configs.DataBase.AutoMigrate(&models.ProjectSettingsDB{})
	configs.DataBase.AutoMigrate(&models.ProjectPermissionDB{})
	//configs.DataBase.AutoMigrate(&models.ProjectCollaboratorDB{})

	configs.DataBase.AutoMigrate(&models.ProjectCharacterDB{})
	configs.DataBase.AutoMigrate(&models.ProjectCharacterRelationDB{})
	configs.DataBase.AutoMigrate(&models.ProjectLocationDB{})
	configs.DataBase.AutoMigrate(&models.ProjectStoryElementDB{})
	configs.DataBase.AutoMigrate(&models.ProjectStoryElementTypeDB{})

	configs.DataBase.AutoMigrate(&models.CharacterRelationTypeDB{})

	configs.DataBase.AutoMigrate(&models.ResourceTextDB{})
}
