package model

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type User struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	Email       string     `json:"email" gorm:"type:varchar(255);unique;not null;column:email"`
	Password    []byte     `json:"-" gorm:"type:varchar(255);not null;column:password"`
	Name        string     `json:"name" gorm:"type:varchar(255);column:name"`
	Age         string     `json:"age" gorm:"type:varchar(255);column:age"`
	Tag         string     `json:"tag" gorm:"type:varchar(255);unique;column:tag"`
	Direction   string     `json:"direction" gorm:"type:varchar(255);column:direction"`
	Status      string     `json:"status" gorm:"type:varchar(255);column:status"`
	Level       string     `json:"level" gorm:"type:varchar(255);column:level"`
	Salary      string     `json:"salary" gorm:"type:varchar(255);column:salary"`
	Skills      string     `json:"skills" gorm:"type:varchar(255);column:skills"`
	Description string     `json:"description" gorm:"type:varchar(255);column:description"`
	Number      string     `json:"number" gorm:"type:varchar(255);column:number"` //unique;
	Gender      string     `json:"gender" gorm:"type:varchar(255);column:gender"`
	Birthday    *time.Time `json:"birthday" gorm:"type:varchar(255);column:birthday"`
	Location    string     `json:"location" gorm:"type:varchar(255);column:location"`
	Role        []Role     `json:"role" gorm:"not null;column:role;gorm:foreignKey:Name"`
	Photo       []byte     `json:"photo" gorm:"gorm:column:photo"`
	Resume      []*Resume  `json:"resume" gorm:"many2many:resumes"`
}

//var (
//	connect repository.DBConnect = new(repository.GormConnect)
//)
//
//func init() {
//	db, err := connect.Connect()
//	if err != nil {
//		log.Println(err)
//	}
//
//	migrator := db.Migrator()
//	if !migrator.HasTable(User{}) {
//		if err := db.AutoMigrate(&User{}); err != nil {
//			log.Println(err)
//		}
//	} else {
//		if err := migrator.DropTable(&User{}); err != nil {
//			log.Println(err)
//		}
//		if err := db.AutoMigrate(&User{}); err != nil {
//			log.Println(err)
//		}
//	}
//}

func (user *User) User(db *gorm.DB, secretKey string, c *fiber.Ctx) error {
	cookie := c.Cookies("ipw_cookie")
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil // using the SecretKey which was generated in th Login function
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

func (user *User) Register(data map[string]string, db *gorm.DB) (*User, error) {
	// Проверка наличия уже зарегистрированного пользователя с указанным email
	var existingUser User
	result := db.Where("email = ?", data["email"]).First(&existingUser)
	if result.Error == nil {
		return nil, fmt.Errorf("user already exists")
	} else if result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	var role Role
	db.Where("name = ?", data["role"]).First(&role)
	if db.Error != nil {
		return nil, db.Error
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 15)
	if err != nil {
		return nil, err
	}
	regUser := &User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
		Tag:      data["tag"],
		//Role:      &role,
		Direction: data["direction"],
		Level:     data["level"],
		Salary:    data["salary"],
		Status:    data["status"],
		Skills:    data["skills"],
		Age:       data["age"],
		Location:  data["location"],
		Resume: []*Resume{
			{
				Name: "OneRes",
			},
		},
	}
	db.Create(regUser)
	return regUser, nil
}

func (user *User) Login(data map[string]string, db *gorm.DB, secretKey string, c *fiber.Ctx) error {
	db.Where("email = ?", data["email"]).First(&user)
	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
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

func (user *User) Logout(c *fiber.Ctx) error {
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
func (user *User) UserUploadPhoto(c *fiber.Ctx, db *gorm.DB) error {
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
