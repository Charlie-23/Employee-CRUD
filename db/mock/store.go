// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pzanwar/employee/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/pzanwar/employee/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateEmployee mocks base method.
func (m *MockStore) CreateEmployee(arg0 context.Context, arg1 db.CreateEmployeeParams) (db.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmployee", arg0, arg1)
	ret0, _ := ret[0].(db.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEmployee indicates an expected call of CreateEmployee.
func (mr *MockStoreMockRecorder) CreateEmployee(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmployee", reflect.TypeOf((*MockStore)(nil).CreateEmployee), arg0, arg1)
}

// CreateEmployeeTx mocks base method.
func (m *MockStore) CreateEmployeeTx(arg0 context.Context, arg1 db.CreateEmployeeTxParams) (db.CreateEmployeeTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmployeeTx", arg0, arg1)
	ret0, _ := ret[0].(db.CreateEmployeeTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEmployeeTx indicates an expected call of CreateEmployeeTx.
func (mr *MockStoreMockRecorder) CreateEmployeeTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmployeeTx", reflect.TypeOf((*MockStore)(nil).CreateEmployeeTx), arg0, arg1)
}

// DeleteEmployee mocks base method.
func (m *MockStore) DeleteEmployee(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEmployee", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEmployee indicates an expected call of DeleteEmployee.
func (mr *MockStoreMockRecorder) DeleteEmployee(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEmployee", reflect.TypeOf((*MockStore)(nil).DeleteEmployee), arg0, arg1)
}

// GetEmployeeByID mocks base method.
func (m *MockStore) GetEmployeeByID(arg0 context.Context, arg1 int64) (db.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployeeByID", arg0, arg1)
	ret0, _ := ret[0].(db.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmployeeByID indicates an expected call of GetEmployeeByID.
func (mr *MockStoreMockRecorder) GetEmployeeByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployeeByID", reflect.TypeOf((*MockStore)(nil).GetEmployeeByID), arg0, arg1)
}

// GetEmployeeForUpdate mocks base method.
func (m *MockStore) GetEmployeeForUpdate(arg0 context.Context, arg1 int64) (db.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployeeForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmployeeForUpdate indicates an expected call of GetEmployeeForUpdate.
func (mr *MockStoreMockRecorder) GetEmployeeForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployeeForUpdate", reflect.TypeOf((*MockStore)(nil).GetEmployeeForUpdate), arg0, arg1)
}

// UpdateEmployee mocks base method.
func (m *MockStore) UpdateEmployee(arg0 context.Context, arg1 db.UpdateEmployeeParams) (db.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmployee", arg0, arg1)
	ret0, _ := ret[0].(db.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEmployee indicates an expected call of UpdateEmployee.
func (mr *MockStoreMockRecorder) UpdateEmployee(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmployee", reflect.TypeOf((*MockStore)(nil).UpdateEmployee), arg0, arg1)
}
