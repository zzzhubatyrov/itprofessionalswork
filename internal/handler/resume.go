package handler

import (
	"github.com/gofiber/fiber/v2"
	"ipw-clean-arch/internal/model"
)

func (h *Handler) createResume(c *fiber.Ctx) error {
	var input model.Resume
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	createResume, err := h.services.CreateResume(input, secretKey, c)
	if err != nil {
		return err
	}
	return c.JSON(createResume)
}

func (h *Handler) updateResume(c *fiber.Ctx) error {
	var input model.Resume
	id := c.Params("id")
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	updateResume, err := h.services.UpdateResume(input, id, secretKey, c)
	if err != nil {
		return err
	}
	return c.JSON(updateResume)
}

func (h *Handler) getAllResumes(c *fiber.Ctx) error {
	var data []model.Resume
	resumes, err := h.services.GetAllResumes(data)
	if err != nil {
		return err
	}
	return c.JSON(resumes)
}
