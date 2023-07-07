package model

type Resume struct {
	ID          int    `json:"resume_id" gorm:"primaryKey"`
	Photo       []byte `json:"user_photo"`
	Name        string `json:"username"`
	Tag         string `json:"user_tag"`
	ResumeName  string `json:"resume_name"`
	Location    string `json:"user_location"`
	Status      string `json:"user_status"`
	Email       string `json:"user_email"`
	Salary      int    `json:"salary"`
	Description string `json:"description"`
}
