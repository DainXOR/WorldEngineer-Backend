package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectCharacterRelationDB struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	IDProject      uint   `json:"id_project" gorm:"foreignKey:id"`
	IDCharacterOne uint   `json:"id_character_one" gorm:"foreignKey:id"`
	IDCharacterTwo uint   `json:"id_character_two" gorm:"foreignKey:id"`
	IDType         uint   `json:"relation_type" gorm:"foreignKey:id"`
	Name           string `json:"name" gorm:"not null"`
	gorm.Model
}

func (pcr ProjectCharacterRelationDB) ToResponse() ProjectCharacterRelationResponse {
	return ProjectCharacterRelationResponse{
		ID:             pcr.ID,
		IDProject:      pcr.IDProject,
		IDCharacterOne: pcr.IDCharacterOne,
		IDCharacterTwo: pcr.IDCharacterTwo,
		IDType:         pcr.IDType,
		Name:           pcr.Name,
		CreatedAt:      pcr.CreatedAt,
		UpdatedAt:      pcr.UpdatedAt,
	}
}

type ProjectCharacterRelationResponse struct {
	ID             uint      `json:"id"`
	IDProject      uint      `json:"id_project"`
	IDCharacterOne uint      `json:"id_character_one"`
	IDCharacterTwo uint      `json:"id_character_two"`
	IDType         uint      `json:"relation_type"`
	Name           string    `json:"name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ProjectCharacterRelationCreate struct {
	IDProject      uint   `json:"id_project"`
	IDCharacterOne uint   `json:"id_character_one"`
	IDCharacterTwo uint   `json:"id_character_two"`
	IDType         uint   `json:"relation_type"`
	Name           string `json:"name"`
}

func (pcr ProjectCharacterRelationCreate) ToDB() ProjectCharacterRelationDB {
	return ProjectCharacterRelationDB{
		IDProject:      pcr.IDProject,
		IDCharacterOne: pcr.IDCharacterOne,
		IDCharacterTwo: pcr.IDCharacterTwo,
		IDType:         pcr.IDType,
		Name:           pcr.Name,
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

type ProjectCharacterRelationUpdate struct {
	IDType uint   `json:"relation_type"`
	Name   string `json:"name"`
}

func (pcr ProjectCharacterRelationUpdate) ToDB() ProjectCharacterRelationDB {
	return ProjectCharacterRelationDB{
		IDType: pcr.IDType,
		Name:   pcr.Name,
		Model: gorm.Model{
			UpdatedAt: time.Now(),
		},
	}
}

func (ProjectCharacterRelationDB) TableName() string {
	return "project_character_relations"
}
