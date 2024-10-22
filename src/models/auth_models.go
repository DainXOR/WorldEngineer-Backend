package models

import (
	"time"

	"gorm.io/gorm"
)

type AuthCodeDB struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey"`
	Code      uint      `json:"code" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}