package models

import (
	"time"

	"gorm.io/gorm"
)

type UserDB struct {
	ID              uint      `json:"id" gorm:"primarykey"`
	NameTag         string    `json:"name_tag" gorm:"not null"`
	Username        string    `json:"username" gorm:"not null"`
	Email           string    `json:"email" gorm:"not null"`
	IDStatus        uint      `json:"id_status" gorm:"foreignKey:ID"`
	StatusTimeStamp time.Time `json:"status_time_stamp" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"not null"`
	gorm.Model
}

type UserResponse struct {
	ID        uint      `json:"id"`
	NameTag   string    `json:"name_tag"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	IDStatus  uint      `json:"id_status"`
	CreatedAt time.Time `json:"created_at"`
}

func (user UserDB) ToResponse() UserResponse {
	return UserResponse{
		ID:        user.ID,
		NameTag:   user.NameTag,
		Username:  user.Username,
		Email:     user.Email,
		IDStatus:  user.IDStatus,
		CreatedAt: user.CreatedAt,
	}
}

type UserCreate struct {
	NameTag  string `json:"name_tag"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (u UserCreate) ToDB() UserDB {
	return UserDB{
		Username:        u.Username,
		Email:           u.Email,
		NameTag:         u.NameTag,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		IDStatus:        1,
		StatusTimeStamp: time.Now(),
	}
}

type UserUpdate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	IDStatus uint   `json:"id_status"`
}

func (u UserUpdate) ToDB() UserDB {
	return UserDB{
		Username:        u.Username,
		Email:           u.Email,
		IDStatus:        u.IDStatus,
		UpdatedAt:       time.Now(),
		StatusTimeStamp: time.Now(),
	}
}

func (UserDB) TableName() string {
	return "users"
}
