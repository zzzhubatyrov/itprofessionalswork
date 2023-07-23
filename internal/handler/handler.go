package handler

import (
	"github.com/gofiber/fiber/v2"
	"ipw-clean-arch/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoute(app *fiber.App) fiber.Handler {
	auth := app.Group("/auth")
	auth.Post("/create-user", h.createUser)

	data := app.Group("/data")
	data.Get("/users", h.getAllUser)
	return nil
}
