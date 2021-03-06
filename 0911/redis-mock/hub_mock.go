// Code generated by MockGen. DO NOT EDIT.
// Source: hub.go

// Package main is a generated GoMock package.
package main

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRedisClient is a mock of RedisClient interface
type MockRedisClient struct {
	ctrl     *gomock.Controller
	recorder *MockRedisClientMockRecorder
}

// MockRedisClientMockRecorder is the mock recorder for MockRedisClient
type MockRedisClientMockRecorder struct {
	mock *MockRedisClient
}

// NewMockRedisClient creates a new mock instance
func NewMockRedisClient(ctrl *gomock.Controller) *MockRedisClient {
	mock := &MockRedisClient{ctrl: ctrl}
	mock.recorder = &MockRedisClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRedisClient) EXPECT() *MockRedisClientMockRecorder {
	return m.recorder
}

// Publish mocks base method
func (m *MockRedisClient) Publish(channel string, message interface{}) {
	m.ctrl.Call(m, "Publish", channel, message)
}

// Publish indicates an expected call of Publish
func (mr *MockRedisClientMockRecorder) Publish(channel, message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockRedisClient)(nil).Publish), channel, message)
}

// Subscribe mocks base method
func (m *MockRedisClient) Subscribe(arg0 string) RedisSubscription {
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(RedisSubscription)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockRedisClientMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockRedisClient)(nil).Subscribe), arg0)
}

// MockRedisSubscription is a mock of RedisSubscription interface
type MockRedisSubscription struct {
	ctrl     *gomock.Controller
	recorder *MockRedisSubscriptionMockRecorder
}

// MockRedisSubscriptionMockRecorder is the mock recorder for MockRedisSubscription
type MockRedisSubscriptionMockRecorder struct {
	mock *MockRedisSubscription
}

// NewMockRedisSubscription creates a new mock instance
func NewMockRedisSubscription(ctrl *gomock.Controller) *MockRedisSubscription {
	mock := &MockRedisSubscription{ctrl: ctrl}
	mock.recorder = &MockRedisSubscriptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRedisSubscription) EXPECT() *MockRedisSubscriptionMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockRedisSubscription) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockRedisSubscriptionMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRedisSubscription)(nil).Close))
}

// ReceiveMessage mocks base method
func (m *MockRedisSubscription) ReceiveMessage() (*Message, error) {
	ret := m.ctrl.Call(m, "ReceiveMessage")
	ret0, _ := ret[0].(*Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveMessage indicates an expected call of ReceiveMessage
func (mr *MockRedisSubscriptionMockRecorder) ReceiveMessage() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveMessage", reflect.TypeOf((*MockRedisSubscription)(nil).ReceiveMessage))
}
