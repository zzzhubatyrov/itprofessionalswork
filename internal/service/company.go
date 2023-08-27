package service

import (
	"ipw-clean-arch/internal/model"
	"ipw-clean-arch/internal/repository"
)

type CompanyService struct {
	repo repository.CompanyHandler
}

func NewCompanyService(repo repository.CompanyHandler) *CompanyService {
	return &CompanyService{repo: repo}
}

func (c *CompanyService) CreateCompany(data model.Company) (*model.Company, error) {
	company := &model.Company{
		Name:        data.Name,
		Tag:         data.Tag,
		Email:       data.Email,
		Phone:       data.Phone,
		Location:    data.Location,
		Description: data.Description,
	}
	createCompany, err := c.repo.CreateCompany(company)
	if err != nil {
		return nil, err
	}
	return createCompany, nil
}
