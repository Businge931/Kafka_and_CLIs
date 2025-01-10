package models

import "errors"

var (
	ErrMaxRetry                = errors.New("failed to send message after max retries: ")
	ErrCreateConsumer          = errors.New("failed to create consumer")
	ErrSubscribeTopic          = errors.New("failed to subscribe to topic")
	ErrConsumer                = errors.New("consumer error")
	ErrEmptyTopic              = errors.New("topic cannot be empty")
	ErrEmptyMessage            = errors.New("message cannot be empty")
	ErrEventType               = errors.New("unexpected event type received")
	ErrSendFail                = errors.New("failed to send")
	ErrReadFail                = errors.New("failed to read")
	ErrSendMessageToTopic      = errors.New("error sending message to topic myTopic: failed to send")
	ErrReceiveMessageFromTopic = errors.New("error receiving messages from topic myTopic: failed to read")
	ErrInvalidTopic            = errors.New("invalid topic")
)
