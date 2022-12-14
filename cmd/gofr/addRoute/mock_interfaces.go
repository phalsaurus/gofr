// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock_addRoute is a generated GoMock package.
package addroute

import (
	"os"
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockfileSystem is a mock of fileSystem interface
type MockfileSystem struct {
	ctrl     *gomock.Controller
	recorder *MockfileSystemMockRecorder
}

// MockfileSystemMockRecorder is the mock recorder for MockfileSystem
type MockfileSystemMockRecorder struct {
	mock *MockfileSystem
}

// NewMockfileSystem creates a new mock instance
func NewMockfileSystem(ctrl *gomock.Controller) *MockfileSystem {
	mock := &MockfileSystem{ctrl: ctrl}
	mock.recorder = &MockfileSystemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockfileSystem) EXPECT() *MockfileSystemMockRecorder {
	return m.recorder
}

// Getwd mocks base method
func (m *MockfileSystem) Getwd() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Getwd")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Getwd indicates an expected call of Getwd
func (mr *MockfileSystemMockRecorder) Getwd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Getwd", reflect.TypeOf((*MockfileSystem)(nil).Getwd))
}

// Chdir mocks base method
func (m *MockfileSystem) Chdir(dir string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chdir", dir)
	ret0, _ := ret[0].(error)
	return ret0
}

// Chdir indicates an expected call of Chdir
func (mr *MockfileSystemMockRecorder) Chdir(dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chdir", reflect.TypeOf((*MockfileSystem)(nil).Chdir), dir)
}

// Mkdir mocks base method
func (m *MockfileSystem) Mkdir(name string, perm os.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mkdir", name, perm)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mkdir indicates an expected call of Mkdir
func (mr *MockfileSystemMockRecorder) Mkdir(name, perm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mkdir", reflect.TypeOf((*MockfileSystem)(nil).Mkdir), name, perm)
}

// OpenFile mocks base method
func (m *MockfileSystem) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenFile", name, flag, perm)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenFile indicates an expected call of OpenFile
func (mr *MockfileSystemMockRecorder) OpenFile(name, flag, perm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenFile", reflect.TypeOf((*MockfileSystem)(nil).OpenFile), name, flag, perm)
}

// Stat mocks base method
func (m *MockfileSystem) Stat(name string) (os.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat", name)
	ret0, _ := ret[0].(os.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat
func (mr *MockfileSystemMockRecorder) Stat(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockfileSystem)(nil).Stat), name)
}

// IsNotExist mocks base method
func (m *MockfileSystem) IsNotExist(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNotExist", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNotExist indicates an expected call of IsNotExist
func (mr *MockfileSystemMockRecorder) IsNotExist(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNotExist", reflect.TypeOf((*MockfileSystem)(nil).IsNotExist), err)
}

// Match mocks base method
func (m *MockfileSystem) Match(pattern string, b []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Match", pattern, b)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Match indicates an expected call of Match
func (mr *MockfileSystemMockRecorder) Match(pattern, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Match", reflect.TypeOf((*MockfileSystem)(nil).Match), pattern, b)
}
