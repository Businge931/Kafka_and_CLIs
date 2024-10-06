package service

// "github.com/Businge931/Kafka_and_CLIs/models"

func (s *Service) SendMessage(topic, message string) error {
	return s.producer.SendMessage(topic, message)
}

func (s *Service) ReceiveMessages(topic string, 
	// testing bool
	) error {
	return s.consumer.ReadMessages(topic)
}

// func (s *Service) SendMessage(ctx context.Context, message models.Message) error {
// 	return s.producer.ProduceMessage(ctx, message)
// }

// func (s *Service) ReceiveMessages(ctx context.Context) (message string,err error) {
// 	return s.consumer.ConsumeMessages(ctx)
// }
