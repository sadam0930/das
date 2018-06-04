// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/reference/federation.go

// Package mock_reference is a generated GoMock package.
package mock_reference

import (
	reference "github.com/DancesportSoftware/das/businesslogic/reference"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIFederationRepository is a mock of IFederationRepository interface
type MockIFederationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIFederationRepositoryMockRecorder
}

// MockIFederationRepositoryMockRecorder is the mock recorder for MockIFederationRepository
type MockIFederationRepositoryMockRecorder struct {
	mock *MockIFederationRepository
}

// NewMockIFederationRepository creates a new mock instance
func NewMockIFederationRepository(ctrl *gomock.Controller) *MockIFederationRepository {
	mock := &MockIFederationRepository{ctrl: ctrl}
	mock.recorder = &MockIFederationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIFederationRepository) EXPECT() *MockIFederationRepositoryMockRecorder {
	return m.recorder
}

// CreateFederation mocks base method
func (m *MockIFederationRepository) CreateFederation(federation reference.Federation) error {
	ret := m.ctrl.Call(m, "CreateFederation", federation)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFederation indicates an expected call of CreateFederation
func (mr *MockIFederationRepositoryMockRecorder) CreateFederation(federation interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFederation", reflect.TypeOf((*MockIFederationRepository)(nil).CreateFederation), federation)
}

// SearchFederation mocks base method
func (m *MockIFederationRepository) SearchFederation(criteria *reference.SearchFederationCriteria) ([]reference.Federation, error) {
	ret := m.ctrl.Call(m, "SearchFederation", criteria)
	ret0, _ := ret[0].([]reference.Federation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchFederation indicates an expected call of SearchFederation
func (mr *MockIFederationRepositoryMockRecorder) SearchFederation(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchFederation", reflect.TypeOf((*MockIFederationRepository)(nil).SearchFederation), criteria)
}

// UpdateFederation mocks base method
func (m *MockIFederationRepository) UpdateFederation(federation reference.Federation) error {
	ret := m.ctrl.Call(m, "UpdateFederation", federation)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFederation indicates an expected call of UpdateFederation
func (mr *MockIFederationRepositoryMockRecorder) UpdateFederation(federation interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederation", reflect.TypeOf((*MockIFederationRepository)(nil).UpdateFederation), federation)
}

// DeleteFederation mocks base method
func (m *MockIFederationRepository) DeleteFederation(federation reference.Federation) error {
	ret := m.ctrl.Call(m, "DeleteFederation", federation)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFederation indicates an expected call of DeleteFederation
func (mr *MockIFederationRepositoryMockRecorder) DeleteFederation(federation interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFederation", reflect.TypeOf((*MockIFederationRepository)(nil).DeleteFederation), federation)
}
