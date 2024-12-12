package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectDB struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name" gorm:"not null"`
	Description     string    `json:"description" gorm:"not null"`
	IDCreator       uint      `json:"id_creator" gorm:"foreignKey:id"`
	IDStatus        uint      `json:"id_status" gorm:"foreignKey:id"`
	IDSettings      uint      `json:"id_settings" gorm:"foreignKey:id"`
	StatusTimeStamp time.Time `json:"status_time_stamp" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"not null"`
	gorm.Model
}

func (project ProjectDB) ToResponse() ProjectResponse {
	return ProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		CreatedAt:   project.CreatedAt.String(),
		IDCreator:   project.IDCreator,
		IDStatus:    project.IDStatus,
		IDSettings:  project.IDSettings,
	}
}

type ProjectResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	IDCreator   uint   `json:"id_creator"`
	IDStatus    uint   `json:"id_status"`
	IDSettings  uint   `json:"id_settings"`
}

type ProjectCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IDCreator   uint   `json:"id_creator"`
	Public      bool   `json:"public"`
}

func (p ProjectCreate) ToDB() ProjectDB {
	return ProjectDB{
		Name:            p.Name,
		Description:     p.Description,
		IDCreator:       p.IDCreator,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		IDStatus:        1,
		StatusTimeStamp: time.Now(),
		IDSettings:      0,
	}
}
func (p ProjectCreate) Settings() ProjectSettingsCreate {
	return ProjectSettingsCreate{
		Public: p.Public,
	}
}

type ProjectUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IDStatus    uint   `json:"id_status"`
}

func (p ProjectUpdate) ToDB() ProjectDB {
	return ProjectDB{
		Name:        p.Name,
		Description: p.Description,
		IDStatus:    p.IDStatus,
		UpdatedAt:   time.Now(),
	}
}

func (ProjectDB) TableName() string {
	return "projects"
}
