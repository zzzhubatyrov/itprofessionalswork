package handler

import (
	"github.com/gofiber/fiber/v2"
	"ipw-clean-arch/internal/model"
)

func (h *Handler) createUser(c *fiber.Ctx) error {
	var input model.User
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	regUser, err := h.services.Register(input)
	if err != nil {
		return err
	}
	return c.JSON(regUser)
}

func (h *Handler) loginUser(c *fiber.Ctx) error {
	var input model.User
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	loginUser, err := h.services.Login(input, secretKey, c)
	if err != nil {
		return err
	}
	return c.JSON(loginUser)
}

func (h *Handler) logoutUser(c *fiber.Ctx) error {
	return h.services.Logout(c)
}
