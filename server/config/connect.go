package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".\\config")
	// Чтение файла конфигурации
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ошибка чтения файла конфигурации: %s", err))
	}
}

func Connection() (*gorm.DB, error) {
	// Получение значения параметра конфигурации
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")     // Изменено на GetString
	dbname := viper.GetString("database.dbname") // Изменено на GetString
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	sslMode := viper.GetString("database.sslMode")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, username, password, dbname, port, sslMode)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
