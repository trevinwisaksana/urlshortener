// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/trevinwisaksana/trevin-urlshortener/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/trevinwisaksana/trevin-urlshortener/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateShortURL mocks base method.
func (m *MockStore) CreateShortURL(arg0 context.Context, arg1 db.CreateShortURLParams) (db.Url, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShortURL", arg0, arg1)
	ret0, _ := ret[0].(db.Url)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateShortURL indicates an expected call of CreateShortURL.
func (mr *MockStoreMockRecorder) CreateShortURL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShortURL", reflect.TypeOf((*MockStore)(nil).CreateShortURL), arg0, arg1)
}

// GetLongURL mocks base method.
func (m *MockStore) GetLongURL(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLongURL", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLongURL indicates an expected call of GetLongURL.
func (mr *MockStoreMockRecorder) GetLongURL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLongURL", reflect.TypeOf((*MockStore)(nil).GetLongURL), arg0, arg1)
}