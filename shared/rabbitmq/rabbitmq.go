package rabbitmq

import (
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	logger  *zap.Logger
}

func New(url string, logger *zap.Logger) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		conn:    conn,
		channel: ch,
		logger:  logger,
	}, nil
}

func (r *RabbitMQ) PublishMessage(exchange, routingKey string, body []byte) error {
	err := r.channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	r.logger.Info("Message published", zap.String("exchange", exchange), zap.String("routingKey", routingKey))
	return nil
}

func (r *RabbitMQ) ConsumeMessages(queue string, handler func([]byte) error) error {
	msgs, err := r.channel.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			if err := handler(d.Body); err != nil {
				r.logger.Error("Error processing message", zap.Error(err))
			}
		}
	}()

	r.logger.Info("Started consuming messages", zap.String("queue", queue))
	return nil
}

func (r *RabbitMQ) Close() {
	r.channel.Close()
	r.conn.Close()
}
