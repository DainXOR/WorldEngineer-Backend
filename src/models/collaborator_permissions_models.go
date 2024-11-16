package models

import (
	"time"

	"gorm.io/gorm"
)

type CollaboratorPermissionDB struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	IDCollaborator uint      `json:"id_collaborator" gorm:"foreignKey:id"`
	IDPermission   uint      `json:"id_permission" gorm:"foreignKey:id"`
	CreatedAt      time.Time `json:"created_at" gorm:"not null"`
	gorm.Model
}

func (cp CollaboratorPermissionDB) ToResponse() CollaboratorPermissionResponse {
	return CollaboratorPermissionResponse{
		ID:             cp.ID,
		IDCollaborator: cp.IDCollaborator,
		IDPermission:   cp.IDPermission,
		CreatedAt:      cp.CreatedAt,
	}
}

type CollaboratorPermissionResponse struct {
	ID             uint      `json:"id"`
	IDCollaborator uint      `json:"id_collaborator"`
	IDPermission   uint      `json:"id_permission"`
	CreatedAt      time.Time `json:"created_at"`
}

type CollaboratorPermissionCreate struct {
	IDCollaborator uint `json:"id_collaborator"`
	IDPermission   uint `json:"id_permission"`
}

func (cp CollaboratorPermissionCreate) ToDB() CollaboratorPermissionDB {
	return CollaboratorPermissionDB{
		IDCollaborator: cp.IDCollaborator,
		IDPermission:   cp.IDPermission,
		CreatedAt:      time.Now(),
	}
}

func (CollaboratorPermissionDB) TableName() string {
	return "collaborator_permissions"
}
