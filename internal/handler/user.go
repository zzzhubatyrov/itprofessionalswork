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

func (h *Handler) updateUser(c *fiber.Ctx) error {
	var data model.User
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	updateUser, err := h.services.UpdateUser(data, secretKey, c)
	if err != nil {
		return err
	}
	return c.JSON(updateUser)
}

func (h *Handler) getAllUsers(c *fiber.Ctx) error {
	var users []model.User
	getAllUsers, err := h.services.GetAllUsers(users)
	if err != nil {
		return err
	}
	return c.JSON(getAllUsers)
}

func (h *Handler) createResponse(c *fiber.Ctx) error {
	id := c.Params("id")
	getVacancy, err := h.services.GetVacancyByID(id)
	createResponse, err := h.services.CreateResponse(*getVacancy, secretKey, c)
	if err != nil {
		return err
	}
	return c.JSON(createResponse)
}
