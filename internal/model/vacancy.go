package model

type Vacancy struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	CompanyName string `json:"companyName"`
	CompanyTag  string `json:"companyTag"`
	Direction   string `json:"direction"`
	Level       string `json:"level"`
	Location    string `json:"location"`
	WorkTime    string `json:"workTime"`
	Description string `json:"description"`
	Skills      string `json:"skills"`
}
