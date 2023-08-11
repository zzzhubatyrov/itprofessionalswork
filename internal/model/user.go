package model

type User struct {
	ID       int      `json:"id" gorm:"primaryKey"`
	RoleID   int      `json:"role_id"`
	ResumeID int      `json:"resume_id"`
	Email    string   `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password string   `json:"password" gorm:"type:varchar(255);not null"`
	Name     string   `json:"name" gorm:"type:varchar(255)"`
	Age      string   `json:"age" gorm:"type:varchar(255)"`
	Tag      string   `json:"tag" gorm:"type:varchar(255);unique"`
	Role     Role     `json:"role" gorm:"foreignKey:RoleID"`
	Resume   []Resume `json:"resume"`
}
