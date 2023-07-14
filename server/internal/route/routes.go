package route

import (
	// "ipw-app/controllers"

	"github.com/gofiber/fiber/v2"
	"ipw-app/internal/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/get-user", controllers.User)
	app.Get("/get-role/:@tag", controllers.GetRole)

	app.Post("/photos", controllers.UploadPhoto)
	//app.Post("/change-role", controllers.ChangeRole)

	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)
}
