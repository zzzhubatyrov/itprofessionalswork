package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnect interface {
	Connect() (*gorm.DB, error)
}

type GormConnect struct {
	db *gorm.DB
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".\\config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ошибка чтения файла конфигурации: %s", err))
	}
}

func (conn *GormConnect) Connect() (*gorm.DB, error) {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")     // Изменено на GetString
	dbname := viper.GetString("database.dbname") // Изменено на GetString
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	sslMode := viper.GetString("database.sslMode")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, username, password, dbname, port, sslMode)
	var err error
	conn.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to databse: %s", err.Error())
	}
	return conn.db, nil
}
