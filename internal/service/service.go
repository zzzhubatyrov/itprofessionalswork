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
	UpdateUser(data model.User, secretKey string, c *fiber.Ctx) (*model.User, error)
	CreateResponse(data model.Vacancy, secretKey string, c *fiber.Ctx) (*model.Response, error)
	ResumeHandler
	CompanyHandler
}

type ResumeHandler interface {
	CreateResume(data model.Resume, secretKey string, c *fiber.Ctx) (*model.Resume, error)
	UpdateResume(data model.Resume, id, secretKey string, c *fiber.Ctx) (*model.Resume, error)
	GetResume()
	GetResumeByID(id string) (*model.Resume, error)
	GetAllResumes(data []model.Resume) ([]model.Resume, error)
	DeleteResume(id string) error
}

// RoleHandler TODO Add GetUserRole()
type RoleHandler interface {
	GetAllRoles(data []model.Role) ([]model.Role, error)
}

type CompanyHandler interface {
	//CreateCompany(data model.Company, secretKey string, ctx *fiber.Ctx) (*model.Company, error)
	CreateCompany(data model.Company, secretKey string, c *fiber.Ctx) (*model.Company, error)
	//GetVacancy()
	//GetVacancyByID()
	//GetAllVacancy()
}

type VacancyHandler interface {
	CreateVacancy(data model.Vacancy) (*model.Vacancy, error)
	GetAllVacancy(data []model.Vacancy) ([]model.Vacancy, error)
	GetVacancyByID(id string) (*model.Vacancy, error)
	UpdateVacancy()
	DeleteVacancy()
}

type Service struct {
	Authorization
	UserHandler
	RoleHandler
	VacancyHandler
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:  NewAuthService(repos.Authorization),
		UserHandler:    NewUserService(repos.UserHandler),
		VacancyHandler: NewVacancyService(repos.VacancyHandler),
		RoleHandler:    NewRoleService(repos.RoleHandler),
	}
}
