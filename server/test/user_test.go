package test_test

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ipw-app/internal/interfaces/mocks"
	"ipw-app/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../internal/config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ошибка чтения файла конфигурации: %s", err))
	}
}

func TestUserHandlerMock_Login(t *testing.T) {
	mockUserHandler := &mocks.UserHandler{}
	data := map[string]string{
		"email":    "worker@gmail.com",
		"password": "test123",
	}
	//user := model.User{
	//	Email:    data["email"],
	//	Password: data["password"],
	//}
	db, _ := Connect()
	secretKey := viper.GetString("SecretKey")
	c := new(fiber.Ctx)
	expectedEmail := data["email"]
	expectedPassword := data["password"]
	mockUserHandler.On("Login", data, db, secretKey, c).Return(nil).Once()
	returnedErr := mockUserHandler.Login(data, db, secretKey, c)
	assert.NoError(t, returnedErr)
	mockUserHandler.AssertCalled(t, "Login", data, db, secretKey, c)
	mockUserHandler.AssertExpectations(t)
	assert.Equal(t, expectedEmail, data["email"])       // Проверяем, что значение data["email"] не изменилось
	assert.Equal(t, expectedPassword, data["password"]) // Проверяем, что значение data["password"] не изменилось
}

func TestUserHandlerMock_Register(t *testing.T) {
	mockUserHandler := &mocks.UserHandler{}
	data := map[string]string{
		"email": "john@doe.ru",
	}
	user := &model.User{
		Email: data["email"],
	}
	db, _ := Connect()
	mockUserHandler.On("Register", data, db).Return(user, nil).Once()
	returnedUser, err := mockUserHandler.Register(data, db)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, returnedUser.Email)
	mockUserHandler.AssertCalled(t, "Register", data, db)
	mockUserHandler.AssertExpectations(t)
	// Дополнительная проверка на правильность введенных значений
	assert.Equal(t, "john@doe.ru", data["email"]) // Проверяем, что значение в data не изменилось
	assert.Equal(t, "john@doe.ru", user.Email)    // Проверяем, что значение в user правильно установлено
}

func Connect() (*gorm.DB, error) {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")     // Изменено на GetString
	dbname := viper.GetString("database.dbname") // Изменено на GetString
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	sslMode := viper.GetString("database.sslMode")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, username, password, dbname, port, sslMode)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
