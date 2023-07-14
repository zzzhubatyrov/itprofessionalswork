package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ipw-app/internal/model"
)

// UserHandler
// TODO add Update method, CheckEmail, CheckUser, VerifyEmail
type UserHandler interface {
	User(db *gorm.DB, secretKey string, c *fiber.Ctx) error
	Register(data map[string]string, db *gorm.DB) (*model.User, error)
	Login(data map[string]string, db *gorm.DB, secretKey string, c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}
