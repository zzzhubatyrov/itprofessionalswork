package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"ipw-clean-arch/internal/model"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (u *AuthPostgres) Register(data *model.User) (*model.User, error) {
	var existingUser model.User
	result := u.db.Where("email = ?", data.Email).First(&existingUser)
	if result.Error == nil {
		return nil, fmt.Errorf("user already exists")
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	if err := u.db.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (u *AuthPostgres) Login(data model.User) (*model.User, error) {
	var user model.User
	if err := u.db.Preload("Role").Where("email = ?", data.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
