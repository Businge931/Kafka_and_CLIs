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

// Ensure `Service` implements the `MessageService` interface.
var _ MessageService = (*Service)(nil)

func New(producer MessageProducer, consumer MessageConsumer) *Service {
	return &Service{
		producer: producer,
		consumer: consumer,
	}
}

// NewMessageService creates and returns a `MessageService` interface.
// This version abstracts the underlying implementation to promote dependency injection.
func NewMessageService(producer MessageProducer, consumer MessageConsumer) MessageService {
	return New(producer, consumer)
}
