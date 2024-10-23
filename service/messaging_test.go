package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Businge931/Kafka_and_CLIs/models"
	"github.com/Businge931/Kafka_and_CLIs/service"
	mock_service "github.com/Businge931/Kafka_and_CLIs/service/mock"
)

func TestSendMessage(t *testing.T) {
	type args struct {
		topic   string
		message string
	}
	tests := []struct {
		name           string
		args           args
		mockProducerFn func(mock *mock_service.MockMessageProducer)
		wantErr        error
	}{
		{
			name: "success/message send",
			args: args{
				topic:   "myTopic",
				message: "Hello, Kafka!",
			},
			mockProducerFn: func(mock *mock_service.MockMessageProducer) {
				mock.EXPECT().SendMessage("myTopic", "Hello, Kafka!").Return(nil).Times(1)
			},
			wantErr: nil,
		},
		{
			name: "error/Empty topic ",
			args: args{
				topic:   "",
				message: "Hello, Kafka!",
			},
			mockProducerFn: nil,
			wantErr:        models.ErrEmptyTopic,
		},
		{
			name: "error/Empty message ",
			args: args{
				topic:   "myTopic",
				message: "",
			},
			mockProducerFn: nil,
			wantErr:        models.ErrEmptyMessage,
		},
		{
			name: "error/Retry on failure",
			args: args{
				topic:   "myTopic",
				message: "Hello, Kafka!",
			},
			mockProducerFn: func(mock *mock_service.MockMessageProducer) {
				mock.EXPECT().SendMessage("myTopic", "Hello, Kafka!").Return(models.ErrSendFail).Times(5)
			},
			wantErr: models.ErrMaxRetry,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockProducer := mock_service.NewMockMessageProducer(ctrl)
			if tt.mockProducerFn != nil {
				tt.mockProducerFn(mockProducer)
			}

			svc := service.New(mockProducer, nil)
			err := svc.SendMessage(tt.args.topic, tt.args.message)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestReadMessages(t *testing.T) {
	type args struct {
		topic string
	}
	tests := []struct {
		name           string
		args           args
		mockConsumerFn func(mock *mock_service.MockMessageConsumer)
		wantErr        error
	}{
		{
			name: "Success/message read",
			args: args{
				topic: "myTopic",
			},
			mockConsumerFn: func(mock *mock_service.MockMessageConsumer) {
				mock.EXPECT().ReadMessages("myTopic").Return(nil).Times(1)
			},
			wantErr: nil,
		},
		{
			name: "error/Empty topic",
			args: args{
				topic: "",
			},
			mockConsumerFn: nil,
			wantErr:        models.ErrEmptyTopic,
		},
		{
			name: "error/Failed to read messages",
			args: args{
				topic: "myTopic",
			},
			mockConsumerFn: func(mock *mock_service.MockMessageConsumer) {
				mock.EXPECT().ReadMessages("myTopic").Return(models.ErrReadFail).Times(1)
			},
			wantErr: models.ErrReadFail,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockConsumer := mock_service.NewMockMessageConsumer(ctrl)
			if tt.mockConsumerFn != nil {
				tt.mockConsumerFn(mockConsumer)
			}

			svc := service.New(nil, mockConsumer)
			err := svc.ReadMessages(tt.args.topic)

			if tt.wantErr != nil {
				assert.ErrorContains(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
