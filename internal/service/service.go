package service

import (
	"github.com/gofiber/fiber/v2"
	"ipw-clean-arch/internal/model"
	"ipw-clean-arch/internal/repository"
)

// Authorization TODO add Update method, CheckEmail, CheckUser, VerifyEmail
type Authorization interface {
	Register(data model.User) (*model.User, error)
	Login(data model.User, secretKey string, c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

type UserHandler interface {
	GetUser(data model.User, secretKey string, c *fiber.Ctx) (*model.User, error)
	GetAllUsers(data []model.User) ([]model.User, error)
}

// RoleHandler TODO Add GetUserRole()
type RoleHandler interface {
	//GetAllRoles(data []model.Role) ([]model.Role, error)
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
	UserHandler
	RoleHandler
	CompanyHandler
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		UserHandler:   NewUserService(repos.UserHandler),
	}
}
