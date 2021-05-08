// Code generated by MockGen. DO NOT EDIT.
// Source: hasher.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHasherInterface is a mock of HasherInterface interface.
type MockHasherInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHasherInterfaceMockRecorder
}

// MockHasherInterfaceMockRecorder is the mock recorder for MockHasherInterface.
type MockHasherInterfaceMockRecorder struct {
	mock *MockHasherInterface
}

// NewMockHasherInterface creates a new mock instance.
func NewMockHasherInterface(ctrl *gomock.Controller) *MockHasherInterface {
	mock := &MockHasherInterface{ctrl: ctrl}
	mock.recorder = &MockHasherInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasherInterface) EXPECT() *MockHasherInterfaceMockRecorder {
	return m.recorder
}

// HashInput mocks base method.
func (m *MockHasherInterface) HashInput(input []byte) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashInput", input)
	ret0, _ := ret[0].(string)
	return ret0
}

// HashInput indicates an expected call of HashInput.
func (mr *MockHasherInterfaceMockRecorder) HashInput(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashInput", reflect.TypeOf((*MockHasherInterface)(nil).HashInput), input)
}
