package service

import (
	"ipw-clean-arch/internal/model"
	"ipw-clean-arch/internal/repository"
)

type AdminServices struct {
	repo repository.AdminHandler
}

func NewAdminService(repo repository.AdminHandler) *AdminServices {
	return &AdminServices{repo: repo}
}

func (a *AdminServices) UpdateUserRoleByID(id int) (*model.User, error) {
	updateUserRoleById, err := a.repo.UpdateUserRoleByID(id)
	if err != nil {
		return nil, err
	}
	return updateUserRoleById, nil
}
