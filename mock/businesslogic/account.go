// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/account.go

// Package mock_businesslogic is a generated GoMock package.
package mock_businesslogic

import (
	businesslogic "github.com/yubing24/das/businesslogic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIAccountRepository is a mock of IAccountRepository interface
type MockIAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountRepositoryMockRecorder
}

// MockIAccountRepositoryMockRecorder is the mock recorder for MockIAccountRepository
type MockIAccountRepositoryMockRecorder struct {
	mock *MockIAccountRepository
}

// NewMockIAccountRepository creates a new mock instance
func NewMockIAccountRepository(ctrl *gomock.Controller) *MockIAccountRepository {
	mock := &MockIAccountRepository{ctrl: ctrl}
	mock.recorder = &MockIAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAccountRepository) EXPECT() *MockIAccountRepositoryMockRecorder {
	return m.recorder
}

// SearchAccount mocks base method
func (m *MockIAccountRepository) SearchAccount(criteria *businesslogic.SearchAccountCriteria) ([]businesslogic.Account, error) {
	ret := m.ctrl.Call(m, "SearchAccount", criteria)
	ret0, _ := ret[0].([]businesslogic.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchAccount indicates an expected call of SearchAccount
func (mr *MockIAccountRepositoryMockRecorder) SearchAccount(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchAccount", reflect.TypeOf((*MockIAccountRepository)(nil).SearchAccount), criteria)
}

// CreateAccount mocks base method
func (m *MockIAccountRepository) CreateAccount(account *businesslogic.Account) error {
	ret := m.ctrl.Call(m, "CreateAccount", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockIAccountRepositoryMockRecorder) CreateAccount(account interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockIAccountRepository)(nil).CreateAccount), account)
}

// UpdateAccount mocks base method
func (m *MockIAccountRepository) UpdateAccount(account businesslogic.Account) error {
	ret := m.ctrl.Call(m, "UpdateAccount", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount
func (mr *MockIAccountRepositoryMockRecorder) UpdateAccount(account interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockIAccountRepository)(nil).UpdateAccount), account)
}

// DeleteAccount mocks base method
func (m *MockIAccountRepository) DeleteAccount(account businesslogic.Account) error {
	ret := m.ctrl.Call(m, "DeleteAccount", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount
func (mr *MockIAccountRepositoryMockRecorder) DeleteAccount(account interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockIAccountRepository)(nil).DeleteAccount), account)
}

// MockICreateAccountStrategy is a mock of ICreateAccountStrategy interface
type MockICreateAccountStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockICreateAccountStrategyMockRecorder
}

// MockICreateAccountStrategyMockRecorder is the mock recorder for MockICreateAccountStrategy
type MockICreateAccountStrategyMockRecorder struct {
	mock *MockICreateAccountStrategy
}

// NewMockICreateAccountStrategy creates a new mock instance
func NewMockICreateAccountStrategy(ctrl *gomock.Controller) *MockICreateAccountStrategy {
	mock := &MockICreateAccountStrategy{ctrl: ctrl}
	mock.recorder = &MockICreateAccountStrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICreateAccountStrategy) EXPECT() *MockICreateAccountStrategyMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method
func (m *MockICreateAccountStrategy) CreateAccount(account businesslogic.Account, password string) error {
	ret := m.ctrl.Call(m, "CreateAccount", account, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockICreateAccountStrategyMockRecorder) CreateAccount(account, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockICreateAccountStrategy)(nil).CreateAccount), account, password)
}
