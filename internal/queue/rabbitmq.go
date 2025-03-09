package queue

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare("tasks", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	// Set QoS (Prefetch Count = 1) for fair distribution
	err = ch.Qos(1, 0, false)
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{conn, ch}, nil
}

func (q *RabbitMQ) Publish(queue, message string) error {
	err := q.ch.Publish("", queue, false, false,
		amqp.Publishing{ContentType: "text/plain", Body: []byte(message)})
	return err
}

func (q *RabbitMQ) Consume(queue string) (<-chan amqp.Delivery, error) {
	err := q.ch.Qos(1, 0, false) // Ensure each worker gets only one unacknowledged task at a time
	if err != nil {
		return nil, err
	}
	return q.ch.Consume(queue, "", false, false, false, false, nil)
}

func (q *RabbitMQ) Close() {
	q.ch.Close()
	q.conn.Close()
}
