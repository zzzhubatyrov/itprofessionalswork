package repository

import (
	"gorm.io/gorm"
	"ipw-clean-arch/internal/model"
)

type AdminPostgres struct {
	db *gorm.DB
}

func NewAdminPostgres(db *gorm.DB) *AdminPostgres {
	return &AdminPostgres{db: db}
}

func (a *AdminPostgres) UpdateUserRoleByID(id int) (*model.User, error) {
	var user model.User
	if err := a.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
