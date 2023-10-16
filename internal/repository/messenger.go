package repository

import (
	"gorm.io/gorm"
)

type MessengerPostgres struct {
	db *gorm.DB
}

func NewMessengerPostgres(db *gorm.DB) *MessengerPostgres {
	return &MessengerPostgres{db: db}
}

func (m *MessengerPostgres) SendMessage() {
	// TODO implement me
	panic("implement me")
}

func (m *MessengerPostgres) UpdateMessage() {
	// TODO implement me
	panic("implement me")
}

func (m *MessengerPostgres) DeleteMessage() {
	// TODO implement me
	panic("implement me")
}
