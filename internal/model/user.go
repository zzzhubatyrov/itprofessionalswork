package model

type User struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	RoleID      int        `json:"roleID"`
	Email       string     `json:"email" gorm:"unique;not null"`
	Password    string     `json:"password" gorm:"not null"`
	Name        string     `json:"name"`
	Age         string     `json:"age"`
	Tag         string     `json:"tag"` //  gorm:"unique"
	Photo       []byte     `json:"photo"`
	Number      string     `json:"number"`
	Location    string     `json:"location"`
	Description string     `json:"description"`
	Gender      string     `json:"gender"`
	Role        Role       `json:"role" gorm:"foreignKey:RoleID"`
	Resume      []Resume   `json:"resumes"`
	Response    []Response `json:"responses"`
}
