package model

type Response struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userID"`
	VacancyID   int    `json:"vacancyID"`
	CompanyName string `json:"companyName"`
	JobTitle    string `json:"jobTitle"`
	Applicant   string `json:"applicant"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}
