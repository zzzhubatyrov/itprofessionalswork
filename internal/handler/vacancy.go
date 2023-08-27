package handler

import (
	"github.com/gofiber/fiber/v2"
	"ipw-clean-arch/internal/model"
)

func (h *Handler) createVacancy(c *fiber.Ctx) error {
	var data model.Vacancy
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	createVacancy, err := h.services.CreateVacancy(data)
	if err != nil {
		return err
	}
	return c.JSON(createVacancy)
}

func (h *Handler) getAllVacancy(c *fiber.Ctx) error {
	var data []model.Vacancy
	getAllVacancy, err := h.services.GetAllVacancy(data)
	if err != nil {
		return nil
	}
	return c.JSON(getAllVacancy)
}

func (h *Handler) getVacancyByID(c *fiber.Ctx) error {
	id := c.Params("id")
	getVacancyByID, err := h.services.GetVacancyByID(id)
	if err != nil {
		return nil
	}
	return c.JSON(getVacancyByID)
}
