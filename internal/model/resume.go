package model

type Resume struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	UserID      string `json:"user_id"`
	UserEmail   string `json:"user_email"`
	UserName    string `json:"user_name"`
	UserTag     string `json:"user_tag"`
	Direction   string `json:"direction"`
	Level       string `json:"level"`
	Salary      string `json:"salary"`
	Location    string `json:"location"`
	Status      string `json:"status"`
	Description string `json:"description"`
}
