// Code generated by MockGen. DO NOT EDIT.
// Source: app/repository/cart_items.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	models "github.com/akashgupta05/shopping-cart-go/app/models"
	gomock "github.com/golang/mock/gomock"
)

// MockCartItemsRepositoryInterface is a mock of CartItemsRepositoryInterface interface.
type MockCartItemsRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCartItemsRepositoryInterfaceMockRecorder
}

// MockCartItemsRepositoryInterfaceMockRecorder is the mock recorder for MockCartItemsRepositoryInterface.
type MockCartItemsRepositoryInterfaceMockRecorder struct {
	mock *MockCartItemsRepositoryInterface
}

// NewMockCartItemsRepositoryInterface creates a new mock instance.
func NewMockCartItemsRepositoryInterface(ctrl *gomock.Controller) *MockCartItemsRepositoryInterface {
	mock := &MockCartItemsRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockCartItemsRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartItemsRepositoryInterface) EXPECT() *MockCartItemsRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCartItemsRepositoryInterface) Delete(sessionID, itemID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", sessionID, itemID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCartItemsRepositoryInterfaceMockRecorder) Delete(sessionID, itemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCartItemsRepositoryInterface)(nil).Delete), sessionID, itemID)
}

// GetByItemID mocks base method.
func (m *MockCartItemsRepositoryInterface) GetByItemID(sessionID, itemId string) (*models.CartItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByItemID", sessionID, itemId)
	ret0, _ := ret[0].(*models.CartItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByItemID indicates an expected call of GetByItemID.
func (mr *MockCartItemsRepositoryInterfaceMockRecorder) GetByItemID(sessionID, itemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByItemID", reflect.TypeOf((*MockCartItemsRepositoryInterface)(nil).GetByItemID), sessionID, itemId)
}

// List mocks base method.
func (m *MockCartItemsRepositoryInterface) List(sessionID string) ([]*models.CartItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", sessionID)
	ret0, _ := ret[0].([]*models.CartItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCartItemsRepositoryInterfaceMockRecorder) List(sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCartItemsRepositoryInterface)(nil).List), sessionID)
}

// Upsert mocks base method.
func (m *MockCartItemsRepositoryInterface) Upsert(cart *models.CartItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", cart)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockCartItemsRepositoryInterfaceMockRecorder) Upsert(cart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockCartItemsRepositoryInterface)(nil).Upsert), cart)
}
