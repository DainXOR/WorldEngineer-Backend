package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectStoryElementDB struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	IDProject   uint   `json:"id_project" gorm:"foreignKey:id"`
	IDType      uint   `json:"id_type" gorm:"foreignKey:id"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	gorm.Model
}

func (pse ProjectStoryElementDB) ToResponse() ProjectStoryElementResponse {
	return ProjectStoryElementResponse{
		ID:          pse.ID,
		IDProject:   pse.IDProject,
		IDType:      pse.IDType,
		Name:        pse.Name,
		Description: pse.Description,
		CreatedAt:   pse.CreatedAt,
		UpdatedAt:   pse.UpdatedAt,
	}
}

type ProjectStoryElementResponse struct {
	ID          uint      `json:"id"`
	IDProject   uint      `json:"id_project"`
	IDType      uint      `json:"id_type"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProjectStoryElementCreate struct {
	IDProject   uint   `json:"id_project"`
	IDType      uint   `json:"id_type"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (psec ProjectStoryElementCreate) ToDB() ProjectStoryElementDB {
	return ProjectStoryElementDB{
		IDProject:   psec.IDProject,
		IDType:      psec.IDType,
		Name:        psec.Name,
		Description: psec.Description,
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

type ProjectStoryElementUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (pseu ProjectStoryElementUpdate) ToDB() ProjectStoryElementDB {
	return ProjectStoryElementDB{
		Name:        pseu.Name,
		Description: pseu.Description,
		Model: gorm.Model{
			UpdatedAt: time.Now(),
		},
	}
}

func (ProjectStoryElementDB) TableName() string {
	return "project_story_elements"
}
