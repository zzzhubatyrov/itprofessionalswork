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
	ResumeHandler
	GetUser(data model.User, secretKey string, c *fiber.Ctx) (*model.User, error)
	GetAllUsers(data []model.User) ([]model.User, error)
}

type ResumeHandler interface {
	CreateResume(data model.Resume, secretKey string, c *fiber.Ctx) (*model.Resume, error)
	UpdateResume(data model.Resume, id, secretKey string, c *fiber.Ctx) (*model.Resume, error)
	GetResume()
	GetAllResumes(data []model.Resume) ([]model.Resume, error)
	DeleteResume()
}

// RoleHandler TODO Add GetUserRole()
type RoleHandler interface {
	GetAllRoles(data []model.Role) ([]model.Role, error)
}

type CompanyHandler interface {
	CreateCompany()
	GetVacancy()
	GetVacancyByID()
	GetAllVacancy()
}

type VacancyHandler interface {
	CreateVacancy()
	UpdateVacancy()
	DeleteVacancy()
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
		RoleHandler:   NewRoleService(repos.RoleHandler),
	}
}
