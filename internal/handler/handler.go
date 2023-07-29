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

	return nil
}
