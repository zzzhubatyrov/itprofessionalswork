package handler

import (
	"github.com/gofiber/fiber/v2"
	"ipw-clean-arch/internal/model"
	"strconv"
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

func (h *Handler) getResumeByID(c *fiber.Ctx) error {
	id := c.Params("id")
	usrID, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	getResumeByID, err := h.services.GetResumeByID(usrID)
	if err != nil {
		return err
	}
	return c.JSON(getResumeByID)
}

func (h *Handler) getAllResumes(c *fiber.Ctx) error {
	var data []model.Resume
	resumes, err := h.services.GetAllResumes(data)
	if err != nil {
		return err
	}
	return c.JSON(resumes)
}

func (h *Handler) deleteResume(c *fiber.Ctx) error {
	id := c.Params("id")
	delResume := h.services.DeleteResume(id)
	return c.JSON(delResume)
}
