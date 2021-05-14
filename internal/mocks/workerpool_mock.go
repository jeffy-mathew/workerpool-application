// Code generated by MockGen. DO NOT EDIT.
// Source: workerpool.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	workerpool "workerpool-application/internal/pkg/workerpool"

	gomock "github.com/golang/mock/gomock"
)

// MockWorkerPoolInterface is a mock of WorkerPoolInterface interface.
type MockWorkerPoolInterface struct {
	ctrl     *gomock.Controller
	recorder *MockWorkerPoolInterfaceMockRecorder
}

// MockWorkerPoolInterfaceMockRecorder is the mock recorder for MockWorkerPoolInterface.
type MockWorkerPoolInterfaceMockRecorder struct {
	mock *MockWorkerPoolInterface
}

// NewMockWorkerPoolInterface creates a new mock instance.
func NewMockWorkerPoolInterface(ctrl *gomock.Controller) *MockWorkerPoolInterface {
	mock := &MockWorkerPoolInterface{ctrl: ctrl}
	mock.recorder = &MockWorkerPoolInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkerPoolInterface) EXPECT() *MockWorkerPoolInterfaceMockRecorder {
	return m.recorder
}

// DoWork mocks base method.
func (m *MockWorkerPoolInterface) DoWork(endpoints <-chan string, responses chan<- workerpool.WorkerResp) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DoWork", endpoints, responses)
}

// DoWork indicates an expected call of DoWork.
func (mr *MockWorkerPoolInterfaceMockRecorder) DoWork(endpoints, responses interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoWork", reflect.TypeOf((*MockWorkerPoolInterface)(nil).DoWork), endpoints, responses)
}