// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/organizer.go

// Package mock_businesslogic is a generated GoMock package.
package mock_businesslogic

import (
	businesslogic "github.com/yubing24/das/businesslogic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIOrganizerProvisionRepository is a mock of IOrganizerProvisionRepository interface
type MockIOrganizerProvisionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIOrganizerProvisionRepositoryMockRecorder
}

// MockIOrganizerProvisionRepositoryMockRecorder is the mock recorder for MockIOrganizerProvisionRepository
type MockIOrganizerProvisionRepositoryMockRecorder struct {
	mock *MockIOrganizerProvisionRepository
}

// NewMockIOrganizerProvisionRepository creates a new mock instance
func NewMockIOrganizerProvisionRepository(ctrl *gomock.Controller) *MockIOrganizerProvisionRepository {
	mock := &MockIOrganizerProvisionRepository{ctrl: ctrl}
	mock.recorder = &MockIOrganizerProvisionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIOrganizerProvisionRepository) EXPECT() *MockIOrganizerProvisionRepositoryMockRecorder {
	return m.recorder
}

// CreateOrganizerProvision mocks base method
func (m *MockIOrganizerProvisionRepository) CreateOrganizerProvision(provision businesslogic.OrganizerProvision) error {
	ret := m.ctrl.Call(m, "CreateOrganizerProvision", provision)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrganizerProvision indicates an expected call of CreateOrganizerProvision
func (mr *MockIOrganizerProvisionRepositoryMockRecorder) CreateOrganizerProvision(provision interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganizerProvision", reflect.TypeOf((*MockIOrganizerProvisionRepository)(nil).CreateOrganizerProvision), provision)
}

// UpdateOrganizerProvision mocks base method
func (m *MockIOrganizerProvisionRepository) UpdateOrganizerProvision(provision businesslogic.OrganizerProvision) error {
	ret := m.ctrl.Call(m, "UpdateOrganizerProvision", provision)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrganizerProvision indicates an expected call of UpdateOrganizerProvision
func (mr *MockIOrganizerProvisionRepositoryMockRecorder) UpdateOrganizerProvision(provision interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrganizerProvision", reflect.TypeOf((*MockIOrganizerProvisionRepository)(nil).UpdateOrganizerProvision), provision)
}

// DeleteOrganizerProvision mocks base method
func (m *MockIOrganizerProvisionRepository) DeleteOrganizerProvision(provision businesslogic.OrganizerProvision) error {
	ret := m.ctrl.Call(m, "DeleteOrganizerProvision", provision)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrganizerProvision indicates an expected call of DeleteOrganizerProvision
func (mr *MockIOrganizerProvisionRepositoryMockRecorder) DeleteOrganizerProvision(provision interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrganizerProvision", reflect.TypeOf((*MockIOrganizerProvisionRepository)(nil).DeleteOrganizerProvision), provision)
}

// SearchOrganizerProvision mocks base method
func (m *MockIOrganizerProvisionRepository) SearchOrganizerProvision(criteria *businesslogic.SearchOrganizerProvisionCriteria) ([]businesslogic.OrganizerProvision, error) {
	ret := m.ctrl.Call(m, "SearchOrganizerProvision", criteria)
	ret0, _ := ret[0].([]businesslogic.OrganizerProvision)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchOrganizerProvision indicates an expected call of SearchOrganizerProvision
func (mr *MockIOrganizerProvisionRepositoryMockRecorder) SearchOrganizerProvision(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchOrganizerProvision", reflect.TypeOf((*MockIOrganizerProvisionRepository)(nil).SearchOrganizerProvision), criteria)
}
