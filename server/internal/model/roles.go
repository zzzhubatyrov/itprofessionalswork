package model

import (
	"gorm.io/gorm"
)

type RoleHandler interface {
	GetByTag(tag string, db *gorm.DB) (*Role, error)
	//Update(tag *model.User, db *gorm.DB) error
}

type Role struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (r *Role) GetByTag(tag string, db *gorm.DB) (*Role, error) {
	var role Role
	getTag := db.Select("role").Model(&User{}).Where("tag = ?", tag).First(&role)
	if getTag.Error != nil {
		return nil, getTag.Error
	}
	return &role, nil
}
