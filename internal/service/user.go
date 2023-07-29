package service

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"ipw-clean-arch/internal/model"
	"ipw-clean-arch/internal/repository"
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
