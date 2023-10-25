package model

type Company struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	UserID      int       `json:"userID"`
	Photo       []byte    `json:"photo"`
	Name        string    `json:"name" gorm:"unique"`
	Tag         string    `json:"tag" gorm:"unique"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	CompanySize string    `json:"companySize"`
	WebSite     string    `json:"webSite"`
	Vacancy     []Vacancy `json:"vacancies"`
	//User        *User     `json:"-" gorm:"foreignKey:UserID"`
}
