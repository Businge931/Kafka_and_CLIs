package consumer_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Businge931/Kafka_and_CLIs/consumer/mocks"
)

// Struct for dependencies
type deps struct {
	consumer *mocks.MockMessageConsumer
}

// Struct for arguments
type args struct {
	topic   string
	testing bool
}

// Struct for expected outcomes
type expected struct {
	result error
}

func TestKafkaConsumer_ReadMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock for MessageConsumer interface
	mockConsumer := mocks.NewMockMessageConsumer(ctrl)

	// Define test cases
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
				consumer: mockConsumer,
			},
			args: args{
				topic:   "test-topic",
				testing: false,
			},
			beforeFunc: func() {
				// Expect ReadMessages to be called and return no error
				mockConsumer.EXPECT().ReadMessages("test-topic", false).Return(nil)
			},
			expected: expected{
				result: nil,
			},
		},
		{
			name: "Error case",
			deps: deps{
				consumer: mockConsumer,
			},
			args: args{
				topic:   "test-topic",
				testing: false,
			},
			beforeFunc: func() {
				// Expect ReadMessages to be called and return an error
				mockConsumer.EXPECT().ReadMessages("test-topic", false).Return(errors.New("failed to read messages"))
			},
			expected: expected{
				result: errors.New("failed to read messages"),
			},
			expectedErr: "failed to read messages",
		},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the beforeFunc if defined
			if tt.beforeFunc != nil {
				tt.beforeFunc()
			}

			// Call the ReadMessages function
			err := tt.deps.consumer.ReadMessages(tt.args.topic, tt.args.testing)

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
