package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectSettingsDB struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Public    bool      `json:"public" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
	gorm.Model
}

func (p ProjectSettingsDB) ToResponse() ProjectSettingsResponse {
	return ProjectSettingsResponse{
		ID:     p.ID,
		Public: p.Public,
	}
}

type ProjectSettingsResponse struct {
	ID     uint `json:"id"`
	Public bool `json:"public"`
}

type ProjectSettingsCreate struct {
	Public bool `json:"public"`
}

func (p ProjectSettingsCreate) ToDB() ProjectSettingsDB {
	return ProjectSettingsDB{
		Public:    p.Public,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type ProjectSettingsUpdate struct {
	Public bool `json:"public"`
}

func (p ProjectSettingsUpdate) ToDB() ProjectSettingsDB {
	return ProjectSettingsDB{
		Public:    p.Public,
		UpdatedAt: time.Now(),
	}
}

func (ProjectSettingsDB) TableName() string {
	return "project_settings"
}
