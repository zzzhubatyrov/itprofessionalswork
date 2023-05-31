package main

import (
	"fmt"
	"ipw-app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true, // Very important while using an HTTP-only Cookie, frontend can easily get and return back the cookie.
	}))

	routes.Setup(app)

	fmt.Println("Server started on port :5000")
	err := app.Listen(":5000")
	if err != nil {
		return
	}
}
