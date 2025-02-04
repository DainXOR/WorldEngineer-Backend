package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectStoryElementTypeDB struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
	gorm.Model
}

func (pset ProjectStoryElementTypeDB) ToResponse() ProjectStoryElementTypeResponse {
	return ProjectStoryElementTypeResponse{
		ID:        pset.ID,
		Name:      pset.Name,
		CreatedAt: pset.CreatedAt,
		UpdatedAt: pset.UpdatedAt,
	}
}

type ProjectStoryElementTypeResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProjectStoryElementTypeCreate struct {
	Name string `json:"name"`
}

func (psetc ProjectStoryElementTypeCreate) ToDB() ProjectStoryElementTypeDB {
	return ProjectStoryElementTypeDB{
		Name: psetc.Name,
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

type ProjectStoryElementTypeUpdate struct {
	Name string `json:"name"`
}

func (psetu ProjectStoryElementTypeUpdate) ToDB() ProjectStoryElementTypeDB {
	return ProjectStoryElementTypeDB{
		Name: psetu.Name,
	}
}

func (ProjectStoryElementTypeDB) TableName() string {
	return "project_story_element_types"
}
