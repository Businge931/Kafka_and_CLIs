package producer_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Businge931/Kafka_and_CLIs/producer"
	"github.com/Businge931/Kafka_and_CLIs/producer/mocks"

)

// Struct for dependencies
type deps struct {
	producer producer.MessageProducer
}

// Struct for arguments
type args struct {
	topic   string
	message string
}

// Struct for expected outcomes
type expected struct {
	result error
}

func TestSendMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock for MessageProducer interface
	mockProducer := mocks.NewMockMessageProducer(ctrl)

	// Define test cases in a table-driven format
	tests := []struct {
		name        string
		deps        deps
		args        args
		beforeFunc  func() 
		expected    expected
		expectedErr string
	}{
		{
			name: "Success case",
			deps: deps{
				producer: mockProducer,
			},
			args: args{
				topic:   "test-topic",
				message: "test-message",
			},
			beforeFunc: func() {
				// Expect the SendMessage method to be called and return no error
				mockProducer.EXPECT().SendMessage("test-topic", "test-message").Return(nil)
			},
			expected: expected{
				result: nil,
			},
		},
		{
			name: "Error case",
			deps: deps{
				producer: mockProducer,
			},
			args: args{
				topic:   "test-topic",
				message: "test-message",
			},
			beforeFunc: func() {
				// Expect the SendMessage method to be called and return an error
				mockProducer.EXPECT().SendMessage("test-topic", "test-message").Return(errors.New("failed to send message"))
			},
			expected: expected{
				result: errors.New("failed to send message"),
			},
			expectedErr: "failed to send message",
		},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the beforeFunc if defined
			if tt.beforeFunc != nil {
				tt.beforeFunc()
			}

			// Call the SendMessage function
			err := tt.deps.producer.SendMessage(tt.args.topic, tt.args.message)

			// Validate the result
			if tt.expected.result != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
