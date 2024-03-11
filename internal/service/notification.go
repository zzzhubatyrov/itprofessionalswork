package service

import (
	"github.com/streadway/amqp"
)

func ConnectRabbitMQ() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func SendNotification(conn *amqp.Connection, message string) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"notifications",
		false,
		false,
		false,
		false,
		nil,
	)

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func ConsumeMessages(conn *amqp.Connection) (<-chan amqp.Delivery, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	msgs, err := ch.Consume(
		"notifications",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	//for msg := range msgs {
	//	fmt.Println(string(msg.Body))
	//}
	return msgs, nil
}
