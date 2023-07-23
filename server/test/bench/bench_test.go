package bench

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ipw-app/internal/interfaces"
	"ipw-app/internal/model"
	"testing"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../internal/config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ошибка чтения файла конфигурации: %s", err))
	}
}

// BENCHMARK TESTS
func BenchFunc(b *testing.B) {
	data := map[string]string{
		"email":    "worker@gmail.com",
		"password": "test123",
	}
	db, _ := Connect()
	secretKey := viper.GetString("SecretKey")
	c := new(fiber.Ctx)
	// Вызываем функцию для сброса таймера перед измерением времени выполнения.
	// Это исключает время выполнения кода до тестовой функции из результатов бенчмарка.
	b.ResetTimer()
	// Вызываем функцию для вывода информации о распределении памяти перед измерением времени выполнения.
	// Это позволяет отслеживать аллокации памяти в тестовой функции.
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var usr interfaces.UserHandler = &model.User{}
		usr.Login(data, db, secretKey, c)
	}
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
