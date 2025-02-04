package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectCharacterDB struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	IDProject   uint   `json:"id_project" gorm:"foreignKey:id"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	gorm.Model
}

func (pc ProjectCharacterDB) ToResponse() ProjectCharacterResponse {
	return ProjectCharacterResponse{
		ID:          pc.ID,
		IDProject:   pc.IDProject,
		Name:        pc.Name,
		Description: pc.Description,
		CreatedAt:   pc.CreatedAt,
		UpdatedAt:   pc.UpdatedAt,
	}
}

type ProjectCharacterResponse struct {
	ID          uint      `json:"id"`
	IDProject   uint      `json:"id_project"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProjectCharacterCreate struct {
	IDProject   uint   `json:"id_project"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (pc ProjectCharacterCreate) ToDB() ProjectCharacterDB {
	return ProjectCharacterDB{
		IDProject:   pc.IDProject,
		Name:        pc.Name,
		Description: pc.Description,
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

type ProjectCharacterUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (pc ProjectCharacterUpdate) ToDB() ProjectCharacterDB {
	return ProjectCharacterDB{
		Name:        pc.Name,
		Description: pc.Description,
		Model: gorm.Model{
			UpdatedAt: time.Now(),
		},
	}
}

func (ProjectCharacterDB) TableName() string {
	return "project_characters"
}
