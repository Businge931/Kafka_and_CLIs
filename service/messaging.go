package service

func (s *Service) SendMessage(topic, message string) error {
	return s.producer.SendMessage(topic, message)
}

func (s *Service) ReceiveMessages(topic string) error {
	return s.consumer.ReadMessages(topic)
}
