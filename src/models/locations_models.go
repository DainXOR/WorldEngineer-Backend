package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectLocationDB struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	IDProject   uint   `json:"id_project" gorm:"foreignKey:id"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	gorm.Model
}

func (pl ProjectLocationDB) ToResponse() ProjectLocationResponse {
	return ProjectLocationResponse{
		ID:          pl.ID,
		IDProject:   pl.IDProject,
		Name:        pl.Name,
		Description: pl.Description,
		CreatedAt:   pl.CreatedAt,
		UpdatedAt:   pl.UpdatedAt,
	}
}

type ProjectLocationResponse struct {
	ID          uint      `json:"id"`
	IDProject   uint      `json:"id_project"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProjectLocationCreate struct {
	IDProject   uint   `json:"id_project"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (pl ProjectLocationCreate) ToDB() ProjectLocationDB {
	return ProjectLocationDB{
		IDProject:   pl.IDProject,
		Name:        pl.Name,
		Description: pl.Description,
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

type ProjectLocationUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (pl ProjectLocationUpdate) ToDB() ProjectLocationDB {
	return ProjectLocationDB{
		Name:        pl.Name,
		Description: pl.Description,
		Model: gorm.Model{
			UpdatedAt: time.Now(),
		},
	}
}

func (ProjectLocationDB) TableName() string {
	return "project_locations"
}
