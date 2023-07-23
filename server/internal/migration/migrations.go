package migration

import (
	"gorm.io/gorm"
	"ipw-app/internal/interfaces"
	"ipw-app/internal/model"
	"ipw-app/internal/model/role"
	"ipw-app/internal/repository"
	"log"
)

var connect repository.DBConnect = new(repository.GormConnect)

func init() {
	db, err := connect.Connect()
	if err != nil {
		log.Println(err)
	}
	data := map[string]string{
		"name":     "Rodion Zhubatyrov",
		"age":      "17",
		"email":    "ipw@bk.ru",
		"password": "test123",
		"tag":      "@zhubatyrov",
	}
	data2 := map[string]string{
		"name":     "Donkey Lover",
		"age":      "19",
		"email":    "worker@gmail.com",
		"password": "test123",
		"tag":      "@donkeyhot",
	}
	dropTable(db)
	migrations(db)
	createItems(db)
	autoRegUsers(data, db)
	autoRegUsers(data2, db)
}

func createItems(db *gorm.DB) {
	db.Create(&model.Role{Name: "Администратор"})
	db.Create(&model.Role{Name: "Модератор"})
	db.Create(&model.Role{Name: "HR"})
	db.Create(&model.Role{Name: "Пользователь"})
}

func migrations(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Vacancy{},
		&model.Company{},

		&role.Admin{},
		&role.HR{},
		&role.Moderator{},
	)
}

func dropTable(db *gorm.DB) {
	migrator := db.Migrator()
	migrator.DropTable(
		&model.User{},
		&model.Role{},
		&model.Vacancy{},
		&model.Company{},

		&role.Admin{},
		&role.HR{},
		&role.Moderator{},
	)
}

func autoRegUsers(data map[string]string, db *gorm.DB) {
	var userHandler interfaces.UserHandler = &model.User{}
	userHandler.Register(data, db)

	// "name": "Rodion Zhubatyrov",
	// "age": "17",
	// "email": "ipw@bk.ru",
	// "password": "test123",
	// "tag": "@zhubatyrov"
}
