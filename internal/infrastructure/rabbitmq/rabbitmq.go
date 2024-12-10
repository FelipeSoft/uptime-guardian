package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
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
	return &RabbitMQ{conn: conn, ch: ch}, nil
}

func (r *RabbitMQ) Publish(queueName string, body []byte) error {
	return r.ch.Publish(
		"",
		queueName,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		})
}

func (r *RabbitMQ) DeclareQueue(queueName string, args amqp.Table) (amqp.Queue, error) {
	queue, err := r.ch.QueueDeclare(
		queueName,
		true, 
		false,
		false,
		false,
		args,
	)
	return queue, err
}

func (r *RabbitMQ) Consume(queueName string, consumerName string) (<-chan amqp.Delivery, error) {
	msgs, err := r.ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (r *RabbitMQ) Close() {
	r.conn.Close()
	r.ch.Close()
}
