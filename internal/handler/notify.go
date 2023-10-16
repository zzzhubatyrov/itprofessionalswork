package handler

import (
	"github.com/gofiber/fiber/v2"
	"ipw-clean-arch/internal/service"
	"log"
)

func (h *Handler) sendNotificationHandler(c *fiber.Ctx) error {
	message := c.Query("message")

	conn, err := service.ConnectRabbitMQ()
	if err != nil {
		log.Println("Failed to connect to RabbitMQ:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to send notify")
	}
	defer conn.Close()

	err = service.SendNotification(conn, message)
	if err != nil {
		log.Println("Failed to send notification:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to send notification")
	}

	return c.SendString("Notification sent:" + message)
}

func (h *Handler) readNotificationsHandler(c *fiber.Ctx) error {
	conn, err := service.ConnectRabbitMQ()
	if err != nil {
		log.Println("Failed to connect to RabbitMQ:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to read notify")
	}
	defer conn.Close()

	msgs, err := service.ConsumeMessages(conn)
	if err != nil {
		log.Println("Failed to consume messages", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to read notify")
	}

	notifications := make([]string, 0)
	for msg := range msgs {
		notifications = append(notifications, string(msg.Body))
	}

	return c.JSON(notifications)
}
