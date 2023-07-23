package model

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"ipw-app/internal/services"
	"strconv"
	"time"
)

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Age      string `json:"age" gorm:"type:varchar(255)"`
	Tag      string `json:"tag" gorm:"type:varchar(255);unique"`
	RoleID   int    `json:"role_id"`
	Role     Role   `json:"role" gorm:"foreignKey:RoleID"`
}

func (user User) GetUser(db *gorm.DB, secretKey string, c *fiber.Ctx) error {
	userData := User{}
	cookie := c.Cookies("ipw_cookie")
	token, err := services.VerifyJWT(secretKey, cookie, c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	if err := db.Preload("Role").Where("id = ?", claims.Issuer).First(&userData).Error; err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed to retrieve user data: %v", err),
		})
	}
	return c.JSON(userData)
}

func (user User) Register(data map[string]string, db *gorm.DB) (*User, error) {
	var existingUser User
	result := db.Where("email = ?", data["email"]).First(&existingUser)
	if result.Error == nil {
		return nil, fmt.Errorf("user already exists")
	} else if result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 15)
	if err != nil {
		return nil, err
	}
	regUser := &User{
		Email:    data["email"],
		Password: string(password),
		Name:     data["name"],
		Age:      data["age"],
		Tag:      data["tag"],
		RoleID:   4,
	}
	var role Role
	roleResult := db.First(&role, regUser.RoleID)
	if roleResult.Error != nil {
		return nil, fmt.Errorf("failed to find role: %v", roleResult.Error)
	}
	role.UserCount++
	if err := db.Save(&role).Error; err != nil {
		return nil, err
	}
	if err := db.Create(regUser).Error; err != nil {
		return nil, err
	}
	return regUser, nil
}

func (user User) Login(data map[string]string, db *gorm.DB, secretKey string, c *fiber.Ctx) error {
	db.Where("email = ?", data["email"]).First(&user)
	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
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
		"message":     "success",
		"cookieName":  cookie.Name,
		"cookieValue": cookie.Value,
	})
}

func (user User) GetUserByID(id string, db *gorm.DB) (*User, error) {
	var getUser User
	if err := db.Preload("Role").Where("id = ?", id).First(&getUser).Error; err != nil {
		return nil, err
	}
	return &getUser, nil
}

func (user User) Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "ipw_cookie",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Sets the expiry time an hour ago in the past.
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success logout",
	})
}

// #TODO WARNING Delete this method, only test case
func (user User) ChangeRole(db *gorm.DB) (*Role, error) {
	// #TODO implements me!
	panic("implements me!")
}

type UploadHandler interface {
	UserUploadPhoto(c *fiber.Ctx, db *gorm.DB) error
}

// UserUploadPhoto
//
// # TODO Warning this test func for UploadPhoto
//
// # TODO FIXME
//
// Метод для загрузки фото для пользователя
// Обработчик для загрузки фото
func (user User) UserUploadPhoto(c *fiber.Ctx, db *gorm.DB) error {
	// Получите файл из запроса
	file, err := c.FormFile("photo")
	if err != nil {
		return err
	}
	// Откройте файл
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func() { _ = src.Close() }()
	// Создайте пустой буфер для чтения файла
	buf := new(bytes.Buffer)
	buf.ReadFrom(src)
	// Получите срез байтов файла
	fileBytes := buf.Bytes()
	// Сохраните файл в базу данных или файловой системе
	err = SavePhoto(fileBytes, db)
	if err != nil {
		return err
	}
	return c.SendString("Фото успешно загружено")
}
