package handler

import (
	"ipw-clean-arch/internal/model"

	"github.com/gofiber/fiber/v2"
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

func (h *Handler) updateCompanyData(c *fiber.Ctx) error {
	var company model.Company
	if err := c.BodyParser(&company); err != nil {
		return err
	}
	//updateCompanyData, err := h.services.UpdateCompanyData(data, secretKey, c)
	updateCompanyData, err := h.services.UpdateCompanyData(company, secretKey, c)
	if err != nil {
		return err
	}
	return c.JSON(updateCompanyData)
}

func (h *Handler) getAllCompanies(c *fiber.Ctx) error {
	var companies []model.Company
	getAllCompanies, err := h.services.GetAllCompanies(companies)
	if err != nil {
		return err
	}
	return c.JSON(getAllCompanies)
}

func (h *Handler) getCompanyByID(c *fiber.Ctx) error {
	id := c.Params("id")
	getCompanyByID, err := h.services.GetCompanyByID(id)
	if err != nil {
		return err
	}
	return c.JSON(getCompanyByID)
}
