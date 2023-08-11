package service

import (
	"ipw-clean-arch/internal/model"
	"ipw-clean-arch/internal/repository"
)

type RoleService struct {
	repo repository.RoleHandler
}

func NewRoleService(repo repository.RoleHandler) *RoleService {
	return &RoleService{repo: repo}
}

func (r *RoleService) GetAllRoles(data []model.Role) ([]model.Role, error) {
	getAllRoles, err := r.repo.GetAllRoles(data)
	if err != nil {
		return nil, err
	}
	return getAllRoles, nil
}
