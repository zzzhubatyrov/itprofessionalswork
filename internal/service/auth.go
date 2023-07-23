package service

import (
	"gorm.io/gorm"
	"ipw-clean-arch/internal/model"
	"ipw-clean-arch/internal/repository"
)

type UserServices struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *UserServices {
	return &UserServices{repo: repo}
}

func (u UserServices) GenerateToken() {
	//TODO implement me
	panic("implement me")
}

func (u UserServices) ParseToken() {
	//TODO implement me
	panic("implement me")
}

func (u UserServices) Register(data map[string]string, db *gorm.DB) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
