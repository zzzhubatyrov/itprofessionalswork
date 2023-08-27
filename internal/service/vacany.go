package service

import (
	"ipw-clean-arch/internal/model"
	"ipw-clean-arch/internal/repository"
)

type VacancyServices struct {
	repo repository.VacancyHandler
}

func NewVacancyService(repo repository.VacancyHandler) *VacancyServices {
	return &VacancyServices{repo: repo}
}

func (v *VacancyServices) CreateVacancy(data model.Vacancy) (*model.Vacancy, error) {
	createVacancy, err := v.repo.CreateVacancy(data)
	if err != nil {
		return nil, err
	}
	return createVacancy, nil
}

func (v *VacancyServices) GetAllVacancy(data []model.Vacancy) ([]model.Vacancy, error) {
	getAllVacancy, err := v.repo.GetAllVacancy(data)
	if err != nil {
		return nil, err
	}
	return getAllVacancy, nil
}

func (v *VacancyServices) GetVacancyByID(id string) (*model.Vacancy, error) {
	getVacancyByID, err := v.repo.GetVacancyByID(id)
	if err != nil {
		return nil, err
	}
	return getVacancyByID, nil
}

func (v *VacancyServices) UpdateVacancy() {
	//TODO implement me
	panic("implement me")
}

func (v *VacancyServices) DeleteVacancy() {
	//TODO implement me
	panic("implement me")
}
