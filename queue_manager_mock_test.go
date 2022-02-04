// Copyright 2020-present Open Networking Foundation
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: queue_manager.go

// Package dbuf is a generated GoMock package.
package dbuf

import (
	gomock "github.com/golang/mock/gomock"
	. "github.com/omec-project/dbuf/api"
	net "net"
	reflect "reflect"
)

// MockQueueManagerInterface is a mock of QueueManagerInterface interface
type MockQueueManagerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockQueueManagerInterfaceMockRecorder
}

// MockQueueManagerInterfaceMockRecorder is the mock recorder for MockQueueManagerInterface
type MockQueueManagerInterfaceMockRecorder struct {
	mock *MockQueueManagerInterface
}

// NewMockQueueManagerInterface creates a new mock instance
func NewMockQueueManagerInterface(ctrl *gomock.Controller) *MockQueueManagerInterface {
	mock := &MockQueueManagerInterface{ctrl: ctrl}
	mock.recorder = &MockQueueManagerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQueueManagerInterface) EXPECT() *MockQueueManagerInterfaceMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockQueueManagerInterface) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockQueueManagerInterfaceMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockQueueManagerInterface)(nil).Start))
}

// Stop mocks base method
func (m *MockQueueManagerInterface) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockQueueManagerInterfaceMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockQueueManagerInterface)(nil).Stop))
}

// RegisterSubscriber mocks base method
func (m *MockQueueManagerInterface) RegisterSubscriber(arg0 chan Notification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterSubscriber", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterSubscriber indicates an expected call of RegisterSubscriber
func (mr *MockQueueManagerInterfaceMockRecorder) RegisterSubscriber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterSubscriber", reflect.TypeOf((*MockQueueManagerInterface)(nil).RegisterSubscriber), arg0)
}

// UnregisterSubscriber mocks base method
func (m *MockQueueManagerInterface) UnregisterSubscriber(arg0 chan Notification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterSubscriber", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterSubscriber indicates an expected call of UnregisterSubscriber
func (mr *MockQueueManagerInterfaceMockRecorder) UnregisterSubscriber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterSubscriber", reflect.TypeOf((*MockQueueManagerInterface)(nil).UnregisterSubscriber), arg0)
}

// GetState mocks base method
func (m *MockQueueManagerInterface) GetState() GetDbufStateResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState")
	ret0, _ := ret[0].(GetDbufStateResponse)
	return ret0
}

// GetState indicates an expected call of GetState
func (mr *MockQueueManagerInterfaceMockRecorder) GetState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockQueueManagerInterface)(nil).GetState))
}

// GetQueueState mocks base method
func (m *MockQueueManagerInterface) GetQueueState(arg0 uint64) (GetQueueStateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueueState", arg0)
	ret0, _ := ret[0].(GetQueueStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueueState indicates an expected call of GetQueueState
func (mr *MockQueueManagerInterfaceMockRecorder) GetQueueState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueueState", reflect.TypeOf((*MockQueueManagerInterface)(nil).GetQueueState), arg0)
}

// ReleasePackets mocks base method
func (m *MockQueueManagerInterface) ReleasePackets(queueId uint32, dst *net.UDPAddr, drop, passthrough bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleasePackets", queueId, dst, drop, passthrough)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReleasePackets indicates an expected call of ReleasePackets
func (mr *MockQueueManagerInterfaceMockRecorder) ReleasePackets(queueId, dst, drop, passthrough interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleasePackets", reflect.TypeOf((*MockQueueManagerInterface)(nil).ReleasePackets), queueId, dst, drop, passthrough)
}
