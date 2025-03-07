package core

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	queue   amqp091.Queue
}

func NewRabbitMQ(queueName string) (*RabbitMQ, error) {
	conn, err := amqp091.Dial("amqp://Fernando:12345@13.216.97.215:5672/")
	if err != nil {
		return nil, fmt.Errorf("error al conectar a RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("error al abrir canal en RabbitMQ: %w", err)
	}

	q, err := ch.QueueDeclare(
		queueName, // nombre de la cola
		true,      // durable: true para que la cola sea duradera
		false,     // auto-delete
		false,     // exclusive
		false,     // no-wait
		nil,       // argumentos adicionales
	)
	if err != nil {
		return nil, fmt.Errorf("error al declarar la cola: %w", err)
	}

	return &RabbitMQ{conn: conn, channel: ch, queue: q}, nil
}

func (r *RabbitMQ) PublishMessage(message string) error {
	err := r.channel.Publish(
		"",           // exchange
		r.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
			DeliveryMode: amqp091.Persistent, 
		},
	)
	if err != nil {
		return fmt.Errorf("error al publicar mensaje en RabbitMQ: %w", err)
	}

	log.Println("Mensaje enviado a la cola:", message)
	return nil
}

func (r *RabbitMQ) Close() {
	r.channel.Close()
	r.conn.Close()
}