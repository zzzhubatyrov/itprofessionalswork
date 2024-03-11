package service

import (
	"ipw-clean-arch/internal/model"
	"ipw-clean-arch/internal/repository"

	"github.com/gofiber/fiber/v2"
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
	UploadPhoto(secretKey string, c *fiber.Ctx) (*model.User, error)
	ResumeHandler
	CompanyHandler
}

type ResumeHandler interface {
	CreateResume(data model.Resume, secretKey string, c *fiber.Ctx) (*model.Resume, error)
	UpdateResume(data model.Resume, id, secretKey string, c *fiber.Ctx) (*model.Resume, error)
	GetResume()
	GetResumeByID(id int) (*model.Resume, error)
	GetAllResumes(data []model.Resume) ([]model.Resume, error)
	DeleteResume(id string) error
}

// RoleHandler TODO Add GetUserRole()
type RoleHandler interface {
	GetAllRoles(data []model.Role) ([]model.Role, error)
}

type CompanyHandler interface {
	UpdateRoleByUserID(userID string, roleID int) error
	CreateCompany(data model.Company, secretKey string, c *fiber.Ctx) (*model.Company, error)
	UpdateCompanyData(company model.Company, secretKey string, c *fiber.Ctx) (*model.Company, error)
	//UpdateCompanyData(company model.Company) (*model.Company, error)
	GetAllCompanies(company []model.Company) ([]model.Company, error)
	GetCompanyByID(id string) (*model.Company, error)
	VacancyHandler
}

type VacancyHandler interface {
	CreateVacancy(data model.Vacancy, secretKey string, c *fiber.Ctx) (*model.Vacancy, error)
	GetAllVacancy(data []model.Vacancy) ([]model.Vacancy, error)
	GetVacancyByID(id string) (*model.Vacancy, error)
	UpdateVacancy()
	DeleteVacancy()
}

type AdminHandler interface {
	UpdateUserRoleByID(id int) (*model.User, error)
}

type Service struct {
	Authorization
	UserHandler
	RoleHandler
	AdminHandler
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		UserHandler:   NewUserService(repos.UserHandler),
		RoleHandler:   NewRoleService(repos.RoleHandler),
		AdminHandler:  NewAdminService(repos.AdminHandler),
	}
}
