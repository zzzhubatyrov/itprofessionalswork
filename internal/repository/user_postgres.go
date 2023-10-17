package repository

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"ipw-clean-arch/internal/model"
	"log"
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

//func (u *UserPostgres) CreateResponse(data model.User, claims *jwt.RegisteredClaims) (*model.User, error) {
//	//TODO implement me
//	panic("implement me")
//}

func (u *UserPostgres) CreateResponse(data *model.Response) (*model.Response, error) {
	if err := u.db.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (u *UserPostgres) CreateCompany(company *model.Company, user *model.User) (*model.Company, error) {
	if user.Company != nil {
		return nil, errors.New("Пользователь уже создал компанию")
	} else if err := u.db.Create(&company).Error; err != nil {
		return nil, err
	}
	return company, nil
}

func (u *UserPostgres) GetVacancy() {
	//TODO implement me
	panic("implement me")
}
