package handler

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"ipw-clean-arch/internal/service"
	"log"
)

func (h *Handler) sendNotificationHandler(c *fiber.Ctx) error {
	var body struct {
		Message string `json:"message"`
	}
	//message := c.Query("message")

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	conn, err := service.ConnectRabbitMQ()
	if err != nil {
		log.Println("Failed to connect to RabbitMQ:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to send notify")
	}
	defer conn.Close()

	err = service.SendNotification(conn, body.Message)
	if err != nil {
		log.Println("Failed to send notification:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to send notification")
	}

	return c.SendString("Notification sent:" + body.Message)
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

func (h *Handler) WebSocketHandler(c *websocket.Conn) {
	defer c.Close()
	// Горутина для чтения уведомлений из RabbitMQ и отправки по WebSocket
	go func() {
		conn, err := service.ConnectRabbitMQ()
		if err != nil {
			log.Println("Failed to connect to RabbitMQ:", err)
			return
		}
		defer conn.Close()
		msgs, err := service.ConsumeMessages(conn)
		if err != nil {
			log.Println("Failed to consume messages", err)
			return
		}
		for msg := range msgs {
			notification := string(msg.Body)
			if err := c.WriteMessage(websocket.TextMessage, []byte(notification)); err != nil {
				log.Println("Failed to send WebSocket message:", err)
				return
			}
		}
	}()
	// Ждем сообщений от клиента, например, если клиент закрыл соединение
	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			log.Println("WebSocket connection closed by the client:", err)
			return
		}
	}
}
