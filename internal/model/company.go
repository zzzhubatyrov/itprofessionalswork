package model

type Company struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	UserID      int       `json:"userID"`
	Name        string    `json:"name"`
	Tag         string    `json:"tag"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	Vacancy     []Vacancy `json:"vacancies"`
	//User        *User     `json:"-" gorm:"foreignKey:UserID"`
}
