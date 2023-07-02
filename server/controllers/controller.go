package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"ipw-app/config"
	usermodel "ipw-app/models"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".\\config")
	// Чтение файла конфигурации
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ошибка чтения файла конфигурации: %s", err))
	}
}

var (
	SecretKey                  = viper.GetString("SecretKey")
	connector config.DBConnect = &config.GormConnect{}
)

func Register(c *fiber.Ctx) error {
	var data usermodel.User
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	user, err := data.UserRegister(&data)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data usermodel.User
	if dataErr := c.BodyParser(&data); dataErr != nil {
		return dataErr
	}
	if err := data.UserLogin(data, c); err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"cookie":  c.Cookies("jwt"),
	})
}

func User(c *fiber.Ctx) error {
	db, err := connector.Connect()
	if err != nil {
		fmt.Println(err)
	}
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil // using the SecretKey which was generated in th Login function
	})
	var user usermodel.User
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
			"data":    &user,
			"cookie":  c.Cookies("jwt"),
		})
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	db.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(fiber.Map{
		"data":   &user,
		"cookie": c.Cookies("jwt"),
	})
}

func Logout(c *fiber.Ctx) error {
	var data usermodel.User
	if err := data.UserLogout(c); err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "success logout",
	})
}
