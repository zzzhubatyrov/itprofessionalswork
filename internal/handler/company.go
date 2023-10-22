package handler

import (
	"github.com/gofiber/fiber/v2"
	"ipw-clean-arch/internal/model"
)

func (h *Handler) createCompany(c *fiber.Ctx) error {
	var data model.Company
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	createCompany, err := h.services.CreateCompany(data, secretKey, c)
	if err != nil {
		return err
	}
	return c.JSON(createCompany)
}

func (h *Handler) getCompanyByID(c *fiber.Ctx) error {
	id := c.Params("id")
	getCompanyByID, err := h.services.GetCompanyByID(id)
	if err != nil {
		return err
	}
	return c.JSON(getCompanyByID)
}
