// Code generated by MockGen. DO NOT EDIT.
// Source: app/repository/access_tokens.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	models "github.com/akashgupta05/shopping-cart-go/app/models"
	gomock "github.com/golang/mock/gomock"
)

// MockAccessTokensRepositoryInterface is a mock of AccessTokensRepositoryInterface interface.
type MockAccessTokensRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAccessTokensRepositoryInterfaceMockRecorder
}

// MockAccessTokensRepositoryInterfaceMockRecorder is the mock recorder for MockAccessTokensRepositoryInterface.
type MockAccessTokensRepositoryInterfaceMockRecorder struct {
	mock *MockAccessTokensRepositoryInterface
}

// NewMockAccessTokensRepositoryInterface creates a new mock instance.
func NewMockAccessTokensRepositoryInterface(ctrl *gomock.Controller) *MockAccessTokensRepositoryInterface {
	mock := &MockAccessTokensRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockAccessTokensRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessTokensRepositoryInterface) EXPECT() *MockAccessTokensRepositoryInterfaceMockRecorder {
	return m.recorder
}

// MarkInactive mocks base method.
func (m *MockAccessTokensRepositoryInterface) MarkInactive(token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkInactive", token)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkInactive indicates an expected call of MarkInactive.
func (mr *MockAccessTokensRepositoryInterfaceMockRecorder) MarkInactive(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkInactive", reflect.TypeOf((*MockAccessTokensRepositoryInterface)(nil).MarkInactive), token)
}

// MarkInactiveForUser mocks base method.
func (m *MockAccessTokensRepositoryInterface) MarkInactiveForUser(userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkInactiveForUser", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkInactiveForUser indicates an expected call of MarkInactiveForUser.
func (mr *MockAccessTokensRepositoryInterfaceMockRecorder) MarkInactiveForUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkInactiveForUser", reflect.TypeOf((*MockAccessTokensRepositoryInterface)(nil).MarkInactiveForUser), userID)
}

// Upsert mocks base method.
func (m *MockAccessTokensRepositoryInterface) Upsert(accessToken *models.AccessToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", accessToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockAccessTokensRepositoryInterfaceMockRecorder) Upsert(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockAccessTokensRepositoryInterface)(nil).Upsert), accessToken)
}

// ValidateToken mocks base method.
func (m *MockAccessTokensRepositoryInterface) ValidateToken(token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockAccessTokensRepositoryInterfaceMockRecorder) ValidateToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockAccessTokensRepositoryInterface)(nil).ValidateToken), token)
}