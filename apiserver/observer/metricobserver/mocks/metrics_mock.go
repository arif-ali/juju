// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prometheus/client_golang/prometheus (interfaces: Summary)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	prometheus "github.com/prometheus/client_golang/prometheus"
	io_prometheus_client "github.com/prometheus/client_model/go"
	gomock "go.uber.org/mock/gomock"
)

// MockSummary is a mock of Summary interface.
type MockSummary struct {
	ctrl     *gomock.Controller
	recorder *MockSummaryMockRecorder
}

// MockSummaryMockRecorder is the mock recorder for MockSummary.
type MockSummaryMockRecorder struct {
	mock *MockSummary
}

// NewMockSummary creates a new mock instance.
func NewMockSummary(ctrl *gomock.Controller) *MockSummary {
	mock := &MockSummary{ctrl: ctrl}
	mock.recorder = &MockSummaryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSummary) EXPECT() *MockSummaryMockRecorder {
	return m.recorder
}

// Collect mocks base method.
func (m *MockSummary) Collect(arg0 chan<- prometheus.Metric) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Collect", arg0)
}

// Collect indicates an expected call of Collect.
func (mr *MockSummaryMockRecorder) Collect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collect", reflect.TypeOf((*MockSummary)(nil).Collect), arg0)
}

// Desc mocks base method.
func (m *MockSummary) Desc() *prometheus.Desc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Desc")
	ret0, _ := ret[0].(*prometheus.Desc)
	return ret0
}

// Desc indicates an expected call of Desc.
func (mr *MockSummaryMockRecorder) Desc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Desc", reflect.TypeOf((*MockSummary)(nil).Desc))
}

// Describe mocks base method.
func (m *MockSummary) Describe(arg0 chan<- *prometheus.Desc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Describe", arg0)
}

// Describe indicates an expected call of Describe.
func (mr *MockSummaryMockRecorder) Describe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Describe", reflect.TypeOf((*MockSummary)(nil).Describe), arg0)
}

// Observe mocks base method.
func (m *MockSummary) Observe(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Observe", arg0)
}

// Observe indicates an expected call of Observe.
func (mr *MockSummaryMockRecorder) Observe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Observe", reflect.TypeOf((*MockSummary)(nil).Observe), arg0)
}

// Write mocks base method.
func (m *MockSummary) Write(arg0 *io_prometheus_client.Metric) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockSummaryMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockSummary)(nil).Write), arg0)
}
