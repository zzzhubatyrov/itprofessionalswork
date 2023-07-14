package model

type Vacancy struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	CompanyName string `json:"companyName" gorm:"type:varchar(255);not null;column:companyName"`
	CompanyTag  string
	VacancyName string
	Location    string
	WorkTime    string
	Description string
	Skills      string
}
