package service_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/mock/gomock"

// 	"github.com/Businge931/Kafka_and_CLIs/models"
// 	"github.com/Businge931/Kafka_and_CLIs/service"
// 	mock_service "github.com/Businge931/Kafka_and_CLIs/service/mock"
// )

// func TestService_SendMessage(t *testing.T) {
// 	type args struct {
// 		topic   string
// 		message string
// 	}

// 	tests := []struct {
// 		name        string
// 		args        args
// 		before      func(producer *mock_service.MockMessageProducer)
// 		expectedErr error
// 	}{
// 		{
// 			name: "success/message send",
// 			args: args{
// 				topic:   "myTopic",
// 				message: "hello world",
// 			},
// 			before: func(producer *mock_service.MockMessageProducer) {
// 				producer.EXPECT().SendMessage("myTopic", "hello world").Return(nil)
// 			},
// 			expectedErr: nil,
// 		},
// 		{
// 			name: "error/empty topic",
// 			args: args{
// 				topic:   "",
// 				message: "hello world",
// 			},
// 			before:      func(_ *mock_service.MockMessageProducer) {},
// 			expectedErr: models.ErrEmptyTopic,
// 		},
// 		{
// 			name: "error/empty message",
// 			args: args{
// 				topic:   "myTopic",
// 				message: "",
// 			},
// 			before:      func(_ *mock_service.MockMessageProducer) {},
// 			expectedErr: models.ErrEmptyMessage,
// 		},
// 		{
// 			name: "error/producer error",
// 			args: args{
// 				topic:   "myTopic",
// 				message: "hello world",
// 			},
// 			before: func(producer *mock_service.MockMessageProducer) {
// 				producer.EXPECT().SendMessage("myTopic", "hello world").Return(models.ErrSendFail)
// 			},
// 			expectedErr: models.ErrSendMessageToTopic,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			mockProducer := mock_service.NewMockMessageProducer(ctrl)
// 			tt.before(mockProducer)

// 			svc := service.New(mockProducer, nil)
// 			err := svc.SendMessage(tt.args.topic, tt.args.message)

// 			if tt.expectedErr != nil {
// 				assert.EqualError(t, err, tt.expectedErr.Error())
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }

// func TestService_ReceiveMessages(t *testing.T) {
// 	type args struct {
// 		topic string
// 	}

// 	tests := []struct {
// 		name        string
// 		args        args
// 		before      func(consumer *mock_service.MockMessageConsumer)
// 		expectedErr error
// 	}{
// 		{
// 			name: "success/message receive",
// 			args: args{
// 				topic: "myTopic",
// 			},
// 			before: func(consumer *mock_service.MockMessageConsumer) {
// 				consumer.EXPECT().ReadMessages("myTopic").Return(nil)
// 			},
// 			expectedErr: nil,
// 		},
// 		{
// 			name: "error/empty topic",
// 			args: args{
// 				topic: "",
// 			},
// 			before:      func(_ *mock_service.MockMessageConsumer) {},
// 			expectedErr: models.ErrEmptyTopic,
// 		},
// 		{
// 			name: "error/consumer error",
// 			args: args{
// 				topic: "myTopic",
// 			},
// 			before: func(consumer *mock_service.MockMessageConsumer) {
// 				consumer.EXPECT().ReadMessages("myTopic").Return(models.ErrReadFail)
// 			},
// 			expectedErr: models.ErrReceiveMessageFromTopic,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			mockConsumer := mock_service.NewMockMessageConsumer(ctrl)
// 			tt.before(mockConsumer)

// 			svc := service.New(nil, mockConsumer)
// 			err := svc.ReceiveMessages(tt.args.topic)

// 			if tt.expectedErr != nil {
// 				assert.EqualError(t, err, tt.expectedErr.Error())
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }
