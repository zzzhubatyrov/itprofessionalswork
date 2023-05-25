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
		AllowCredentials: true, // Very important while using a HTTPonly Cookie, frontend can easily get and return back the cookie.
	}))

	routes.Setup(app)

	fmt.Println("Server started on port :5000")
	app.Listen(":5000")
}
