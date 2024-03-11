package repository

import (
	"github.com/redis/go-redis/v9"
	"ipw-clean-arch/internal/model"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// Authorization TODO add Update method, CheckEmail, CheckUser, VerifyEmail
type Authorization interface {
	Register(data *model.User) (*model.User, error)
	Login(data model.User) (*model.User, error)
}

// UserHandler TODO Add GetUserRole()
type UserHandler interface {
	GetUser(data model.User, claims *jwt.RegisteredClaims) (*model.User, error)
	GetUserByTag(tag string) (*model.User, error)
	GetAllUsers(data []model.User) ([]model.User, error)
	UpdateUser(data model.User, claims *jwt.RegisteredClaims) (*model.User, error)
	CreateResponse(data *model.Response) (*model.Response, error)
	UploadPhoto(claims *jwt.RegisteredClaims, photoData []byte) (*model.User, error)
	ResumeHandler
	CompanyHandler
}

type ResumeHandler interface {
	CreateResume(data *model.Resume) (*model.Resume, error)
	UpdateResume(data *model.Resume, id string) (*model.Resume, error)
	GetResume()
	GetResumeByID(id int) (*model.Resume, error)
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
	//UpdateCompanyData(company model.Company, user model.User, claims *jwt.RegisteredClaims) (*model.Company, error)
	UpdateCompanyData(company *model.Company, user *model.User) (*model.Company, error)
	GetAllCompanies(company []model.Company) ([]model.Company, error)
	GetCompanyByID(id string) (*model.Company, error)
	VacancyHandler
}

type VacancyHandler interface {
	CreateVacancy(data model.Vacancy, claims *jwt.RegisteredClaims) (*model.Vacancy, error)
	GetAllVacancy(data []model.Vacancy) ([]model.Vacancy, error)
	GetVacancyByID(id string) (*model.Vacancy, error)
	UpdateVacancy()
	DeleteVacancy()
}

type AdminHandler interface {
	UpdateUserRoleByID(id int) (*model.User, error)
}

type Repository struct {
	Authorization
	UserHandler
	RoleHandler
	AdminHandler
}

func NewRepository(db *gorm.DB, client *redis.Client) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		UserHandler:   NewUserPostgres(db, client),
		RoleHandler:   NewRolePostgres(db),
		AdminHandler:  NewAdminPostgres(db),
	}
}
