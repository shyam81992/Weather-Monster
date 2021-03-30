// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shyam81992/Weather-Monster/db (interfaces: IDB,IRow)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/shyam81992/Weather-Monster/db"
)

// MockIDB is a mock of IDB interface.
type MockIDB struct {
	ctrl     *gomock.Controller
	recorder *MockIDBMockRecorder
}

// MockIDBMockRecorder is the mock recorder for MockIDB.
type MockIDBMockRecorder struct {
	mock *MockIDB
}

// NewMockIDB creates a new mock instance.
func NewMockIDB(ctrl *gomock.Controller) *MockIDB {
	mock := &MockIDB{ctrl: ctrl}
	mock.recorder = &MockIDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDB) EXPECT() *MockIDBMockRecorder {
	return m.recorder
}

// ExecContext mocks base method.
func (m *MockIDB) ExecContext(arg0 context.Context, arg1 string, arg2 ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext.
func (mr *MockIDBMockRecorder) ExecContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockIDB)(nil).ExecContext), varargs...)
}

// QueryRowContext mocks base method.
func (m *MockIDB) QueryRowContext(arg0 context.Context, arg1 string, arg2 ...interface{}) db.IRow {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRowContext", varargs...)
	ret0, _ := ret[0].(db.IRow)
	return ret0
}

// QueryRowContext indicates an expected call of QueryRowContext.
func (mr *MockIDBMockRecorder) QueryRowContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowContext", reflect.TypeOf((*MockIDB)(nil).QueryRowContext), varargs...)
}

// MockIRow is a mock of IRow interface.
type MockIRow struct {
	ctrl     *gomock.Controller
	recorder *MockIRowMockRecorder
}

// MockIRowMockRecorder is the mock recorder for MockIRow.
type MockIRowMockRecorder struct {
	mock *MockIRow
}

// NewMockIRow creates a new mock instance.
func NewMockIRow(ctrl *gomock.Controller) *MockIRow {
	mock := &MockIRow{ctrl: ctrl}
	mock.recorder = &MockIRowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRow) EXPECT() *MockIRowMockRecorder {
	return m.recorder
}

// Scan mocks base method.
func (m *MockIRow) Scan(arg0 ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockIRowMockRecorder) Scan(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockIRow)(nil).Scan), arg0...)
}