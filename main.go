package main

import (
	"fmt"
	"ipw-clean-arch/internal/handler"
	"ipw-clean-arch/internal/model"
	"ipw-clean-arch/internal/repository"
	"ipw-clean-arch/internal/service"
	"time"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		_ = fmt.Errorf("error initializing configs: %s", err.Error())
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
		_ = fmt.Errorf("failed to initialize db: %s", err.Error())
	}

	client := repository.NewRedisDB()

	models := []interface{}{
		&model.User{},
		&model.Company{},
		&model.Vacancy{},
		//&model.Response{},
		&model.Resume{},
		//&model.Role{},
	}
	//migrator := db.Migrator()
	//_ = migrator.DropTable(models...)
	_ = db.AutoMigrate(models...)
	//db.Create(&model.Role{Name: "Администратор"})
	//db.Create(&model.Role{Name: "Модератор"})
	//db.Create(&model.Role{Name: "HR"})
	//db.Create(&model.Role{Name: "Пользователь"})

	repos := repository.NewRepository(db, client)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, https://itprofessionalswork.ru",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.yaml",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}
	app.Use(swagger.New(cfg))
	app.Use(func(c *fiber.Ctx) error {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		message := fmt.Sprintf("[IPW-Log][%s] port: %s method: %s - %s", currentTime, c.IP(), c.Method(), c.Path())
		fmt.Println(message)
		return c.Next()
	})
	handlers.InitRoute(app)
	//messenger.InitializeMessenger(app)
	app.Listen(":5000")
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("internal/config")
	return viper.ReadInConfig()
}
