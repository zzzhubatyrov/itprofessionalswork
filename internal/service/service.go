package service

import (
	"gorm.io/gorm"
	"ipw-clean-arch/internal/model"
	"ipw-clean-arch/internal/repository"
)

type Authorization interface {
	Register(data map[string]string, db *gorm.DB) (*model.User, error)
	GenerateToken()
	ParseToken()
	//ChangeRole(db *gorm.DB) (*model.Role, error)
}

// RoleHandler TODO Add GetUserRole()
type RoleHandler interface {
	CompanyHandler
	//GetAllRoles(db *gorm.DB) ([]model.Role, error)
	//CreateVacancy(data model.Vacancy, db *gorm.DB) (*model.Vacancy, error)
	//UpdateVacancy()
	//DeleteVacancy()
}

type CompanyHandler interface {
	//CreateCompany(data model.Company, secretKey string, db *gorm.DB, c *fiber.Ctx) (*model.Company, error)
	//GetVacancy(db *gorm.DB) (*model.Vacancy, error)
	//GetVacancyByID()
}

type Service struct {
	Authorization
	RoleHandler
	CompanyHandler
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
