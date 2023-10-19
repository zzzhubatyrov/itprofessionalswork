package service

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"ipw-clean-arch/internal/model"
	"ipw-clean-arch/internal/repository"
	"strconv"
)

type UserServices struct {
	repo repository.UserHandler
}

func NewUserService(repo repository.UserHandler) *UserServices {
	return &UserServices{repo: repo}
}

func (u *UserServices) GetUser(data model.User, secretKey string, c *fiber.Ctx) (*model.User, error) {
	cookie := c.Cookies("ipwCookie")
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, errors.New("unauthenticated")
	}
	if !token.Valid {
		return nil, errors.New("недействительный JWT токен")
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	user, err := u.repo.GetUser(data, claims)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserServices) UpdateUser(data model.User, secretKey string, c *fiber.Ctx) (*model.User, error) {
	cookie := c.Cookies("ipwCookie")
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, errors.New("unauthenticated")
	}
	if !token.Valid {
		return nil, errors.New("недействительный JWT токен")
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	updateUser, err := u.repo.UpdateUser(data, claims)
	if err != nil {
		return nil, err
	}
	//if err := elasticsearch.IndexUser(*updateUser); err != nil {
	//	return nil, err
	//}
	return updateUser, nil
}

func (u *UserServices) GetAllUsers(data []model.User) ([]model.User, error) {
	getAllUsers, err := u.repo.GetAllUsers(data)
	if err != nil {
		return nil, err
	}
	return getAllUsers, nil
}

func (u *UserServices) CreateResume(data model.Resume, secretKey string, c *fiber.Ctx) (*model.Resume, error) {
	var user model.User
	cookie := c.Cookies("ipwCookie")
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, errors.New("unauthenticated")
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	getUser, err := u.repo.GetUser(user, claims)
	if !ok {
		return nil, fmt.Errorf("неверный формат токена")
	}
	if claims.Valid() != nil {
		return nil, fmt.Errorf("невалидный токен: %v", claims.Valid())
	}
	userIDStr := strconv.Itoa(getUser.ID)
	if claims.Issuer != userIDStr {
		return nil, errors.New("вы не можете создавать резюме для других пользователей")
	}
	resume := &model.Resume{
		UserID:    claims.Issuer,
		UserEmail: getUser.Email,
		UserName:  getUser.Name,
		//UserAge:     getUser.Age,
		UserGender:  getUser.Gender,
		UserTag:     getUser.Tag,
		Direction:   data.Direction,
		Level:       data.Level,
		Salary:      data.Salary,
		Location:    data.Location,
		Status:      data.Status,
		Description: data.Description,
		Skills:      data.Skills,
	}
	//UserID:      claims.Issuer,
	//UserEmail:   getUser.Email,
	//UserName:    getUser.Name,
	//UserTag:     getUser.Tag,
	//Direction:   data.Direction,
	//Level:       data.Level,
	//Salary:      data.Salary,
	//Location:    data.Location,
	//Status:      data.Status,
	//Description: data.Description,
	//Skills:      data.Skills,
	createResume, err := u.repo.CreateResume(resume)
	if err != nil {
		return nil, err
	}
	return createResume, nil
}

func (u *UserServices) UpdateResume(data model.Resume, id, secretKey string, c *fiber.Ctx) (*model.Resume, error) {
	var user model.User
	cookie := c.Cookies("ipwCookie")
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, errors.New("unauthenticated")
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	getUser, err := u.repo.GetUser(user, claims)
	if !ok {
		return nil, fmt.Errorf("неверный формат токена")
	}
	if claims.Valid() != nil {
		return nil, fmt.Errorf("невалидный токен: %v", claims.Valid())
	}
	userIDStr := strconv.Itoa(getUser.ID)
	if claims.Issuer != userIDStr {
		return nil, errors.New("вы не можете создавать резюме для других пользователей")
	}
	resume := &model.Resume{
		UserEmail:   data.UserEmail,
		UserName:    data.UserName,
		UserTag:     data.UserTag,
		Direction:   data.Direction,
		Level:       data.Level,
		Salary:      data.Salary,
		Location:    data.Location,
		Status:      data.Status,
		Description: data.Description,
		Skills:      data.Skills,
	}
	updateResume, err := u.repo.UpdateResume(resume, id)
	if err != nil {
		return nil, err
	}
	return updateResume, nil
}

func (u *UserServices) GetResume() {
	//TODO implement me
	panic("implement me")
}

func (u *UserServices) GetResumeByID(id string) (*model.Resume, error) {
	getResumeByID, err := u.repo.GetResumeByID(id)
	if err != nil {
		return nil, err
	}
	return getResumeByID, nil
}

func (u *UserServices) GetAllResumes(data []model.Resume) ([]model.Resume, error) {
	getAllResumes, err := u.repo.GetAllResumes(data)
	if err != nil {
		return nil, err
	}
	return getAllResumes, nil
}

func (u *UserServices) DeleteResume(id string) error {
	return u.repo.DeleteResume(id)
}

//func (u *UserServices) CreateResponse(data model.User, secretKey string, c *fiber.Ctx) (*model.User, error) {
//	//TODO implement me
//	panic("implement me")
//}

func (u *UserServices) CreateResponse(data model.Vacancy, secretKey string, c *fiber.Ctx) (*model.Response, error) {
	var user model.User
	cookie := c.Cookies("ipwCookie")
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, errors.New("unauthenticated")
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	getUser, err := u.repo.GetUser(user, claims)
	if !ok {
		return nil, fmt.Errorf("неверный формат токена")
	}
	if claims.Valid() != nil {
		return nil, fmt.Errorf("невалидный токен: %v", claims.Valid())
	}
	response := &model.Response{
		UserID:      getUser.ID,
		VacancyID:   data.ID,
		CompanyName: data.CompanyName,
		JobTitle:    data.Level + " " + data.Direction,
		Applicant:   getUser.Name,
		Email:       getUser.Email,
		Phone:       getUser.Number,
	}
	createResponse, err := u.repo.CreateResponse(response)
	if err != nil {
		return nil, err
	}
	return createResponse, nil
}

func (u *UserServices) CreateCompany(data model.Company, secretKey string, c *fiber.Ctx) (*model.Company, error) {
	var user model.User
	cookie := c.Cookies("ipwCookie")
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, errors.New("unauthenticated")
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	getUser, err := u.repo.GetUser(user, claims)
	if !ok {
		return nil, fmt.Errorf("неверный формат токена")
	}
	if claims.Valid() != nil {
		return nil, fmt.Errorf("невалидный токен: %v", claims.Valid())
	}
	company := &model.Company{
		UserID:      getUser.ID,
		Name:        data.Name,
		Tag:         data.Tag,
		Email:       data.Email,
		Phone:       data.Phone,
		Location:    data.Location,
		Description: data.Description,
	}
	createCompany, err := u.repo.CreateCompany(company, getUser, claims)
	if err != nil {
		return nil, err
	}
	return createCompany, nil
}

func (u *UserServices) UpdateRoleByUserID(userID string, roleID int) error {
	updateRoleByUserID := u.repo.UpdateRoleByUserID(userID, roleID)
	return updateRoleByUserID
}
