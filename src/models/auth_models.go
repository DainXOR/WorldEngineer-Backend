package models

import (
	"database/sql"
	"time"
)

type AuthCodeDB struct {
	ID         uint         `json:"id" gorm:"primaryKey"`
	Code       string       `json:"code" gorm:"not null"`
	Email      string       `json:"email" gorm:"not null"`
	CreatedAt  time.Time    `json:"created_at" gorm:"not null"`
	ConsumedAt sql.NullTime `json:"consumed_at"`
}

func AuthCode(code string, email string) AuthCodeDB {
	return AuthCodeDB{
		Code:      code,
		Email:     email,
		CreatedAt: time.Now(),
	}
}

type AuthCodeUpdate struct {
	ConsumedAt time.Time `json:"consumed_at"`
}

func (u *AuthCodeDB) TableName() string {
	return "auth_codes"
}
