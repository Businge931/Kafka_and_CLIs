package service

import (
	"context"

	"github.com/Businge931/Kafka_and_CLIs/models"
)

type (
	MessageProducer interface {
		ProduceMessage(
			ctx context.Context, 
			message models.Message) error
	}
	MessageConsumer interface {
		ConsumeMessages(ctx context.Context) ([]models.Message, error)
	}

	Service struct {
		producer MessageProducer
		consumer MessageConsumer
	}
)

func New(producer MessageProducer, consumer MessageConsumer) *Service {
	return &Service{
		producer: producer,
		consumer: consumer,
	}
}
