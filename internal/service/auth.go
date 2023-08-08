package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"ipw-clean-arch/internal/model"
	"ipw-clean-arch/internal/repository"
	"strconv"
	"time"
)

type AuthServices struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthServices {
	return &AuthServices{repo: repo}
}

func (u *AuthServices) Register(data model.User) (*model.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(data.Password), 15)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Email:    data.Email,
		Password: string(password),
		Name:     data.Name,
		Age:      data.Age,
		Tag:      data.Tag,
		RoleID:   4,
	}
	//var role Role
	//roleResult := db.First(&role, regUser.RoleID)
	//if roleResult.Error != nil {
	//	return nil, fmt.Errorf("failed to find role: %v", roleResult.Error)
	//}
	//role.UserCount++
	//if err := db.Save(&role).Error; err != nil {
	//	return nil, err
	//}
	regUser, err := u.repo.Register(user)
	if err != nil {
		return nil, err
	}
	return regUser, nil
}

func (u *AuthServices) Login(data model.User, secretKey string, c *fiber.Ctx) error {
	user, err := u.repo.Login(data)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(user.ID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})
	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not generate token",
		})
	}
	cookie := fiber.Cookie{
		Name:     "ipw_cookie",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	if cookie.Value != "" {
		c.Set("Authorization", "Bearer "+cookie.Value)
	}
	return c.JSON(fiber.Map{
		"user":        user,
		"cookie_name": cookie.Name,
		"token":       cookie.Value,
	})
}

func (u *AuthServices) Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "ipw_cookie",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success logout",
	})
}