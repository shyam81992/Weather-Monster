// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shyam81992/Weather-Monster/models (interfaces: WebHookCtlInteface)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockWebHookCtlInteface is a mock of WebHookCtlInteface interface.
type MockWebHookCtlInteface struct {
	ctrl     *gomock.Controller
	recorder *MockWebHookCtlIntefaceMockRecorder
}

// MockWebHookCtlIntefaceMockRecorder is the mock recorder for MockWebHookCtlInteface.
type MockWebHookCtlIntefaceMockRecorder struct {
	mock *MockWebHookCtlInteface
}

// NewMockWebHookCtlInteface creates a new mock instance.
func NewMockWebHookCtlInteface(ctrl *gomock.Controller) *MockWebHookCtlInteface {
	mock := &MockWebHookCtlInteface{ctrl: ctrl}
	mock.recorder = &MockWebHookCtlIntefaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebHookCtlInteface) EXPECT() *MockWebHookCtlIntefaceMockRecorder {
	return m.recorder
}

// CreateWebHookTable mocks base method.
func (m *MockWebHookCtlInteface) CreateWebHookTable() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateWebHookTable")
}

// CreateWebHookTable indicates an expected call of CreateWebHookTable.
func (mr *MockWebHookCtlIntefaceMockRecorder) CreateWebHookTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWebHookTable", reflect.TypeOf((*MockWebHookCtlInteface)(nil).CreateWebHookTable))
}

// CreateWebHooks mocks base method.
func (m *MockWebHookCtlInteface) CreateWebHooks(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateWebHooks", arg0)
}

// CreateWebHooks indicates an expected call of CreateWebHooks.
func (mr *MockWebHookCtlIntefaceMockRecorder) CreateWebHooks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWebHooks", reflect.TypeOf((*MockWebHookCtlInteface)(nil).CreateWebHooks), arg0)
}

// DeleteWebHooks mocks base method.
func (m *MockWebHookCtlInteface) DeleteWebHooks(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteWebHooks", arg0)
}

// DeleteWebHooks indicates an expected call of DeleteWebHooks.
func (mr *MockWebHookCtlIntefaceMockRecorder) DeleteWebHooks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWebHooks", reflect.TypeOf((*MockWebHookCtlInteface)(nil).DeleteWebHooks), arg0)
}
