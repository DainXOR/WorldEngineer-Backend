package models

import (
	"time"
	//"gorm.io/gorm"
)

type AuthCodeDB struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Code       string    `json:"code" gorm:"not null"`
	Email      string    `json:"email" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null"`
	ConsumedAt time.Time `json:"consumed_at"`
}

type AuthCodeUpdate struct {
	Code       string    `json:"code"`
	Email      string    `json:"email"`
	ConsumedAt time.Time `json:"consumed_at"`
}

func (u *AuthCodeDB) TableName() string {
	return "auth_codes"
}
