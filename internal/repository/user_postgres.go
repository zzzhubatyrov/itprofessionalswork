package repository

import (
	"gorm.io/gorm"
	"ipw-clean-arch/internal/model"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u UserPostgres) GetUser(data model.User) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserPostgres) GetAllUsers(data []model.User) ([]model.User, error) {
	//TODO implement me
	panic("implement me")
}
