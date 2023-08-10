// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/secrets (interfaces: ListSecretsAPI,AddSecretsAPI,GrantRevokeSecretsAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	secrets "github.com/juju/juju/api/client/secrets"
	secrets0 "github.com/juju/juju/core/secrets"
	gomock "go.uber.org/mock/gomock"
)

// MockListSecretsAPI is a mock of ListSecretsAPI interface.
type MockListSecretsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockListSecretsAPIMockRecorder
}

// MockListSecretsAPIMockRecorder is the mock recorder for MockListSecretsAPI.
type MockListSecretsAPIMockRecorder struct {
	mock *MockListSecretsAPI
}

// NewMockListSecretsAPI creates a new mock instance.
func NewMockListSecretsAPI(ctrl *gomock.Controller) *MockListSecretsAPI {
	mock := &MockListSecretsAPI{ctrl: ctrl}
	mock.recorder = &MockListSecretsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListSecretsAPI) EXPECT() *MockListSecretsAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockListSecretsAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockListSecretsAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockListSecretsAPI)(nil).Close))
}

// ListSecrets mocks base method.
func (m *MockListSecretsAPI) ListSecrets(arg0 bool, arg1 secrets0.Filter) ([]secrets.SecretDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", arg0, arg1)
	ret0, _ := ret[0].([]secrets.SecretDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets.
func (mr *MockListSecretsAPIMockRecorder) ListSecrets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockListSecretsAPI)(nil).ListSecrets), arg0, arg1)
}

// MockAddSecretsAPI is a mock of AddSecretsAPI interface.
type MockAddSecretsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAddSecretsAPIMockRecorder
}

// MockAddSecretsAPIMockRecorder is the mock recorder for MockAddSecretsAPI.
type MockAddSecretsAPIMockRecorder struct {
	mock *MockAddSecretsAPI
}

// NewMockAddSecretsAPI creates a new mock instance.
func NewMockAddSecretsAPI(ctrl *gomock.Controller) *MockAddSecretsAPI {
	mock := &MockAddSecretsAPI{ctrl: ctrl}
	mock.recorder = &MockAddSecretsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAddSecretsAPI) EXPECT() *MockAddSecretsAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockAddSecretsAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAddSecretsAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAddSecretsAPI)(nil).Close))
}

// CreateSecret mocks base method.
func (m *MockAddSecretsAPI) CreateSecret(arg0, arg1 string, arg2 map[string]string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecret indicates an expected call of CreateSecret.
func (mr *MockAddSecretsAPIMockRecorder) CreateSecret(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockAddSecretsAPI)(nil).CreateSecret), arg0, arg1, arg2)
}

// MockGrantRevokeSecretsAPI is a mock of GrantRevokeSecretsAPI interface.
type MockGrantRevokeSecretsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockGrantRevokeSecretsAPIMockRecorder
}

// MockGrantRevokeSecretsAPIMockRecorder is the mock recorder for MockGrantRevokeSecretsAPI.
type MockGrantRevokeSecretsAPIMockRecorder struct {
	mock *MockGrantRevokeSecretsAPI
}

// NewMockGrantRevokeSecretsAPI creates a new mock instance.
func NewMockGrantRevokeSecretsAPI(ctrl *gomock.Controller) *MockGrantRevokeSecretsAPI {
	mock := &MockGrantRevokeSecretsAPI{ctrl: ctrl}
	mock.recorder = &MockGrantRevokeSecretsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGrantRevokeSecretsAPI) EXPECT() *MockGrantRevokeSecretsAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockGrantRevokeSecretsAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockGrantRevokeSecretsAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockGrantRevokeSecretsAPI)(nil).Close))
}

// GrantSecret mocks base method.
func (m *MockGrantRevokeSecretsAPI) GrantSecret(arg0 *secrets0.URI, arg1 []string) ([]error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantSecret", arg0, arg1)
	ret0, _ := ret[0].([]error)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GrantSecret indicates an expected call of GrantSecret.
func (mr *MockGrantRevokeSecretsAPIMockRecorder) GrantSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantSecret", reflect.TypeOf((*MockGrantRevokeSecretsAPI)(nil).GrantSecret), arg0, arg1)
}

// RevokeSecret mocks base method.
func (m *MockGrantRevokeSecretsAPI) RevokeSecret(arg0 *secrets0.URI, arg1 []string) ([]error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeSecret", arg0, arg1)
	ret0, _ := ret[0].([]error)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeSecret indicates an expected call of RevokeSecret.
func (mr *MockGrantRevokeSecretsAPIMockRecorder) RevokeSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeSecret", reflect.TypeOf((*MockGrantRevokeSecretsAPI)(nil).RevokeSecret), arg0, arg1)
}
