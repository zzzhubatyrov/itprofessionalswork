package controllers

import (
	"ipw-app/config"
	usermodel "ipw-app/models/user-model"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	db, dbErr := config.Connection()
	if dbErr != nil {
		log.Println(dbErr)
	}
	var data usermodel.User
	if dataErr := c.BodyParser(&data); dataErr != nil {
		return dataErr
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 15)

	user := usermodel.User{
		Surname:  data.Surname,
		Name:     data.Name,
		Lastname: data.Lastname,
		Email:    data.Email,
		Password: password,
	}

	db.Create(&user)

	return c.JSON(user)
}

const SecretKey = "02639fa6273c7e789b3ed709a9210dc05f4188037bbdf3ebc87b10199d405061"

func Login(c *fiber.Ctx) error {
	db, _ := config.Connection()

	var data usermodel.User
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user usermodel.User
	db.Where("email = ?", data.Email).First(&user) // Check the email is present in the DB

	if user.ID == 0 { // If the ID return is '0' then there is no such email present in the DB
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	} // If the email is present in the DB then compare the Passwords and if incorrect password then return error.

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),            // issuer contains the ID of the user.
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Adds time to the token i.e. 24 hours.
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	} // Creates the cookie to be passed.

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    cookie,
	})
}

func User(c *fiber.Ctx) error {
	db, _ := config.Connection()
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil // using the SecretKey which was generated in th Login function
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user usermodel.User
	db.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Sets the expiry time an hour ago in the past.
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    cookie,
	})
}
