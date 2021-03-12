// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sapawarga/phonebook-service/repository (interfaces: PhoneBookI)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/sapawarga/phonebook-service/model"
	reflect "reflect"
)

// MockPhoneBookI is a mock of PhoneBookI interface
type MockPhoneBookI struct {
	ctrl     *gomock.Controller
	recorder *MockPhoneBookIMockRecorder
}

// MockPhoneBookIMockRecorder is the mock recorder for MockPhoneBookI
type MockPhoneBookIMockRecorder struct {
	mock *MockPhoneBookI
}

// NewMockPhoneBookI creates a new mock instance
func NewMockPhoneBookI(ctrl *gomock.Controller) *MockPhoneBookI {
	mock := &MockPhoneBookI{ctrl: ctrl}
	mock.recorder = &MockPhoneBookIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPhoneBookI) EXPECT() *MockPhoneBookIMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockPhoneBookI) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockPhoneBookIMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPhoneBookI)(nil).Delete), arg0, arg1)
}

// GetCategoryNameByID mocks base method
func (m *MockPhoneBookI) GetCategoryNameByID(arg0 context.Context, arg1 int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryNameByID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryNameByID indicates an expected call of GetCategoryNameByID
func (mr *MockPhoneBookIMockRecorder) GetCategoryNameByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryNameByID", reflect.TypeOf((*MockPhoneBookI)(nil).GetCategoryNameByID), arg0, arg1)
}

// GetListPhoneBook mocks base method
func (m *MockPhoneBookI) GetListPhoneBook(arg0 context.Context, arg1 *model.GetListRequest) ([]*model.PhoneBookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListPhoneBook", arg0, arg1)
	ret0, _ := ret[0].([]*model.PhoneBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListPhoneBook indicates an expected call of GetListPhoneBook
func (mr *MockPhoneBookIMockRecorder) GetListPhoneBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListPhoneBook", reflect.TypeOf((*MockPhoneBookI)(nil).GetListPhoneBook), arg0, arg1)
}

// GetLocationNameByID mocks base method
func (m *MockPhoneBookI) GetLocationNameByID(arg0 context.Context, arg1 int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocationNameByID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocationNameByID indicates an expected call of GetLocationNameByID
func (mr *MockPhoneBookIMockRecorder) GetLocationNameByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocationNameByID", reflect.TypeOf((*MockPhoneBookI)(nil).GetLocationNameByID), arg0, arg1)
}

// GetMetaDataPhoneBook mocks base method
func (m *MockPhoneBookI) GetMetaDataPhoneBook(arg0 context.Context, arg1 *model.GetListRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetaDataPhoneBook", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetaDataPhoneBook indicates an expected call of GetMetaDataPhoneBook
func (mr *MockPhoneBookIMockRecorder) GetMetaDataPhoneBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetaDataPhoneBook", reflect.TypeOf((*MockPhoneBookI)(nil).GetMetaDataPhoneBook), arg0, arg1)
}

// GetPhonebookDetailByID mocks base method
func (m *MockPhoneBookI) GetPhonebookDetailByID(arg0 context.Context, arg1 int64) (*model.PhoneBookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhonebookDetailByID", arg0, arg1)
	ret0, _ := ret[0].(*model.PhoneBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhonebookDetailByID indicates an expected call of GetPhonebookDetailByID
func (mr *MockPhoneBookIMockRecorder) GetPhonebookDetailByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhonebookDetailByID", reflect.TypeOf((*MockPhoneBookI)(nil).GetPhonebookDetailByID), arg0, arg1)
}

// Insert mocks base method
func (m *MockPhoneBookI) Insert(arg0 context.Context, arg1 *model.AddPhonebook) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockPhoneBookIMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPhoneBookI)(nil).Insert), arg0, arg1)
}

// Update mocks base method
func (m *MockPhoneBookI) Update(arg0 context.Context, arg1 *model.UpdatePhonebook) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockPhoneBookIMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPhoneBookI)(nil).Update), arg0, arg1)
}
