// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package app is a generated GoMock package.
package app

import (
	reflect "reflect"
	domain "url-changer/domain"

	gomock "github.com/golang/mock/gomock"
)

// MockKeyGenerator is a mock of KeyGenerator interface.
type MockKeyGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockKeyGeneratorMockRecorder
}

// MockKeyGeneratorMockRecorder is the mock recorder for MockKeyGenerator.
type MockKeyGeneratorMockRecorder struct {
	mock *MockKeyGenerator
}

// NewMockKeyGenerator creates a new mock instance.
func NewMockKeyGenerator(ctrl *gomock.Controller) *MockKeyGenerator {
	mock := &MockKeyGenerator{ctrl: ctrl}
	mock.recorder = &MockKeyGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyGenerator) EXPECT() *MockKeyGeneratorMockRecorder {
	return m.recorder
}

// GetURL mocks base method.
func (m *MockKeyGenerator) GetURL(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURL", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetURL indicates an expected call of GetURL.
func (mr *MockKeyGeneratorMockRecorder) GetURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURL", reflect.TypeOf((*MockKeyGenerator)(nil).GetURL), arg0)
}

// MakeKey mocks base method.
func (m *MockKeyGenerator) MakeKey(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeKey", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeKey indicates an expected call of MakeKey.
func (mr *MockKeyGeneratorMockRecorder) MakeKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeKey", reflect.TypeOf((*MockKeyGenerator)(nil).MakeKey), arg0)
}

// MockUrlSaver is a mock of UrlSaver interface.
type MockUrlSaver struct {
	ctrl     *gomock.Controller
	recorder *MockUrlSaverMockRecorder
}

// MockUrlSaverMockRecorder is the mock recorder for MockUrlSaver.
type MockUrlSaverMockRecorder struct {
	mock *MockUrlSaver
}

// NewMockUrlSaver creates a new mock instance.
func NewMockUrlSaver(ctrl *gomock.Controller) *MockUrlSaver {
	mock := &MockUrlSaver{ctrl: ctrl}
	mock.recorder = &MockUrlSaverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUrlSaver) EXPECT() *MockUrlSaverMockRecorder {
	return m.recorder
}

// GetFullString mocks base method.
func (m *MockUrlSaver) GetFullString(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullString", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullString indicates an expected call of GetFullString.
func (mr *MockUrlSaverMockRecorder) GetFullString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullString", reflect.TypeOf((*MockUrlSaver)(nil).GetFullString), arg0)
}

// Save mocks base method.
func (m *MockUrlSaver) Save(arg0 domain.LongURL, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockUrlSaverMockRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockUrlSaver)(nil).Save), arg0, arg1)
}