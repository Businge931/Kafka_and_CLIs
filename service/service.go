package service

// "context"

// "github.com/Businge931/Kafka_and_CLIs/models"

type (
	MessageProducer interface {
		SendMessage(topic, message string) error
		Close()
	}

	MessageConsumer interface {
		ReadMessages(topic string) error
		Close()
	}

	Service struct {
		producer MessageProducer
		consumer MessageConsumer
	}
)

// type (
// 	MessageProducer interface {
// 		ProduceMessage(ctx context.Context, message models.Message) error
// 	}
// 	MessageConsumer interface {
// 		ConsumeMessages(ctx context.Context) (
// 			// []models.Message
// 		message string,err error)
// 	}

	// Service struct {
	// 	producer MessageProducer
	// 	consumer MessageConsumer
	// }
// )

func New(producer MessageProducer, consumer MessageConsumer) *Service {
	return &Service{
		producer: producer,
		consumer: consumer,
	}
}
