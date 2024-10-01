package broker

import (
	"context"

	"github.com/Businge931/Kafka_and_CLIs/models"
)

// Producer is a generic struct that implements the MessageProducer interface.
// It can work with any messaging broker (e.g., Kafka, RabbitMQ).
type Producer struct {
	produceMessageFunc func(ctx context.Context, message models.Message) error
}

// NewProducer creates a new generic Producer instance.
// The broker-specific message production function is passed as a parameter.
func NewProducer(produceFunc func(ctx context.Context, message models.Message) error) *Producer {
	return &Producer{
		produceMessageFunc: produceFunc,
	}
}

// ProduceMessage produces a message using the generic production function.
func (p *Producer) ProduceMessage(ctx context.Context, message models.Message) error {
	return p.produceMessageFunc(ctx, message)
}

// Consumer is a generic struct that implements the MessageConsumer interface.
// It can work with any messaging broker (e.g., Kafka, RabbitMQ).
type Consumer struct {
	consumeMessagesFunc func(ctx context.Context) ([]models.Message, error)
}

// NewConsumer creates a new generic Consumer instance.
// The broker-specific message consumption function is passed as a parameter.
func NewConsumer(consumeFunc func(ctx context.Context) ([]models.Message, error)) *Consumer {
	return &Consumer{
		consumeMessagesFunc: consumeFunc,
	}
}

// ConsumeMessages consumes messages using the generic consumption function.
func (c *Consumer) ConsumeMessages(ctx context.Context) ([]models.Message, error) {
	return c.consumeMessagesFunc(ctx)
}
