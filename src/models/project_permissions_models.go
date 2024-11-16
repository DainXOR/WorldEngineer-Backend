package models

type ProjectPermissionDB struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}

func (p ProjectPermissionDB) ToResponse() ProjectPermissionResponse {
	return ProjectPermissionResponse(p)
}

type ProjectPermissionResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (ProjectPermissionDB) TableName() string {
	return "project_permissions"
}
