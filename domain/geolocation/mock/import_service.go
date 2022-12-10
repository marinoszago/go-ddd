// Code generated by MockGen. DO NOT EDIT.
// Source: go-ddd/domain/geolocation (interfaces: ImportService)

// Package mock is a generated GoMock package.
package mock

import (
	geolocation "go-ddd/domain/geolocation"
	internalimport "go-ddd/internal/import"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockImportService is a mock of ImportService interface.
type MockImportService struct {
	ctrl     *gomock.Controller
	recorder *MockImportServiceMockRecorder
}

// MockImportServiceMockRecorder is the mock recorder for MockImportService.
type MockImportServiceMockRecorder struct {
	mock *MockImportService
}

// NewMockImportService creates a new mock instance.
func NewMockImportService(ctrl *gomock.Controller) *MockImportService {
	mock := &MockImportService{ctrl: ctrl}
	mock.recorder = &MockImportServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImportService) EXPECT() *MockImportServiceMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockImportService) Save(arg0 []internalimport.ParsedRecord) (geolocation.ImportStats, geolocation.ImportErrors) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(geolocation.ImportStats)
	ret1, _ := ret[1].(geolocation.ImportErrors)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockImportServiceMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockImportService)(nil).Save), arg0)
}