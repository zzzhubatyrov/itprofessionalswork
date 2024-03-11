package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"ipw-clean-arch/internal/model"
	"log"
	"strconv"
	"time"
)

type UserPostgres struct {
	db *gorm.DB
	rb *redis.Client
}

func NewUserPostgres(db *gorm.DB, rb *redis.Client) *UserPostgres {
	return &UserPostgres{db: db, rb: rb}
}

func (u *UserPostgres) GetUser(data model.User, claims *jwt.RegisteredClaims) (*model.User, error) {
	if err := u.db.Preload("Resume").Preload("Company").Preload("Company.Vacancy").Preload("Response").Preload("Role").Where("id = ?", claims.Issuer).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (u *UserPostgres) GetUserByTag(tag string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

/*
GetAllUsers Сделать через switch

Исп. errors.Is

Использовать Generics
*/
func (u *UserPostgres) GetAllUsers(data []model.User) ([]model.User, error) {
	//if err := u.db.Preload("Resume").Preload("Role").Find(&data).Error; err != nil {
	//	return nil, err
	//}
	//
	// Сделать через switch
	// Исп. errors.Is
	// Использовать Generics
	ctx := context.Background()
	cachedUsers, err := u.rb.Get(ctx, "users").Result()
	if errors.Is(err, redis.Nil) {
		// Данные отсутствуют в кеше, получение данных из PostgreSQL
		if err := u.db.Preload("Resume").Preload("Role").Find(&data).Error; err != nil {
			return nil, err
		}
		// Сохранение данных в кеше Redis
		usersJSON, err := json.Marshal(data)
		err = u.rb.Set(ctx, "users", usersJSON, time.Minute).Err()
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	} else {
		// Данные найдены в кеше, использование их
		log.Println("Cache hit for users")
		err = json.Unmarshal([]byte(cachedUsers), &data)
		if err != nil {
			return nil, err
		}
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
	if data.Stack != "" {
		user.Stack = data.Stack
	}
	if err := u.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserPostgres) CreateResponse(data *model.Response) (*model.Response, error) {
	if err := u.db.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
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
	if data.Email != "" {
		resume.Email = data.Email
	}
	if data.Name != "" {
		resume.Name = data.Name
	}
	if data.Tag != "" {
		resume.Tag = data.Tag
	}
	if data.Direction != "" {
		resume.Direction = data.Direction
	}
	if data.Level != "" {
		resume.Level = data.Level
	}
	if data.Salary != "" {
		resume.Salary = data.Salary
	}
	if data.Location != "" {
		resume.Location = data.Location
	}
	if data.Status != "" {
		resume.Status = data.Status
	}
	if data.Description != "" {
		resume.Description = data.Description
	}
	if err := u.db.Save(&resume).Error; err != nil {
		return nil, err
	}
	return &resume, nil
}

func (u *UserPostgres) GetResume() {
	//TODO implement me
	panic("implement me")
}

func (u *UserPostgres) GetResumeByID(id int) (*model.Resume, error) {
	var resume model.Resume
	var user model.User
	if err := u.db.First(&resume, id).Error; err != nil {
		return nil, err
	}
	if err := u.db.First(&user, resume.UserID).Error; err != nil {
		return nil, err
	}
	resume.Name = user.Name
	resume.Email = user.Email
	resume.Age = user.Birthday
	resume.Gender = user.Gender
	resume.Tag = user.Tag
	resume.Number = user.Number
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

func (u *UserPostgres) UpdateCompanyData(company *model.Company, user *model.User) (*model.Company, error) {
	var updateCompany model.Company
	if company.ID != 0 {
		updateCompany.ID = company.ID
	}
	if company.UserID != 0 {
		updateCompany.UserID = company.UserID
	}
	if company.Photo != nil {
		updateCompany.Photo = company.Photo
	}
	if company.Name != "" {
		updateCompany.Name = company.Name
	}
	if company.Tag != "" {
		updateCompany.Tag = company.Tag
	}
	if company.Email != "" {
		updateCompany.Email = company.Email
	}
	if company.Phone != "" {
		updateCompany.Phone = company.Phone
	}
	if company.Location != "" {
		updateCompany.Location = company.Location
	}
	if company.Description != "" {
		updateCompany.Description = company.Description
	}
	if company.CompanySize != "" {
		updateCompany.CompanySize = company.CompanySize
	}
	if company.WebSite != "" {
		updateCompany.WebSite = company.WebSite
	}
	if company.Vacancy != nil {
		updateCompany.Vacancy = company.Vacancy
	}
	if err := u.db.Save(&updateCompany).Error; err != nil {
		return nil, err
	}
	return &updateCompany, nil
}

func (u *UserPostgres) GetAllCompanies(company []model.Company) ([]model.Company, error) {
	if err := u.db.Find(&company).Error; err != nil {
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

func (u *UserPostgres) CreateVacancy(data model.Vacancy, claims *jwt.RegisteredClaims) (*model.Vacancy, error) {
	var user model.User
	if err := u.db.Preload("Company").Preload("Role").Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		return nil, err
	}
	if user.Company.ID == 0 {
		return nil, errors.New("user does not have a company and cannot create a vacancy")
	}
	vacancy := &model.Vacancy{
		CompanyID:    user.Company.ID,
		CompanyPhoto: user.Company.Photo,
		CompanyTag:   user.Company.Tag,
		CompanyName:  user.Company.Name,
		Direction:    data.Direction,
		Level:        data.Level,
		Location:     data.Location,
		WorkTime:     data.WorkTime,
		Description:  data.Description,
		Skills:       data.Skills,
		Salary:       data.Salary,
		Experience:   data.Experience,
	}
	if err := u.db.Debug().Create(vacancy).Error; err != nil {
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
