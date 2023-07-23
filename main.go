package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"ipw-clean-arch/internal/handler"
	"ipw-clean-arch/internal/repository"
	"ipw-clean-arch/internal/service"
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

	//migrator := db.Migrator()
	//migrator.DropTable(&model.User{})
	//db.AutoMigrate(&model.User{})

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	app := fiber.New()
	app.Use(cors.New())
	handlers.InitRoute(app)
	app.Listen(":5000")
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("internal/config")
	return viper.ReadInConfig()
}
