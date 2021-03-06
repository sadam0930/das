// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/reference/studio.go

// Package mock_reference is a generated GoMock package.
package mock_reference

import (
	reference "github.com/yubing24/das/businesslogic/reference"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIStudioRepository is a mock of IStudioRepository interface
type MockIStudioRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIStudioRepositoryMockRecorder
}

// MockIStudioRepositoryMockRecorder is the mock recorder for MockIStudioRepository
type MockIStudioRepositoryMockRecorder struct {
	mock *MockIStudioRepository
}

// NewMockIStudioRepository creates a new mock instance
func NewMockIStudioRepository(ctrl *gomock.Controller) *MockIStudioRepository {
	mock := &MockIStudioRepository{ctrl: ctrl}
	mock.recorder = &MockIStudioRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIStudioRepository) EXPECT() *MockIStudioRepositoryMockRecorder {
	return m.recorder
}

// CreateStudio mocks base method
func (m *MockIStudioRepository) CreateStudio(studio reference.Studio) error {
	ret := m.ctrl.Call(m, "CreateStudio", studio)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStudio indicates an expected call of CreateStudio
func (mr *MockIStudioRepositoryMockRecorder) CreateStudio(studio interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStudio", reflect.TypeOf((*MockIStudioRepository)(nil).CreateStudio), studio)
}

// SearchStudio mocks base method
func (m *MockIStudioRepository) SearchStudio(criteria *reference.SearchStudioCriteria) ([]reference.Studio, error) {
	ret := m.ctrl.Call(m, "SearchStudio", criteria)
	ret0, _ := ret[0].([]reference.Studio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchStudio indicates an expected call of SearchStudio
func (mr *MockIStudioRepositoryMockRecorder) SearchStudio(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchStudio", reflect.TypeOf((*MockIStudioRepository)(nil).SearchStudio), criteria)
}

// DeleteStudio mocks base method
func (m *MockIStudioRepository) DeleteStudio(studio reference.Studio) error {
	ret := m.ctrl.Call(m, "DeleteStudio", studio)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStudio indicates an expected call of DeleteStudio
func (mr *MockIStudioRepositoryMockRecorder) DeleteStudio(studio interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStudio", reflect.TypeOf((*MockIStudioRepository)(nil).DeleteStudio), studio)
}

// UpdateStudio mocks base method
func (m *MockIStudioRepository) UpdateStudio(studio reference.Studio) error {
	ret := m.ctrl.Call(m, "UpdateStudio", studio)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStudio indicates an expected call of UpdateStudio
func (mr *MockIStudioRepositoryMockRecorder) UpdateStudio(studio interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStudio", reflect.TypeOf((*MockIStudioRepository)(nil).UpdateStudio), studio)
}
