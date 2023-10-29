package model

type Vacancy struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	CompanyID    int    `json:"companyID"`
	CompanyPhoto []byte `json:"companyPhoto"`
	CompanyName  string `json:"companyName"`
	Direction    string `json:"direction"`
	Level        string `json:"level"`
	Salary       string `json:"salary"`
	Experience   string `json:"experience"`
	Location     string `json:"location"`
	WorkTime     string `json:"workTime"`
	Description  string `json:"description"`
	Skills       string `json:"skills"`
}
