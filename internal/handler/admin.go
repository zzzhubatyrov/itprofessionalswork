package handler

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h *Handler) updateRoleByID(c *fiber.Ctx) error {
	id := c.Params("id")
	usrId, _ := strconv.Atoi(id)
	updateUserRoleById, err := h.services.UpdateUserRoleByID(usrId)
	if err != nil {
		return err
	}
	return c.JSON(updateUserRoleById)
}
