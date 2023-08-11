package repository

import (
	"gorm.io/gorm"
	"ipw-clean-arch/internal/model"
)

type RolePostgres struct {
	db *gorm.DB
}

func NewRolePostgres(db *gorm.DB) *RolePostgres {
	return &RolePostgres{db: db}
}

func (r *RolePostgres) GetAllRoles(data []model.Role) ([]model.Role, error) {
	if err := r.db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
