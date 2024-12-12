package models

type ProjectStatusDB struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}

func (p ProjectStatusDB) ToResponse() ProjectStatusResponse {
	return ProjectStatusResponse(p)

	//return ProjectStatusResponse{
	//	ID:   p.ID,
	//	Name: p.Name,
	//}
}

type ProjectStatusResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ProjectStatusCreate struct {
	Name string `json:"name"`
}

func (p ProjectStatusCreate) ToDB() ProjectStatusDB {
	return ProjectStatusDB{
		Name: p.Name,
	}
}

func (ProjectStatusDB) TableName() string {
	return "project_status"
}
