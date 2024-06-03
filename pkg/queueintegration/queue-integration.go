package queue_integration

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageQueue interface {
	Connect(url string) error
	Close() error
	Publish(ctx context.Context, queueName string, msg string) error
	Consume(queueName string) (<-chan string, error)
}

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func (r *RabbitMQ) Connect(url string) error {
	var err error
	r.conn, err = amqp.Dial(url)
	if err != nil {
		return err
	}

	r.channel, err = r.conn.Channel()
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQ) Close() error {
	if err := r.channel.Close(); err != nil {
		return err
	}
	return r.conn.Close()
}

func (r *RabbitMQ) Publish(ctx context.Context, queueName string, msg string) error {
	q, err := r.channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return r.channel.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	})
}

func (r *RabbitMQ) Consume(queueName string) (<-chan string, error) {
	q, err := r.channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	msgs, err := r.channel.Consume(
		q.Name,
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

	out := make(chan string)
	go func() {
		for d := range msgs {
			out <- string(d.Body)
		}
		close(out)
	}()

	return out, nil
}
