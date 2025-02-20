// Code generated by MockGen. DO NOT EDIT.
// Source: unit-test/payment/repository (interfaces: HttpConector)
//
// Generated by this command:
//
//	mockgen -build_flags=--mod=mod -destination=mock/xendit_mock.go -package=mock . HttpConector
//

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockHttpConector is a mock of HttpConector interface.
type MockHttpConector struct {
	ctrl     *gomock.Controller
	recorder *MockHttpConectorMockRecorder
	isgomock struct{}
}

// MockHttpConectorMockRecorder is the mock recorder for MockHttpConector.
type MockHttpConectorMockRecorder struct {
	mock *MockHttpConector
}

// NewMockHttpConector creates a new mock instance.
func NewMockHttpConector(ctrl *gomock.Controller) *MockHttpConector {
	mock := &MockHttpConector{ctrl: ctrl}
	mock.recorder = &MockHttpConectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpConector) EXPECT() *MockHttpConectorMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockHttpConector) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockHttpConectorMockRecorder) Do(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHttpConector)(nil).Do), req)
}
