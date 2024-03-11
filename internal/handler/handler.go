package handler

import (
	"github.com/gofiber/contrib/websocket"
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
	dataV1.Put("/user/update", h.updateUser)
	dataV1.Put("/user/upload-photo", h.uploadPhoto)
	dataV1.Get("/users", h.getAllUsers)
	//dataV1.Get("/user/resume")
	dataV1.Post("/user/create-resume", h.createResume)
	dataV1.Put("/user/update/resume/:id", h.updateResume)
	dataV1.Post("/user/create-company", h.createCompany)
	dataV1.Get("/create-response/:id", h.createResponse)

	resume := app.Group("/resume")
	resumeV1 := resume.Group("/v1")
	resumeV1.Get("/resumes", h.getAllResumes)
	resumeV1.Get("/:id", h.getResumeByID)
	resumeV1.Post("/delete/:id", h.deleteResume)

	vacancy := app.Group("/vacancy")
	vacancyV1 := vacancy.Group("/v1")
	vacancyV1.Post("/create-vacancy", h.createVacancy)
	vacancyV1.Get("/all-vacancies", h.getAllVacancy)
	vacancyV1.Get("/:id", h.getVacancyByID)

	company := app.Group("/company")
	companyV1 := company.Group("/v1")
	companyV1.Get("/companies", h.getAllCompanies)
	companyV1.Get("/:id", h.getCompanyByID)
	companyV1.Put("update", h.updateCompanyData)

	//companyV1.Get("/all-vacancies")

	//admin := app.Group("/admin")
	//adminV1 := admin.Group("/v1")
	//adminV1.Put("/update-role/:id", h.updateRoleByID)

	role := app.Group("/role")
	roleV1 := role.Group("/v1")
	roleV1.Get("/roles", h.getAllRoles)
	roleV1.Put("/update/:id", h.updateRoleByUserID)

	search := app.Group("/search")
	searchV1 := search.Group("/v1")
	searchV1.Get("/search/user/:tag", h.searchUser)
	searchV1.Get("/data", h.esData)

	app.Get("/message", websocket.New(h.sendMessage))
	app.Post("/send-notify", h.sendNotificationHandler)
	app.Get("/read-notifications", h.readNotificationsHandler)
	//app.Get("/read-notifications", websocket.New(func(conn *websocket.Conn) {
	//	rabbitConn, err := service.ConnectRabbitMQ()
	//	if err != nil {
	//		log.Println("Failed to connect to RabbitMQ:", err)
	//		return
	//	}
	//	defer rabbitConn.Close()
	//	msgs, err := service.ConsumeMessages(rabbitConn)
	//	if err != nil {
	//		log.Println("Failed to consume messages", err)
	//		return
	//	}
	//	for msg := range msgs {
	//		notification := string(msg.Body)
	//		// Отправка уведомления через WebSocket
	//		if err := conn.WriteJSON(notification); err != nil {
	//			log.Println("Error writing JSON to WebSocket:", err)
	//			return
	//		}
	//	}
	//}))
	//app.Get("/ws", websocket.New(h.WebSocketHandler))

	return nil
}
