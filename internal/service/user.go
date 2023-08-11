package service

import (
	"ipw-clean-arch/internal/model"
	"ipw-clean-arch/internal/repository"
)

type UserServices struct {
	repo repository.UserHandler
}

func NewUserService(repo repository.UserHandler) *UserServices {
	return &UserServices{repo: repo}
}

func (u UserServices) GetUser(data model.User) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServices) GetAllUsers(data []model.User) ([]model.User, error) {
	//TODO implement me
	panic("implement me")
}
