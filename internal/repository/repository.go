package repository

import (
	"github.com/golang-jwt/jwt/v4"
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
	GetUser(data model.User, claims *jwt.RegisteredClaims) (*model.User, error)
	GetAllUsers(data []model.User) ([]model.User, error)
}

// RoleHandler TODO Add GetUserRole(), CheckUserRole()
type RoleHandler interface {
	GetAllRoles(data []model.Role) ([]model.Role, error)
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
		RoleHandler:   NewRolePostgres(db),
	}
}
