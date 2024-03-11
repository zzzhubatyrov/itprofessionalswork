package model

type Resume struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	UserID      int    `json:"userID"`
	Email       string `json:"email" gorm:"-"`
	Name        string `json:"name" gorm:"-"`
	Age         string `json:"age" gorm:"-"`
	Gender      string `json:"gender" gorm:"-"`
	Tag         string `json:"tag" gorm:"-"`
	Number      string `json:"number" gorm:"-"`
	Direction   string `json:"direction"`
	Level       string `json:"level"`
	Salary      string `json:"salary"`
	Location    string `json:"location"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Skills      string `json:"skills"`
}
