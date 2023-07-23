package model

type Vacancy struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	CompanyID   int    `json:"company_id"`
	CompanyName string `json:"company_name"`
	CompanyTag  string `json:"company_tag" gorm:"unique;"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	WorkTime    string `json:"work_time"`
	Description string `json:"description"`
	Skills      string `json:"skills"`
}
