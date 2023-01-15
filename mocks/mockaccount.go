// Code generated by MockGen. DO NOT EDIT.
// Source: mmt.com/lolbank/ports (interfaces: AccountRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "mmt.com/lolbank/domain"
)

// MockAccountRepo is a mock of AccountRepo interface.
type MockAccountRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepoMockRecorder
}

// MockAccountRepoMockRecorder is the mock recorder for MockAccountRepo.
type MockAccountRepoMockRecorder struct {
	mock *MockAccountRepo
}

// NewMockAccountRepo creates a new mock instance.
func NewMockAccountRepo(ctrl *gomock.Controller) *MockAccountRepo {
	mock := &MockAccountRepo{ctrl: ctrl}
	mock.recorder = &MockAccountRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepo) EXPECT() *MockAccountRepoMockRecorder {
	return m.recorder
}

// AddAccount mocks base method.
func (m *MockAccountRepo) AddAccount(arg0 domain.Account) domain.Account {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccount", arg0)
	ret0, _ := ret[0].(domain.Account)
	return ret0
}

// AddAccount indicates an expected call of AddAccount.
func (mr *MockAccountRepoMockRecorder) AddAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccount", reflect.TypeOf((*MockAccountRepo)(nil).AddAccount), arg0)
}

// GetAccount mocks base method.
func (m *MockAccountRepo) GetAccount(arg0 string) (domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0)
	ret0, _ := ret[0].(domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountRepoMockRecorder) GetAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountRepo)(nil).GetAccount), arg0)
}

// GetAccountByIdNumber mocks base method.
func (m *MockAccountRepo) GetAccountByIdNumber(arg0 string) (domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByIdNumber", arg0)
	ret0, _ := ret[0].(domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByIdNumber indicates an expected call of GetAccountByIdNumber.
func (mr *MockAccountRepoMockRecorder) GetAccountByIdNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByIdNumber", reflect.TypeOf((*MockAccountRepo)(nil).GetAccountByIdNumber), arg0)
}

// UpdateAccount mocks base method.
func (m *MockAccountRepo) UpdateAccount(arg0 domain.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockAccountRepoMockRecorder) UpdateAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockAccountRepo)(nil).UpdateAccount), arg0)
}