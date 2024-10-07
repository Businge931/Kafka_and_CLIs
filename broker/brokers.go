package broker

// These are generic structs that implements the MessageProducer interface.
// It can work with any messaging broker (e.g., Kafka, RabbitMQ).

type Producer struct {
	produceMessageFunc func(topic, message string) error
	closeFunc          func()
}

type Consumer struct {
	consumeMessagesFunc func(topic string) error
	closeFunc           func()
}

func NewProducer(produceFunc func(topic, message string) error, closeFunc func()) *Producer {
	return &Producer{
		produceMessageFunc: produceFunc,
		closeFunc:          closeFunc,
	}
}

func (p *Producer) SendMessage(topic, message string) error {
	return p.produceMessageFunc(topic, message)
}

func (p *Producer) Close() {
	p.closeFunc()
}

func NewConsumer(consumeFunc func(topic string) error, closeFunc func()) *Consumer {
	return &Consumer{
		consumeMessagesFunc: consumeFunc,
		closeFunc:           closeFunc,
	}
}

func (c *Consumer) ReadMessages(topic string) error {
	return c.consumeMessagesFunc(topic)
}

func (c *Consumer) Close() {
	c.closeFunc()
}

//////*************************************************************

// type Producer struct {
// 	produceMessageFunc func(ctx context.Context, message models.Message) error
// }

// type Consumer struct {
// 	consumeMessagesFunc func(ctx context.Context) (message string,err error)
// }

// // The Producers create new generic Producer instances.
// // The broker-specific message production function is passed as a parameter.

// func NewProducer(produceFunc func(ctx context.Context, message models.Message) error) *Producer {
// 	return &Producer{
// 		produceMessageFunc: produceFunc,
// 	}
// }

// func (p *Producer) ProduceMessage(ctx context.Context, message models.Message) error {
// 	return p.produceMessageFunc(ctx, message)
// }

// func NewConsumer(consumeFunc func(ctx context.Context) (message string,err error)) *Consumer {
// 	return &Consumer{
// 		consumeMessagesFunc: consumeFunc,
// 	}
// }

// func (c *Consumer) ConsumeMessages(ctx context.Context) (message string,err error) {
// 	return c.consumeMessagesFunc(ctx)
// }
