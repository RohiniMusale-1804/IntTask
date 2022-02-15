// Code generated by MockGen. DO NOT EDIT.
// Source: services/pkg/service/service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	models "intTask/services/pkg/models"
	reflect "reflect"
)

// MockCityService is a mock of CityService interface
type MockCityService struct {
	ctrl     *gomock.Controller
	recorder *MockCityServiceMockRecorder
}

// MockCityServiceMockRecorder is the mock recorder for MockCityService
type MockCityServiceMockRecorder struct {
	mock *MockCityService
}

// NewMockCityService creates a new mock instance
func NewMockCityService(ctrl *gomock.Controller) *MockCityService {
	mock := &MockCityService{ctrl: ctrl}
	mock.recorder = &MockCityServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCityService) EXPECT() *MockCityServiceMockRecorder {
	return m.recorder
}

// CreateCity mocks base method
func (m *MockCityService) CreateCity(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCity", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCity indicates an expected call of CreateCity
func (mr *MockCityServiceMockRecorder) CreateCity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCity", reflect.TypeOf((*MockCityService)(nil).CreateCity), arg0)
}

// GetCity mocks base method
func (m *MockCityService) GetCity(arg0 context.Context, arg1 string) (*models.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCity", arg0, arg1)
	ret0, _ := ret[0].(*models.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCity indicates an expected call of GetCity
func (mr *MockCityServiceMockRecorder) GetCity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCity", reflect.TypeOf((*MockCityService)(nil).GetCity), arg0, arg1)
}
