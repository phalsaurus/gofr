// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package metrics is a generated GoMock package.
package metrics

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMetric is a mock of Metric interface.
type MockMetric struct {
	ctrl     *gomock.Controller
	recorder *MockMetricMockRecorder
}

// MockMetricMockRecorder is the mock recorder for MockMetric.
type MockMetricMockRecorder struct {
	mock *MockMetric
}

// NewMockMetric creates a new mock instance.
func NewMockMetric(ctrl *gomock.Controller) *MockMetric {
	mock := &MockMetric{ctrl: ctrl}
	mock.recorder = &MockMetricMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetric) EXPECT() *MockMetricMockRecorder {
	return m.recorder
}

// AddCounter mocks base method.
func (m *MockMetric) AddCounter(name string, val float64, labels ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, val}
	for _, a := range labels {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddCounter", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCounter indicates an expected call of AddCounter.
func (mr *MockMetricMockRecorder) AddCounter(name, val interface{}, labels ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, val}, labels...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCounter", reflect.TypeOf((*MockMetric)(nil).AddCounter), varargs...)
}

// IncCounter mocks base method.
func (m *MockMetric) IncCounter(name string, labels ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range labels {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IncCounter", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncCounter indicates an expected call of IncCounter.
func (mr *MockMetricMockRecorder) IncCounter(name interface{}, labels ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, labels...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncCounter", reflect.TypeOf((*MockMetric)(nil).IncCounter), varargs...)
}

// ObserveHistogram mocks base method.
func (m *MockMetric) ObserveHistogram(name string, val float64, labels ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, val}
	for _, a := range labels {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ObserveHistogram", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ObserveHistogram indicates an expected call of ObserveHistogram.
func (mr *MockMetricMockRecorder) ObserveHistogram(name, val interface{}, labels ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, val}, labels...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveHistogram", reflect.TypeOf((*MockMetric)(nil).ObserveHistogram), varargs...)
}

// ObserveSummary mocks base method.
func (m *MockMetric) ObserveSummary(name string, val float64, labels ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, val}
	for _, a := range labels {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ObserveSummary", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ObserveSummary indicates an expected call of ObserveSummary.
func (mr *MockMetricMockRecorder) ObserveSummary(name, val interface{}, labels ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, val}, labels...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveSummary", reflect.TypeOf((*MockMetric)(nil).ObserveSummary), varargs...)
}

// SetGauge mocks base method.
func (m *MockMetric) SetGauge(name string, val float64, labels ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, val}
	for _, a := range labels {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetGauge", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetGauge indicates an expected call of SetGauge.
func (mr *MockMetricMockRecorder) SetGauge(name, val interface{}, labels ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, val}, labels...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGauge", reflect.TypeOf((*MockMetric)(nil).SetGauge), varargs...)
}
