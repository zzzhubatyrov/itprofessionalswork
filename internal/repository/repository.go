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
	UpdateUser(data model.User, claims *jwt.RegisteredClaims) (*model.User, error)
	CreateResponse(data *model.Response) (*model.Response, error)
	ResumeHandler
	CompanyHandler
}

type ResumeHandler interface {
	CreateResume(data *model.Resume) (*model.Resume, error)
	UpdateResume(data *model.Resume, id string) (*model.Resume, error)
	GetResume()
	GetResumeByID(id string) (*model.Resume, error)
	GetAllResumes(data []model.Resume) ([]model.Resume, error)
	DeleteResume(id string) error
}

// RoleHandler TODO Add GetUserRole(), CheckUserRole()
type RoleHandler interface {
	GetAllRoles(data []model.Role) ([]model.Role, error)
}

type CompanyHandler interface {
	UpdateRoleByUserID(userID string, roleID int) error
	CreateCompany(company *model.Company, user *model.User, claims *jwt.RegisteredClaims) (*model.Company, error)
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

type Repository struct {
	Authorization
	UserHandler
	RoleHandler
	VacancyHandler
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization:  NewAuthPostgres(db),
		UserHandler:    NewUserPostgres(db),
		VacancyHandler: NewVacancyPostgres(db),
		RoleHandler:    NewRolePostgres(db),
	}
}
