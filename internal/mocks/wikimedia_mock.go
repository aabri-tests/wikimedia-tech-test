package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWikiMedia is a mock of WikiMedia interface.
type MockWikiMedia struct {
	ctrl     *gomock.Controller
	recorder *MockWikiMediaMockRecorder
}

// MockWikiMediaMockRecorder is the mock recorder for MockWikiMedia.
type MockWikiMediaMockRecorder struct {
	mock *MockWikiMedia
}

// NewMockWikiMedia creates a new mock instance.
func NewMockWikiMedia(ctrl *gomock.Controller) *MockWikiMedia {
	mock := &MockWikiMedia{ctrl: ctrl}
	mock.recorder = &MockWikiMediaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWikiMedia) EXPECT() *MockWikiMediaMockRecorder {
	return m.recorder
}

// Search mocks base method.
func (m *MockWikiMedia) Search(query, language string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query, language)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockWikiMediaMockRecorder) Search(query, language interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockWikiMedia)(nil).Search), query, language)
}
