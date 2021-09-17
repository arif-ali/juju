// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/provider/oci/common (interfaces: OCIIdentityClient)

// Package testing is a generated GoMock package.
package testing

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	identity "github.com/oracle/oci-go-sdk/v47/identity"
	reflect "reflect"
)

// MockOCIIdentityClient is a mock of OCIIdentityClient interface
type MockOCIIdentityClient struct {
	ctrl     *gomock.Controller
	recorder *MockOCIIdentityClientMockRecorder
}

// MockOCIIdentityClientMockRecorder is the mock recorder for MockOCIIdentityClient
type MockOCIIdentityClientMockRecorder struct {
	mock *MockOCIIdentityClient
}

// NewMockOCIIdentityClient creates a new mock instance
func NewMockOCIIdentityClient(ctrl *gomock.Controller) *MockOCIIdentityClient {
	mock := &MockOCIIdentityClient{ctrl: ctrl}
	mock.recorder = &MockOCIIdentityClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOCIIdentityClient) EXPECT() *MockOCIIdentityClientMockRecorder {
	return m.recorder
}

// ListAvailabilityDomains mocks base method
func (m *MockOCIIdentityClient) ListAvailabilityDomains(arg0 context.Context, arg1 identity.ListAvailabilityDomainsRequest) (identity.ListAvailabilityDomainsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAvailabilityDomains", arg0, arg1)
	ret0, _ := ret[0].(identity.ListAvailabilityDomainsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAvailabilityDomains indicates an expected call of ListAvailabilityDomains
func (mr *MockOCIIdentityClientMockRecorder) ListAvailabilityDomains(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailabilityDomains", reflect.TypeOf((*MockOCIIdentityClient)(nil).ListAvailabilityDomains), arg0, arg1)
}

// ListCompartments mocks base method
func (m *MockOCIIdentityClient) ListCompartments(arg0 context.Context, arg1 identity.ListCompartmentsRequest) (identity.ListCompartmentsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCompartments", arg0, arg1)
	ret0, _ := ret[0].(identity.ListCompartmentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCompartments indicates an expected call of ListCompartments
func (mr *MockOCIIdentityClientMockRecorder) ListCompartments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCompartments", reflect.TypeOf((*MockOCIIdentityClient)(nil).ListCompartments), arg0, arg1)
}
