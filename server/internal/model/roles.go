package model

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ipw-app/internal/services"
)

type Role struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	UserCount uint64 `json:"user_count" gorm:"default:0"`
}

var admRoles = []string{"Администратор", "Модератор"}
var allRoles = []string{"Администратор", "Модератор", "Менеджер", "Пользователь"}

func contains(roles []string, role string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}

func (r Role) CreateCompany(data Company, secretKey string, db *gorm.DB, c *fiber.Ctx) (*Company, error) {
	cookie := c.Cookies("ipw_cookie")
	_, err := services.VerifyJWT(secretKey, cookie, c)
	if err != nil {
		return nil, err
	}
	var user User
	if err := db.Preload("Role").First(&user).Error; err != nil {
		return nil, err
	}
	if !contains(admRoles, user.Role.Name) {
		return nil, errors.New("недостаточно прав доступа для создания компании")
	}
	company := &Company{
		Name:  data.Name,
		Email: data.Email,
		Tag:   data.Tag,
	}
	if err := db.Create(company).Error; err != nil {
		return nil, err
	}
	return company, nil
}

func (r Role) GetAllRoles(db *gorm.DB) ([]Role, error) {
	var roles []Role
	if err := db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r Role) GetVacancy(db *gorm.DB) (*Vacancy, error) {
	var vacancy Vacancy
	if err := db.Find(&vacancy).Error; err != nil {
		return nil, err
	}
	return &vacancy, nil
}

func (r Role) GetVacancyByID() {
	//TODO implement me
	panic("implement me")
}

func (r Role) CreateVacancy(data Vacancy, db *gorm.DB) (*Vacancy, error) {
	company := &Company{}
	if err := db.First(company, data.CompanyID).Error; err != nil {
		return nil, err
	}
	vacancy := &Vacancy{
		CompanyID:   data.CompanyID,
		CompanyName: company.Name,
		CompanyTag:  company.Tag,
		Name:        data.Name,
		Location:    data.Location,
		WorkTime:    data.WorkTime,
	}
	if err := db.Create(vacancy).Error; err != nil {
		return nil, err
	}
	return vacancy, nil
}

func (r Role) UpdateVacancy() {
	panic("implement me")
}

func (r Role) DeleteVacancy() {
	//TODO implement me
	panic("implement me")
}
