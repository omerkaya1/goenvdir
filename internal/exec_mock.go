// Code generated by MockGen. DO NOT EDIT.
// Source: internal/exec.go

// Package internal is a generated GoMock package.
package internal

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProgRunner is a mock of ProgRunner interface
type MockProgRunner struct {
	ctrl     *gomock.Controller
	recorder *MockProgRunnerMockRecorder
}

// MockProgRunnerMockRecorder is the mock recorder for MockProgRunner
type MockProgRunnerMockRecorder struct {
	mock *MockProgRunner
}

// NewMockProgRunner creates a new mock instance
func NewMockProgRunner(ctrl *gomock.Controller) *MockProgRunner {
	mock := &MockProgRunner{ctrl: ctrl}
	mock.recorder = &MockProgRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProgRunner) EXPECT() *MockProgRunnerMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockProgRunner) Execute() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute")
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockProgRunnerMockRecorder) Execute() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockProgRunner)(nil).Execute))
}
