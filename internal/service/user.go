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
	cookie := c.Cookies("ipw_cookie")
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

func (u *UserServices) GetAllUsers(data []model.User) ([]model.User, error) {
	getAllUsers, err := u.repo.GetAllUsers(data)
	if err != nil {
		return nil, err
	}
	return getAllUsers, nil
}

func (u *UserServices) CreateResume(data model.Resume, secretKey string, c *fiber.Ctx) (*model.Resume, error) {
	var user model.User
	cookie := c.Cookies("ipw_cookie")
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
		UserID:      claims.Issuer,
		UserEmail:   data.UserEmail,
		UserName:    data.UserName,
		UserTag:     data.UserTag,
		Direction:   data.Direction,
		Level:       data.Level,
		Salary:      data.Salary,
		Location:    data.Location,
		Status:      data.Status,
		Description: data.Description,
	}
	createResume, err := u.repo.CreateResume(resume)
	if err != nil {
		return nil, err
	}
	return createResume, nil
}

func (u *UserServices) UpdateResume(data model.Resume, id, secretKey string, c *fiber.Ctx) (*model.Resume, error) {
	var user model.User
	cookie := c.Cookies("ipw_cookie")
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

func (u *UserServices) GetAllResumes(data []model.Resume) ([]model.Resume, error) {
	getAllResumes, err := u.repo.GetAllResumes(data)
	if err != nil {
		return nil, err
	}
	return getAllResumes, nil
}

func (u *UserServices) DeleteResume() {
	//TODO implement me
	panic("implement me")
}
