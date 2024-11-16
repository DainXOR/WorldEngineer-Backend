package models

import (
	"reflect"
	"time"

	"gorm.io/gorm"
)

type ResourceDB[Data any] struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	IDProject    uint      `json:"id_project" gorm:"foreignKey:id"`
	Name         string    `json:"name" gorm:"not null"`
	ResourceType int       `json:"resource_type" gorm:"not null"`
	Data         Data      `json:"data" gorm:"not null"`
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
		Data:         pr.Data,
		CreatedAt:    pr.CreatedAt,
		UpdatedAt:    pr.UpdatedAt,
	}
}

type ResourceResponse[Res any] struct {
	ID           uint      `json:"id"`
	IDProject    uint      `json:"id_project"`
	Name         string    `json:"name"`
	ResourceType int       `json:"resource_type"`
	Data         Res       `json:"data"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ResourceCreate[Res any] struct {
	IDProject    uint   `json:"id_project"`
	Name         string `json:"name"`
	ResourceType int    `json:"resource_type"`
	Data         Res    `json:"data"`
}
type ResourceTextCreate struct {
	ResourceCreate[string]
}

func (pr ResourceCreate[Res]) ToDB() ResourceDB[Res] {
	return ResourceDB[Res]{
		IDProject:    pr.IDProject,
		Name:         pr.Name,
		ResourceType: pr.ResourceType,
		Data:         pr.Data,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

type ResourceUpdate[Res any] struct {
	ID           uint   `json:"id"`
	IDProject    uint   `json:"id_project"`
	Name         string `json:"name"`
	ResourceType int    `json:"resource_type"`
	Data         Res    `json:"data"`
}
type ResourceTextUpdate struct {
	ResourceUpdate[string]
}

func (pr ResourceUpdate[Res]) ToDB() ResourceDB[Res] {
	return ResourceDB[Res]{
		ID:           pr.ID,
		IDProject:    pr.IDProject,
		Name:         pr.Name,
		ResourceType: pr.ResourceType,
		Data:         pr.Data,
		UpdatedAt:    time.Now(),
	}
}

func (ResourceDB[Res]) TableName() string {
	var r Res

	switch reflect.TypeOf(r).String() {
	case "string":
		return "project_resources_text"

	default:
		return "project_resources"
	}
}
