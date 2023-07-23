package repository

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ipw-clean-arch/internal/model"
)

// Authorization TODO add Update method, CheckEmail, CheckUser, VerifyEmail
type Authorization interface {
	GetUser(c *fiber.Ctx) error
	GetUserByID(id string, db *gorm.DB) (*model.User, error)
	Register(data map[string]string, db *gorm.DB) (*model.User, error)
	Login(data map[string]string, db *gorm.DB, secretKey string, c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
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

type Repository struct {
	Authorization
	RoleHandler
	CompanyHandler
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
