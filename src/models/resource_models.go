package models

import (
	"time"

	"gorm.io/gorm"
)

type ResourceDB[Res any] struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	IDProject    uint      `json:"id_project" gorm:"foreignKey:id"`
	Name         string    `json:"name" gorm:"not null"`
	ResourceType int       `json:"resource_type" gorm:"not null"`
	Resource     Res       `json:"resource" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"not null"`
	gorm.Model
}

type ResourceTextDB struct {
	ResourceDB[string]
}

func (pr ResourceDB[Res]) ToResponse() ResourceResponse[Res] {
	return ResourceResponse[Res]{
		ID:           pr.ID,
		IDProject:    pr.IDProject,
		Name:         pr.Name,
		ResourceType: pr.ResourceType,
		Resource:     pr.Resource,
		CreatedAt:    pr.CreatedAt,
		UpdatedAt:    pr.UpdatedAt,
	}
}

type ResourceResponse[Res any] struct {
	ID           uint      `json:"id"`
	IDProject    uint      `json:"id_project"`
	Name         string    `json:"name"`
	ResourceType int       `json:"resource_type"`
	Resource     Res       `json:"resource"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ResourceCreate[Res any] struct {
	IDProject    uint   `json:"id_project"`
	Name         string `json:"name"`
	ResourceType int    `json:"resource_type"`
	Resource     Res    `json:"resource"`
}
type ResourceTextCreate struct {
	ResourceCreate[string]
}

func (pr ResourceCreate[Res]) ToDB() ResourceDB[Res] {
	return ResourceDB[Res]{
		IDProject:    pr.IDProject,
		Name:         pr.Name,
		ResourceType: pr.ResourceType,
		Resource:     pr.Resource,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (ResourceDB[Res]) TableName() string {
	return "project_resources_text"
}
