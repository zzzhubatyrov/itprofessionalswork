package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"ipw-app/internal/interfaces"
	"ipw-app/internal/model"
	"ipw-app/internal/repository"
)

// #TODO FIXME Полностью переписать handlers

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ошибка чтения файла конфигурации: %s", err))
	}
}

var (
	userHandler interfaces.UserHandler = new(model.User)
	connect     repository.DBConnect   = new(repository.GormConnect)
	roleHandler interfaces.RoleHandler = new(model.Role)
	secretKey                          = viper.GetString("SecretKey")
)

func User(c *fiber.Ctx) error {
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	return userHandler.GetUser(db, secretKey, c)
}
func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	getUserById, err := userHandler.GetUserByID(id, db)
	if err != nil {
		return err
	}
	return c.JSON(getUserById)
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

func CreateCompany(c *fiber.Ctx) error {
	var data model.Company
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	company, err := roleHandler.CreateCompany(data, secretKey, db, c)
	if err != nil {
		return err
	}
	return c.JSON(company)
}

func CreateVacancy(c *fiber.Ctx) error {
	var data model.Vacancy
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	vacancy, err := roleHandler.CreateVacancy(data, db)
	if err != nil {
		return err
	}
	return c.JSON(vacancy)
}

func GetVacancy(c *fiber.Ctx) error {
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	vacancy, err := roleHandler.GetVacancy(db)
	if err != nil {
		return err
	}
	return c.JSON(vacancy)
}

//func GetRole(c *fiber.Ctx) error {
//	db, err := connect.Connect()
//	if err != nil {
//		panic(err)
//	}
//	getUserRole, err := userHandler.GetRole(db, secretKey, c)
//	if err != nil {
//		return err
//	}
//	return c.JSON(getUserRole)
//}

// GetAllRoles get all role in DB
func GetAllRoles(c *fiber.Ctx) error {
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	roles, err := roleHandler.GetAllRoles(db)
	if err != nil {
		return err
	}
	return c.JSON(roles)
}

func ChangeRole(c *fiber.Ctx) error {
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	roles, err := roleHandler.GetAllRoles(db)
	if err != nil {
		return err
	}
	return c.JSON(roles)
}
