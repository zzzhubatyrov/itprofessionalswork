package repository

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ipw-clean-arch/internal/model"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u UserPostgres) GetUser(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u UserPostgres) GetUserByID(id string, db *gorm.DB) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserPostgres) Register(data map[string]string, db *gorm.DB) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserPostgres) Login(data map[string]string, db *gorm.DB, secretKey string, c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u UserPostgres) Logout(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u UserPostgres) CreateUser(user model.User) (model.User, error) {
	//TODO implement me
	panic("implement me")
}
