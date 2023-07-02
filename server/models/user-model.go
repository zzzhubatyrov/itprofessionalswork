package models

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"ipw-app/config"
	"strconv"
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey; serial"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);unique;not null;column:email"`
	Password  []byte    `json:"-" gorm:"type:varchar(255);not null;column:password"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	SecretKey                  = viper.GetString("SecretKey")
	connector config.DBConnect = &config.GormConnect{}
)

func (user *User) UserRegister(data *User) (*User, error) {
	db, err := connector.Connect()
	if err != nil {
		return nil, err
	}

	// Проверка наличия уже зарегистрированного пользователя с указанным email
	db.Where("email = ?", data.Email).First(user)
	if user.ID != 0 {
		return nil, errors.New("user already exists")
	}

	password, err := bcrypt.GenerateFromPassword(data.Password, 15)
	if err != nil {
		return nil, err
	}

	regUser := &User{
		Name:     data.Name,
		Email:    data.Email,
		Password: password,
	}

	db.Create(&regUser)
	return regUser, nil
}

func (user *User) UserLogin(data User, c *fiber.Ctx) error {
	db, err := connector.Connect()
	if err != nil {
		fmt.Println(err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not connect to database",
		})
	}
	db.Where("email = ?", data.Email).First(user)
	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, data.Password); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})
	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not generate token",
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return nil
}

func (user *User) GetUser(c *fiber.Ctx) error {
	db, err := connector.Connect()
	if err != nil {
		fmt.Println(err)
	}
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil // using the SecretKey which was generated in th Login function
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	db.Where("id = ?", claims.Issuer).First(user)

	return c.JSON(user)
}

func (user *User) UserLogout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Sets the expiry time an hour ago in the past.
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return nil
}
