// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shyam81992/Weather-Monster/models (interfaces: TemperatureCtlInteface)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockTemperatureCtlInteface is a mock of TemperatureCtlInteface interface.
type MockTemperatureCtlInteface struct {
	ctrl     *gomock.Controller
	recorder *MockTemperatureCtlIntefaceMockRecorder
}

// MockTemperatureCtlIntefaceMockRecorder is the mock recorder for MockTemperatureCtlInteface.
type MockTemperatureCtlIntefaceMockRecorder struct {
	mock *MockTemperatureCtlInteface
}

// NewMockTemperatureCtlInteface creates a new mock instance.
func NewMockTemperatureCtlInteface(ctrl *gomock.Controller) *MockTemperatureCtlInteface {
	mock := &MockTemperatureCtlInteface{ctrl: ctrl}
	mock.recorder = &MockTemperatureCtlIntefaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTemperatureCtlInteface) EXPECT() *MockTemperatureCtlIntefaceMockRecorder {
	return m.recorder
}

// CreateTemperature mocks base method.
func (m *MockTemperatureCtlInteface) CreateTemperature(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateTemperature", arg0)
}

// CreateTemperature indicates an expected call of CreateTemperature.
func (mr *MockTemperatureCtlIntefaceMockRecorder) CreateTemperature(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTemperature", reflect.TypeOf((*MockTemperatureCtlInteface)(nil).CreateTemperature), arg0)
}

// CreateTemperatureTable mocks base method.
func (m *MockTemperatureCtlInteface) CreateTemperatureTable() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateTemperatureTable")
}

// CreateTemperatureTable indicates an expected call of CreateTemperatureTable.
func (mr *MockTemperatureCtlIntefaceMockRecorder) CreateTemperatureTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTemperatureTable", reflect.TypeOf((*MockTemperatureCtlInteface)(nil).CreateTemperatureTable))
}

// GetForecasts mocks base method.
func (m *MockTemperatureCtlInteface) GetForecasts(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetForecasts", arg0)
}

// GetForecasts indicates an expected call of GetForecasts.
func (mr *MockTemperatureCtlIntefaceMockRecorder) GetForecasts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForecasts", reflect.TypeOf((*MockTemperatureCtlInteface)(nil).GetForecasts), arg0)
}