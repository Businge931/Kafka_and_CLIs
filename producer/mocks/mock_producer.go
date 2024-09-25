// Code generated by MockGen. DO NOT EDIT.
// Source: producer/producer.go
//
// Generated by this command:
//
//	mockgen -source=producer/producer.go -destination=producer/mocks/mock_producer.go -package=mocks
//

// Package producer is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMessageProducer is a mock of MessageProducer interface.
type MockMessageProducer struct {
	ctrl     *gomock.Controller
	recorder *MockMessageProducerMockRecorder
}

// MockMessageProducerMockRecorder is the mock recorder for MockMessageProducer.
type MockMessageProducerMockRecorder struct {
	mock *MockMessageProducer
}

// NewMockMessageProducer creates a new mock instance.
func NewMockMessageProducer(ctrl *gomock.Controller) *MockMessageProducer {
	mock := &MockMessageProducer{ctrl: ctrl}
	mock.recorder = &MockMessageProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageProducer) EXPECT() *MockMessageProducerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockMessageProducer) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockMessageProducerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMessageProducer)(nil).Close))
}

// SendMessage mocks base method.
func (m *MockMessageProducer) SendMessage(topic, message string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", topic, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockMessageProducerMockRecorder) SendMessage(topic, message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockMessageProducer)(nil).SendMessage), topic, message)
}