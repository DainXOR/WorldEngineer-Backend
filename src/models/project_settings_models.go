package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectSettingsDB struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey"`
	Public    bool      `json:"public" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

type ProjectSettingsResponse struct {
	ID     uint `json:"id"`
	Public bool `json:"public"`
}

type ProjectSettingsCreate struct {
	Public bool `json:"public"`
}

func (p *ProjectSettingsDB) TableName() string {
	return "project_settings"
}
