package repository

import (
	"gorm.io/gorm"
)

type CompanyPostgres struct {
	db *gorm.DB
}

func NewCompanyPostgres(db *gorm.DB) *CompanyPostgres {
	return &CompanyPostgres{db: db}
}

func (c CompanyPostgres) CreateCompany() {
	//TODO implement me
	panic("implement me")
}

func (c CompanyPostgres) GetVacancy() {
	//TODO implement me
	panic("implement me")
}
