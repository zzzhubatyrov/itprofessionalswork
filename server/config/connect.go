package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=tigserf dbname=ipwtest port=5432 sslmode=disable TimeZone=Asia/Yekaterinburg"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
