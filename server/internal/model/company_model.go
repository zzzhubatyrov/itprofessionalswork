package model

type Company struct {
	ID      int       `json:"id" gorm:"primaryKey"`
	Name    string    `json:"name" gorm:"unique;"`
	Email   string    `json:"email"`
	Tag     string    `json:"tag" gorm:"unique;"`
	Vacancy []Vacancy `json:"vacancies" gorm:"foreignKey:CompanyID"`
}
