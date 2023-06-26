// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dipdup-net/go-lib/database (interfaces: SchemeCommenter)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSchemeCommenter is a mock of SchemeCommenter interface.
type MockSchemeCommenter struct {
	ctrl     *gomock.Controller
	recorder *MockSchemeCommenterMockRecorder
}

// MockSchemeCommenterMockRecorder is the mock recorder for MockSchemeCommenter.
type MockSchemeCommenterMockRecorder struct {
	mock *MockSchemeCommenter
}

// NewMockSchemeCommenter creates a new mock instance.
func NewMockSchemeCommenter(ctrl *gomock.Controller) *MockSchemeCommenter {
	mock := &MockSchemeCommenter{ctrl: ctrl}
	mock.recorder = &MockSchemeCommenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchemeCommenter) EXPECT() *MockSchemeCommenterMockRecorder {
	return m.recorder
}

// MakeColumnComment mocks base method.
func (m *MockSchemeCommenter) MakeColumnComment(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeColumnComment", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeColumnComment indicates an expected call of MakeColumnComment.
func (mr *MockSchemeCommenterMockRecorder) MakeColumnComment(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeColumnComment", reflect.TypeOf((*MockSchemeCommenter)(nil).MakeColumnComment), arg0, arg1, arg2, arg3)
}

// MakeTableComment mocks base method.
func (m *MockSchemeCommenter) MakeTableComment(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeTableComment", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeTableComment indicates an expected call of MakeTableComment.
func (mr *MockSchemeCommenterMockRecorder) MakeTableComment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeTableComment", reflect.TypeOf((*MockSchemeCommenter)(nil).MakeTableComment), arg0, arg1, arg2)
}