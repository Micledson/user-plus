// Code generated by MockGen. DO NOT EDIT.
// Source: services/interfaces/interfaces.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	reflect "reflect"
	errs "user-plus/domain/errs"
	response "user-plus/endpoints/dto/response"

	gomock "github.com/golang/mock/gomock"
)

// MockIUserService is a mock of IUserService interface.
type MockIUserService struct {
	ctrl     *gomock.Controller
	recorder *MockIUserServiceMockRecorder
}

// MockIUserServiceMockRecorder is the mock recorder for MockIUserService.
type MockIUserServiceMockRecorder struct {
	mock *MockIUserService
}

// NewMockIUserService creates a new mock instance.
func NewMockIUserService(ctrl *gomock.Controller) *MockIUserService {
	mock := &MockIUserService{ctrl: ctrl}
	mock.recorder = &MockIUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserService) EXPECT() *MockIUserServiceMockRecorder {
	return m.recorder
}

// FindUserByEmail mocks base method.
func (m *MockIUserService) FindUserByEmail(email string) (*response.UserReponse, *errs.ApiErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", email)
	ret0, _ := ret[0].(*response.UserReponse)
	ret1, _ := ret[1].(*errs.ApiErr)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockIUserServiceMockRecorder) FindUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockIUserService)(nil).FindUserByEmail), email)
}
