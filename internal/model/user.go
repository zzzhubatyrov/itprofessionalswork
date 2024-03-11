package model

type User struct {
	ID          int      `json:"id" gorm:"primaryKey"`
	RoleID      int      `json:"roleID"`
	Email       string   `json:"email" gorm:"unique;not null"`
	Password    string   `json:"password" gorm:"not null"`
	Name        string   `json:"name"`
	Birthday    string   `json:"birthday"`
	Tag         string   `json:"tag" gorm:"unique"`
	Photo       []byte   `json:"photo"`
	Number      string   `json:"number"`
	Location    string   `json:"location"`
	Description string   `json:"description"`
	Gender      string   `json:"gender"`
	Stack       string   `json:"stack"`
	Blocked     bool     `json:"blocked" gorm:"default:false"`
	Role        Role     `json:"role" gorm:"foreignKey:RoleID"`
	Company     *Company `json:"company" gorm:"foreignKey:UserID"`
	Resume      []Resume `json:"resumes"`
	//Response    []Response `json:"responses"`
	//Chats       []Chat     `json:"chats" gorm:"many2many:user_chats"`
}
