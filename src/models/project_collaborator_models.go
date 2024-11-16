package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectCollaboratorDB struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	IDProject uint      `json:"id_project" gorm:"foreignKey:id"`
	IDUser    uint      `json:"id_user" gorm:"foreignKey:id"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	gorm.Model
}

func (cp ProjectCollaboratorDB) ToResponse() ProjectCollaboratorResponse {
	return ProjectCollaboratorResponse{
		ID:        cp.ID,
		IDProject: cp.IDProject,
		IDUser:    cp.IDUser,
		CreatedAt: cp.CreatedAt,
	}
}

type ProjectCollaboratorResponse struct {
	ID        uint      `json:"id"`
	IDProject uint      `json:"id_project"`
	IDUser    uint      `json:"id_user"`
	CreatedAt time.Time `json:"created_at"`
}

type ProjectCollaboratorCreate struct {
	IDProject uint `json:"id_project"`
	IDUser    uint `json:"id_user"`
}

func (cp ProjectCollaboratorCreate) ToDB() ProjectCollaboratorDB {
	return ProjectCollaboratorDB{
		IDProject: cp.IDProject,
		IDUser:    cp.IDUser,
		CreatedAt: time.Now(),
	}
}

func (ProjectCollaboratorDB) TableName() string {
	return "project_collaborators"
}
