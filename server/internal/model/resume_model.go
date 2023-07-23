package model

type Resume struct {
	ID          uint   `json:"resume_id" gorm:"primaryKey"`
	Photo       []byte `json:"user_photo"`
	Name        string `json:"username"`
	Tag         string `json:"user_tag"`
	ResumeName  string `json:"resume_name"`
	Location    string `json:"user_location"`
	Status      string `json:"user_status"`
	Email       string `json:"user_email"`
	Salary      uint32 `json:"salary"`
	Description string `json:"description"`
}
