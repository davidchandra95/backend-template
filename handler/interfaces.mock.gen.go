// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package handler is a generated GoMock package.
package handler

import (
	gomock "github.com/golang/mock/gomock"
	v4 "github.com/labstack/echo/v4"
	reflect "reflect"
)

// MockServerInterface is a mock of ServerInterface interface
type MockServerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockServerInterfaceMockRecorder
}

// MockServerInterfaceMockRecorder is the mock recorder for MockServerInterface
type MockServerInterfaceMockRecorder struct {
	mock *MockServerInterface
}

// NewMockServerInterface creates a new mock instance
func NewMockServerInterface(ctrl *gomock.Controller) *MockServerInterface {
	mock := &MockServerInterface{ctrl: ctrl}
	mock.recorder = &MockServerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerInterface) EXPECT() *MockServerInterfaceMockRecorder {
	return m.recorder
}

// RegistAccount mocks base method
func (m *MockServerInterface) RegistAccount(ctx v4.Context, params RegistrationParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistAccount", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegistAccount indicates an expected call of RegistAccount
func (mr *MockServerInterfaceMockRecorder) RegistAccount(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistAccount", reflect.TypeOf((*MockServerInterface)(nil).RegistAccount), ctx, params)
}

// Login mocks base method
func (m *MockServerInterface) Login(ctx v4.Context, params LoginParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login
func (mr *MockServerInterfaceMockRecorder) Login(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockServerInterface)(nil).Login), ctx, params)
}

// UpdateAccount mocks base method
func (m *MockServerInterface) UpdateAccount(ctx v4.Context, params UpdateAccountParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount
func (mr *MockServerInterfaceMockRecorder) UpdateAccount(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockServerInterface)(nil).UpdateAccount), ctx, params)
}

// GetAccountByID mocks base method
func (m *MockServerInterface) GetAccountByID(ctx v4.Context, accountID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByID", ctx, accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAccountByID indicates an expected call of GetAccountByID
func (mr *MockServerInterfaceMockRecorder) GetAccountByID(ctx, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByID", reflect.TypeOf((*MockServerInterface)(nil).GetAccountByID), ctx, accountID)
}