package route

import (
	// "ipw-app/controllers"

	"github.com/gofiber/fiber/v2"
	"ipw-app/internal/controllers"
)

func Setup(app *fiber.App) {
	// USER HANDLER
	app.Get("/get-user", controllers.User)
	app.Get("/get-user/:id", controllers.GetUserByID)
	app.Post("/photos", controllers.UploadPhoto)
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)

	// COMPANY HANDLER
	app.Post("/register-company", controllers.CreateCompany)
	app.Get("/company/vacancy", controllers.GetVacancy)
	//app.Get("/company/vacancy/:id")

	// VACANCY HANDLER
	app.Post("/create-vacancy", controllers.CreateVacancy)
	app.Get("/vacancy", controllers.GetVacancy)

	// ROLE HANDLER
	//app.Get("/get-role", controllers.GetRole)
	app.Get("/get-roles", controllers.GetAllRoles)
	app.Post("/change-role", controllers.ChangeRole)

	//app.Get("/get-role/:@tag", controllers.GetRoleByTag)
}
