package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"ipw-app/internal/handler"
	"ipw-app/internal/model"
	"ipw-app/internal/repository"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ошибка чтения файла конфигурации: %s", err))
	}
}

var (
	userHandler handler.UserHandler  = new(model.User)
	connect     repository.DBConnect = new(repository.GormConnect)
	secretKey                        = viper.GetString("SecretKey")
)

func User(c *fiber.Ctx) error {
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	return userHandler.User(db, secretKey, c)
}

func Register(c *fiber.Ctx) error {
	var data map[string]string
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	user, err := userHandler.Register(data, db)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	return userHandler.Login(data, db, secretKey, c)
}

func Logout(c *fiber.Ctx) error {
	return userHandler.Logout(c)
}

// UploadPhoto
// TODO Warning test code for upload photo
func UploadPhoto(c *fiber.Ctx) error {
	var data model.User
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	var userPhoto model.UploadHandler = new(model.User)
	return userPhoto.UserUploadPhoto(c, db)
}

func GetRole(c *fiber.Ctx) error {
	tag := c.Params("@tag")
	//getTag := model.User{Tag: tag}

	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	var userRoleHandler model.RoleHandler = new(model.Role)
	getTag, err := userRoleHandler.GetByTag(tag, db)
	if err != nil {
		return err
	}
	return c.JSON(getTag)
}

//func ChangeRole(c *fiber.Ctx) error {
//	db, err := connect.Connect()
//	if err != nil {
//		panic(err)
//	}
//
//	return nil
//}
