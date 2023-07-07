package dto

import "time"

type User struct {
	ID          int
	ResumeID    int
	Email       string
	Password    []byte
	Name        string
	Age         uint8
	Tag         string
	Description string
	Number      int
	Gender      string
	Birthday    *time.Time
	City        string
	Role        string
	Photo       []byte
}
