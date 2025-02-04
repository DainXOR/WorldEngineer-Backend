package models

import (
	"time"

	"gorm.io/gorm"
)

type nature struct{}

var RelationNature nature

func (nature) Neutral() uint {
	return 0
}
func (nature) Beneficial() uint {
	return 1
}
func (nature) Harmful() uint {
	return 2
}

type CharacterRelationTypeDB struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Name          string `json:"name" gorm:"not null"`
	IDNatureOnOne uint   `json:"id_nature_one" gorm:"not null"`
	IDNatureOnTwo uint   `json:"id_nature_two" gorm:"not null"`

	gorm.Model
}

func (crt CharacterRelationTypeDB) ToResponse() CharacterRelationTypeResponse {
	return CharacterRelationTypeResponse{
		ID:        crt.ID,
		Name:      crt.Name,
		CreatedAt: crt.CreatedAt,
		UpdatedAt: crt.UpdatedAt,
	}
}

type CharacterRelationTypeResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CharacterRelationTypeCreate struct {
	Name string `json:"name"`
}

func (crt CharacterRelationTypeCreate) ToDB() CharacterRelationTypeDB {
	return CharacterRelationTypeDB{
		Name: crt.Name,
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

type CharacterRelationTypeUpdate struct {
	Name string `json:"name"`
}

func (crt CharacterRelationTypeUpdate) ToDB() CharacterRelationTypeDB {
	return CharacterRelationTypeDB{
		Name: crt.Name,
		Model: gorm.Model{
			UpdatedAt: time.Now(),
		},
	}
}

func (CharacterRelationTypeDB) TableName() string {
	return "character_relation_types"
}
