package repository

import (
	"gorm.io/gorm"
	"ipw-clean-arch/internal/model"
)

type VacancyPostgres struct {
	db *gorm.DB
}

func NewVacancyPostgres(db *gorm.DB) *VacancyPostgres {
	return &VacancyPostgres{db: db}
}

func (v *VacancyPostgres) CreateVacancy(data model.Vacancy) (*model.Vacancy, error) {
	vacancy := &model.Vacancy{
		CompanyID:   data.CompanyID,
		CompanyName: data.CompanyName,
		CompanyTag:  data.CompanyTag,
		Direction:   data.Direction,
		Level:       data.Level,
		Location:    data.Location,
		WorkTime:    data.WorkTime,
		Description: data.Description,
		Skills:      data.Skills,
	}
	if err := v.db.Create(vacancy).Error; err != nil {
		return nil, err
	}
	return vacancy, nil
}

func (v *VacancyPostgres) GetAllVacancy(data []model.Vacancy) ([]model.Vacancy, error) {
	if err := v.db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (v *VacancyPostgres) GetVacancyByID(id string) (*model.Vacancy, error) {
	var vacancy model.Vacancy
	if err := v.db.First(&vacancy, id).Error; err != nil {
		return nil, err
	}
	return &vacancy, nil
}

func (v *VacancyPostgres) UpdateVacancy() {
	//TODO implement me
	panic("implement me")
}

func (v *VacancyPostgres) DeleteVacancy() {
	//TODO implement me
	panic("implement me")
}
