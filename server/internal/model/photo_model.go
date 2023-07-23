package model

import (
	"gorm.io/gorm"
)

// Photo Модель для фото
type Photo struct {
	ID   uint   `gorm:"primaryKey"`
	Data []byte `json:"data"`
}

// SavePhoto Функция для сохранения фото
func SavePhoto(data []byte, db *gorm.DB) error {
	// Миграция модели
	err := db.AutoMigrate(&Photo{})
	if err != nil {
		return err
	}
	// Создание экземпляра модели фото
	photo := &Photo{
		Data: data,
	}
	// Сохранение фото в базе данных
	err = db.Create(photo).Error
	if err != nil {
		return err
	}
	return nil
}
