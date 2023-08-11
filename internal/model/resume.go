package model

type Resume struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	UserID      string `json:"userID"`
	UserEmail   string `json:"userEmail"`
	UserName    string `json:"userName"`
	UserTag     string `json:"userTag"`
	Direction   string `json:"direction"`
	Level       string `json:"level"`
	Salary      string `json:"salary"`
	Location    string `json:"location"`
	Status      string `json:"status"`
	Description string `json:"description"`
}
