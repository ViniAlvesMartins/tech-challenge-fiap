// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	types "github.com/aws/aws-sdk-go-v2/service/sqs/types"
	gomock "github.com/golang/mock/gomock"
)

// MockQueueService is a mock of QueueService interface.
type MockQueueService struct {
	ctrl     *gomock.Controller
	recorder *MockQueueServiceMockRecorder
}

// MockQueueServiceMockRecorder is the mock recorder for MockQueueService.
type MockQueueServiceMockRecorder struct {
	mock *MockQueueService
}

// NewMockQueueService creates a new mock instance.
func NewMockQueueService(ctrl *gomock.Controller) *MockQueueService {
	mock := &MockQueueService{ctrl: ctrl}
	mock.recorder = &MockQueueServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueueService) EXPECT() *MockQueueServiceMockRecorder {
	return m.recorder
}

// DeleteMessage mocks base method.
func (m *MockQueueService) DeleteMessage(queueURL, receiptHandle string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessage", queueURL, receiptHandle)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessage indicates an expected call of DeleteMessage.
func (mr *MockQueueServiceMockRecorder) DeleteMessage(queueURL, receiptHandle interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessage", reflect.TypeOf((*MockQueueService)(nil).DeleteMessage), queueURL, receiptHandle)
}

// ReceiveMessage mocks base method.
func (m *MockQueueService) ReceiveMessage(queueUrl string) (*types.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveMessage", queueUrl)
	ret0, _ := ret[0].(*types.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveMessage indicates an expected call of ReceiveMessage.
func (mr *MockQueueServiceMockRecorder) ReceiveMessage(queueUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveMessage", reflect.TypeOf((*MockQueueService)(nil).ReceiveMessage), queueUrl)
}

// SendMessage mocks base method.
func (m *MockQueueService) SendMessage(queueUrl, message, messageGroupId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", queueUrl, message, messageGroupId)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockQueueServiceMockRecorder) SendMessage(queueUrl, message, messageGroupId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockQueueService)(nil).SendMessage), queueUrl, message, messageGroupId)
}
