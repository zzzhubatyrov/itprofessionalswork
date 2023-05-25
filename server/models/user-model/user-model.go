package usermodel

import (
	"ipw-app/config"
	"log"
	"time"
)

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Surname    string    `json:"surname" gorm:"type:varchar(255);not null"`
	Name       string    `json:"name" gorm:"type:varchar(255);not null"`
	Lastname   string    `json:"lastname" gorm:"type:varchar(255)"`
	Email      string    `json:"email" gorm:"type:varchar(255);unique;not null;column:email"`
	Password   []byte    `json:"-" gorm:"type:varchar(255);not null;column:password"`
	Created_at time.Time `json:"created_at"`
	// Updated_At time.Time `json:"updated_at" gorm:"autoUpdateTime:false"`
	// Deleted_At time.Time `json:"deleted_at" gorm:"autoDeleteTime:false"`
}

func init() {
	db, err := config.Connection()
	if err != nil {
		log.Println(err)
	}

	migrator := db.Migrator()
	if !migrator.HasTable(User{}) {
		if err := db.AutoMigrate(&User{}); err != nil {
			log.Println(err)
		}
	} else {
		if err := migrator.DropTable(&User{}); err != nil {
			log.Println(err)
		}
		if err := db.AutoMigrate(&User{}); err != nil {
			log.Println(err)
		}
	}
}
