package model

type Role struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"role_name"`
	User []User `json:"-" gorm:"foreignKey:RoleID"`
}
