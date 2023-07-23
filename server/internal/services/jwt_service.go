package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// VerifyJWT Вспомогательная функция для проверки и верификации JWT токена из кук
func VerifyJWT(secretKey string, cookie string, c *fiber.Ctx) (*jwt.Token, error) {
	// &jwt.MapClaims
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil // используем SecretKey, который был сгенерирован в функции Login
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	if !token.Valid {
		return nil, errors.New("недействительный JWT токен")
	}
	return token, nil
}
