package handler

import (
	"github.com/gofiber/fiber/v2"
	"ipw-clean-arch/internal/model"
)

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
