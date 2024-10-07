package models

import "errors"

var (
	ErrMaxRetry       = errors.New("failed to send message after max retries: ")
	ErrCreateConsumer = errors.New("failed to create consumer")
	ErrSubscribeTopic = errors.New("failed to subscribe to topic")
	ErrConsumer       = errors.New("consumer error")
	ErrEmptyTopic     = errors.New("topic cannot be empty")
	ErrEmptyMessage   = errors.New("message cannot be empty")
	ErrEventType      = errors.New("unexpected event type received")
)
