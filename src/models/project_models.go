package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectDB struct {
	gorm.Model
	ID              uint      `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name" gorm:"not null"`
	Description     string    `json:"description" gorm:"not null"`
	IDCreator       uint      `json:"id_creator" gorm:"foreignKey:ID"`
	CreatedAt       time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"not null"`
	IDStatus        uint      `json:"id_status" gorm:"foreignKey:ID"`
	StatusTimeStamp time.Time `json:"status_time_stamp" gorm:"not null"`
	IDSettings      uint      `json:"id_settings" gorm:"foreignKey:ID"`
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

type ProjectUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IDStatus    int    `json:"id_status"`
}

func (p *ProjectDB) TableName() string {
	return "projects"
}
