package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

// ParseToken разбирает токен и возвращает объект типа *jwt.Token или ошибку.
func ParseToken(tokenString, secretKey string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// GetToken возвращает значение поля из токена или ошибку, если поле не найдено.
func GetToken(token *jwt.Token) (*jwt.RegisteredClaims, error) {
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, errors.New("неверный формат токена")
	}
	return claims, nil
}
