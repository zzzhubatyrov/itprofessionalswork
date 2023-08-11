package repository

import (
	"gorm.io/gorm"
	"ipw-clean-arch/internal/model"
)

// Authorization TODO add Update method, CheckEmail, CheckUser, VerifyEmail
type Authorization interface {
	Register(data *model.User) (*model.User, error)
	Login(data model.User) (*model.User, error)
}

// UserHandler TODO Add GetUserRole()
type UserHandler interface {
	GetUser(data model.User) (*model.User, error)
	GetAllUsers(data []model.User) ([]model.User, error)
}

// RoleHandler TODO Add GetUserRole()
type RoleHandler interface {
}

type CompanyHandler interface {
	//CreateCompany()
	//GetVacancy()
	//GetVacancyByID()
}

type Repository struct {
	Authorization
	UserHandler
	RoleHandler
	CompanyHandler
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		UserHandler:   NewUserPostgres(db),
	}
}
