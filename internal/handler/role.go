package handler

import (
	"github.com/gofiber/fiber/v2"
	"ipw-clean-arch/internal/model"
)

func (h *Handler) getAllRoles(c *fiber.Ctx) error {
	var data []model.Role
	getAllRoles, err := h.services.GetAllRoles(data)
	if err != nil {
		return err
	}
	return c.JSON(getAllRoles)
}

func (h *Handler) updateRoleByUserID(c *fiber.Ctx) error {
	id := c.Params("id")
	updateRoleByUserID := h.services.UpdateRoleByUserID(id, 1)
	return c.JSON(updateRoleByUserID)
}
