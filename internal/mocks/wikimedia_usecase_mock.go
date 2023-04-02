package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWikiMediaUseCase is a mock of WikiMediaUseCase interface.
type MockWikiMediaUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockWikiMediaUseCaseMockRecorder
}

// MockWikiMediaUseCaseMockRecorder is the mock recorder for MockWikiMediaUseCase.
type MockWikiMediaUseCaseMockRecorder struct {
	mock *MockWikiMediaUseCase
}

// NewMockWikiMediaUseCase creates a new mock instance.
func NewMockWikiMediaUseCase(ctrl *gomock.Controller) *MockWikiMediaUseCase {
	mock := &MockWikiMediaUseCase{ctrl: ctrl}
	mock.recorder = &MockWikiMediaUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWikiMediaUseCase) EXPECT() *MockWikiMediaUseCaseMockRecorder {
	return m.recorder
}

// Search mocks base method.
func (m *MockWikiMediaUseCase) Search(query, language string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query, language)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockWikiMediaUseCaseMockRecorder) Search(query, language interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockWikiMediaUseCase)(nil).Search), query, language)
}
