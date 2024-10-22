package models

import (
	"time"

	"gorm.io/gorm"
)

type UserDB struct {
	gorm.Model
	ID              uint      `json:"id" gorm:"primaryKey"`
	NameTag         string    `json:"name_tag" gorm:"not null"`
	Username        string    `json:"username" gorm:"not null"`
	Email           string    `json:"email" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"not null"`
	IDStatus        uint      `json:"id_status" gorm:"foreignKey:ID"`
	StatusTimeStamp time.Time `json:"status_time_stamp" gorm:"not null"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	NameTag  string `json:"name_tag"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IDStatus uint   `json:"id_status"`
}

func (user UserDB) ToResponse() UserResponse {
	return UserResponse{
		ID:       user.ID,
		NameTag:  user.NameTag,
		Username: user.Username,
		Email:    user.Email,
		IDStatus: user.IDStatus,
	}
}

type UserCreate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserUpdate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	IDStatus uint   `json:"id_status"`
}

func (u *UserDB) TableName() string {
	return "users"
}
