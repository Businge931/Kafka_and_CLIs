package service

type (
	MessageProducer interface {
		SendMessage(topic, message string) error
		Close()
	}

	MessageConsumer interface {
		ReadMessages(topic string) error
		Close()
	}

	MessageService interface {
		SendMessage(topic, message string) error
		ReadMessages(topic string) error
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
