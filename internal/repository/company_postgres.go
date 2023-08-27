package repository

import (
	"gorm.io/gorm"
	"ipw-clean-arch/internal/model"
)

type CompanyPostgres struct {
	db *gorm.DB
}

func NewCompanyPostgres(db *gorm.DB) *CompanyPostgres {
	return &CompanyPostgres{db: db}
}

func (c *CompanyPostgres) CreateCompany(data *model.Company) (*model.Company, error) {
	if err := c.db.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (c *CompanyPostgres) GetVacancy() {
	//TODO implement me
	panic("implement me")
}
