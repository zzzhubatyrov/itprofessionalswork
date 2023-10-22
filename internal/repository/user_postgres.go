package repository

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"ipw-clean-arch/internal/model"
	"log"
	"strconv"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) GetUser(data model.User, claims *jwt.RegisteredClaims) (*model.User, error) {
	if err := u.db.Preload("Resume").Preload("Company").Preload("Response").Preload("Role").Where("id = ?", claims.Issuer).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (u *UserPostgres) GetAllUsers(data []model.User) ([]model.User, error) {
	if err := u.db.Preload("Resume").Preload("Role").Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (u *UserPostgres) UpdateUser(data model.User, claims *jwt.RegisteredClaims) (*model.User, error) {
	var user model.User
	if err := u.db.Preload("Role").Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		return nil, err
	}
	if data.Birthday != "" {
		user.Birthday = data.Birthday
	}
	if data.Tag != "" {
		user.Tag = data.Tag
	}
	if data.Number != "" {
		user.Number = data.Number
	}
	if data.Location != "" {
		user.Location = data.Location
	}
	if data.Description != "" {
		user.Description = data.Description
	}
	if data.Gender != "" {
		user.Gender = data.Gender
	}
	// Отправляем успешный ответ
	if err := u.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserPostgres) UploadPhoto(claims *jwt.RegisteredClaims, photoData []byte) (*model.User, error) {
	var user model.User
	if err := u.db.Preload("Role").Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		return nil, err
	}
	user.Photo = photoData
	if err := u.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserPostgres) CreateResume(data *model.Resume) (*model.Resume, error) {
	if err := u.db.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (u *UserPostgres) UpdateResume(data *model.Resume, id string) (*model.Resume, error) {
	var resume model.Resume
	if err := u.db.First(&resume, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("resume not found")
		}
		return nil, err
	}
	resume.UserEmail = data.UserEmail
	resume.UserName = data.UserName
	resume.UserTag = data.UserTag
	resume.Direction = data.Direction
	resume.Level = data.Level
	resume.Salary = data.Salary
	resume.Location = data.Location
	resume.Status = data.Status
	resume.Description = data.Description
	if err := u.db.Save(&resume).Error; err != nil {
		return nil, err
	}
	return &resume, nil
}

func (u *UserPostgres) GetResume() {
	//TODO implement me
	panic("implement me")
}

func (u *UserPostgres) GetResumeByID(id string) (*model.Resume, error) {
	var resume model.Resume
	if err := u.db.First(&resume, id).Error; err != nil {
		return nil, err
	}
	return &resume, nil
}

func (u *UserPostgres) GetAllResumes(data []model.Resume) ([]model.Resume, error) {
	if err := u.db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (u *UserPostgres) DeleteResume(id string) error {
	var resume model.Resume
	if err := u.db.First(&resume, id).Delete(resume).Error; err != nil {
		log.Println(err)
		return err
	}
	return errors.New("delete)")
}

func (u *UserPostgres) CreateResponse(data *model.Response) (*model.Response, error) {
	if err := u.db.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (u *UserPostgres) UpdateRoleByUserID(userID string, roleID int) error {
	var user model.User
	if err := u.db.First(&user, userID).Error; err != nil {
		return err
	}
	user.RoleID = roleID
	if err := u.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserPostgres) CreateCompany(company *model.Company, user *model.User, claims *jwt.RegisteredClaims) (*model.Company, error) {
	if err := u.db.Preload("Role").Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		return nil, err
	}
	if user.Company != nil {
		return nil, errors.New("пользователь уже создал компанию")
	} else if err := u.db.Preload("Role").Create(&company).Error; err != nil {
		return nil, err
	}
	if err := u.UpdateRoleByUserID(strconv.Itoa(user.ID), 3); err != nil {
		return nil, err
	}
	return company, nil
}

func (u *UserPostgres) GetCompanyByID(id string) (*model.Company, error) {
	var company model.Company
	if err := u.db.First(&company, id).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

//
//func (u *UserPostgres) UploadPhotoCompany(data model.Company, claims *jwt.RegisteredClaims, photoData []byte) (*model.Company, error) {
//	var company model.Company
//	var user model.User
//	if err := u.db.Preload("Role").Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
//		return nil, err
//	}
//	company.Photo = photoData
//	if err := u.db.Save(&company).Error; err != nil {
//		return nil, err
//	}
//	return &company, nil
//}

func (u *UserPostgres) CreateVacancy(data model.Vacancy) (*model.Vacancy, error) {
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
	if err := u.db.Create(vacancy).Error; err != nil {
		return nil, err
	}
	return vacancy, nil
}

func (u *UserPostgres) GetAllVacancy(data []model.Vacancy) ([]model.Vacancy, error) {
	if err := u.db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (u *UserPostgres) GetVacancyByID(id string) (*model.Vacancy, error) {
	var vacancy model.Vacancy
	if err := u.db.First(&vacancy, id).Error; err != nil {
		return nil, err
	}
	return &vacancy, nil
}

func (u *UserPostgres) UpdateVacancy() {
	//TODO implement me
	panic("implement me")
}

func (u *UserPostgres) DeleteVacancy() {
	//TODO implement me
	panic("implement me")
}
