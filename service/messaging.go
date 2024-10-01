package service

import (
	"context"

	"github.com/Businge931/Kafka_and_CLIs/models"
)

func (s *Service) SendMessage(
	ctx context.Context,
	 message models.Message) error {
	return s.producer.ProduceMessage(
		ctx,
		 message)
}

func (s *Service) ReceiveMessages(ctx context.Context) ([]models.Message, error) {
	return s.consumer.ConsumeMessages(ctx)
}
