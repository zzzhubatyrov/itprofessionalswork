package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"ipw-clean-arch/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

var secretKey = viper.GetString("SecretKey")

func (h *Handler) InitRoute(app *fiber.App) fiber.Handler {
	auth := app.Group("/auth")
	authV1 := auth.Group("/v1")
	authV1.Post("/register", h.createUser)
	authV1.Post("/login", h.loginUser)
	authV1.Post("/logout", h.logoutUser)

	data := app.Group("/data")
	dataV1 := data.Group("/v1")
	dataV1.Get("/user", h.getUserData)
	dataV1.Get("/users", h.getAllUsers)
	//dataV1.Get("/user/resume")
	dataV1.Post("/user/create-resume", h.createResume)
	dataV1.Put("/user/update/resume/:id", h.updateResume)
	dataV1.Post("/user/create-company", h.createCompany)

	resume := app.Group("/resume")
	resumeV1 := resume.Group("/v1")
	resumeV1.Get("/resumes", h.getAllResumes)

	vacancy := app.Group("/vacancy")
	vacancyV1 := vacancy.Group("/v1")
	vacancyV1.Post("/create-vacancy", h.createVacancy)
	vacancyV1.Get("/all-vacancies", h.getAllVacancy)
	vacancyV1.Get("/:id", h.getVacancyByID)

	//company := app.Group("/company")
	//companyV1 := vacancy.Group("/v1")
	//companyV1.Get("/all-vacancies")

	// Add middleware on check user role
	admin := app.Group("/admin-panel")
	adminV1 := admin.Group("/v1")
	//adminV1.Get("/roles")
	role := adminV1.Group("/role")
	roleV1 := role.Group("/v1")
	roleV1.Get("/roles", h.getAllRoles)

	return nil
}
