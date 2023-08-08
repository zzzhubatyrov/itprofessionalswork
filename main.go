package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"ipw-clean-arch/internal/handler"
	"ipw-clean-arch/internal/model"
	"ipw-clean-arch/internal/repository"
	"ipw-clean-arch/internal/service"
	_ "ipw-clean-arch/internal/service"
)

func main() {
	if err := initConfig(); err != nil {
		fmt.Errorf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		fmt.Errorf("failed to initialize db: %s", err.Error())
	}

	models := []interface{}{
		&model.User{},
		&model.Role{},
		&model.Resume{},
	}
	//migrator := db.Migrator()
	//_ = migrator.DropTable(models...)
	db.AutoMigrate(models...)
	//db.Create(&model.Role{Name: "Администратор"})
	//db.Create(&model.Role{Name: "Модератор"})
	//db.Create(&model.Role{Name: "HR"})
	//db.Create(&model.Role{Name: "Пользователь"})

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))
	handlers.InitRoute(app)
	app.Listen(":5000")
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("internal/config")
	return viper.ReadInConfig()
}