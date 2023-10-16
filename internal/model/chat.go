package model

type Chat struct {
	ID       int     `json:"id"`
	Users    []*User `gorm:"many2many:user_chats"`
	Messages []Message
}
