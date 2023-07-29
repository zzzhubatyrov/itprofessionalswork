package handler

import (
	"github.com/gofiber/fiber/v2"
	"ipw-clean-arch/internal/model"
)

func (h *Handler) getUserData(c *fiber.Ctx) error {
	var data model.User
	getUserData, err := h.services.GetUser(data, secretKey, c)
	if err != nil {
		return err
	}
	return c.JSON(getUserData)
}

func (h *Handler) getAllUsers(c *fiber.Ctx) error {
	var users []model.User
	getAllUsers, err := h.services.GetAllUsers(users)
	if err != nil {
		return err
	}
	return c.JSON(getAllUsers)
}
