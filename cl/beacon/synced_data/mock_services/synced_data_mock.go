// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ledgerwatch/erigon/cl/beacon/synced_data (interfaces: SyncedData)
//
// Generated by this command:
//
//	mockgen -typed=true -destination=./mock_services/synced_data_mock.go -package=mock_services . SyncedData
//

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"

	state "github.com/ledgerwatch/erigon/cl/phase1/core/state"
	gomock "go.uber.org/mock/gomock"
)

// MockSyncedData is a mock of SyncedData interface.
type MockSyncedData struct {
	ctrl     *gomock.Controller
	recorder *MockSyncedDataMockRecorder
}

// MockSyncedDataMockRecorder is the mock recorder for MockSyncedData.
type MockSyncedDataMockRecorder struct {
	mock *MockSyncedData
}

// NewMockSyncedData creates a new mock instance.
func NewMockSyncedData(ctrl *gomock.Controller) *MockSyncedData {
	mock := &MockSyncedData{ctrl: ctrl}
	mock.recorder = &MockSyncedDataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyncedData) EXPECT() *MockSyncedDataMockRecorder {
	return m.recorder
}

// HeadSlot mocks base method.
func (m *MockSyncedData) HeadSlot() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeadSlot")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// HeadSlot indicates an expected call of HeadSlot.
func (mr *MockSyncedDataMockRecorder) HeadSlot() *MockSyncedDataHeadSlotCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadSlot", reflect.TypeOf((*MockSyncedData)(nil).HeadSlot))
	return &MockSyncedDataHeadSlotCall{Call: call}
}

// MockSyncedDataHeadSlotCall wrap *gomock.Call
type MockSyncedDataHeadSlotCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSyncedDataHeadSlotCall) Return(arg0 uint64) *MockSyncedDataHeadSlotCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSyncedDataHeadSlotCall) Do(f func() uint64) *MockSyncedDataHeadSlotCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSyncedDataHeadSlotCall) DoAndReturn(f func() uint64) *MockSyncedDataHeadSlotCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HeadState mocks base method.
func (m *MockSyncedData) HeadState() *state.CachingBeaconState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeadState")
	ret0, _ := ret[0].(*state.CachingBeaconState)
	return ret0
}

// HeadState indicates an expected call of HeadState.
func (mr *MockSyncedDataMockRecorder) HeadState() *MockSyncedDataHeadStateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadState", reflect.TypeOf((*MockSyncedData)(nil).HeadState))
	return &MockSyncedDataHeadStateCall{Call: call}
}

// MockSyncedDataHeadStateCall wrap *gomock.Call
type MockSyncedDataHeadStateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSyncedDataHeadStateCall) Return(arg0 *state.CachingBeaconState) *MockSyncedDataHeadStateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSyncedDataHeadStateCall) Do(f func() *state.CachingBeaconState) *MockSyncedDataHeadStateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSyncedDataHeadStateCall) DoAndReturn(f func() *state.CachingBeaconState) *MockSyncedDataHeadStateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HeadStateReader mocks base method.
func (m *MockSyncedData) HeadStateReader() state.BeaconStateReader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeadStateReader")
	ret0, _ := ret[0].(state.BeaconStateReader)
	return ret0
}

// HeadStateReader indicates an expected call of HeadStateReader.
func (mr *MockSyncedDataMockRecorder) HeadStateReader() *MockSyncedDataHeadStateReaderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadStateReader", reflect.TypeOf((*MockSyncedData)(nil).HeadStateReader))
	return &MockSyncedDataHeadStateReaderCall{Call: call}
}

// MockSyncedDataHeadStateReaderCall wrap *gomock.Call
type MockSyncedDataHeadStateReaderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSyncedDataHeadStateReaderCall) Return(arg0 state.BeaconStateReader) *MockSyncedDataHeadStateReaderCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSyncedDataHeadStateReaderCall) Do(f func() state.BeaconStateReader) *MockSyncedDataHeadStateReaderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSyncedDataHeadStateReaderCall) DoAndReturn(f func() state.BeaconStateReader) *MockSyncedDataHeadStateReaderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OnHeadState mocks base method.
func (m *MockSyncedData) OnHeadState(arg0 *state.CachingBeaconState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnHeadState", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnHeadState indicates an expected call of OnHeadState.
func (mr *MockSyncedDataMockRecorder) OnHeadState(arg0 any) *MockSyncedDataOnHeadStateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnHeadState", reflect.TypeOf((*MockSyncedData)(nil).OnHeadState), arg0)
	return &MockSyncedDataOnHeadStateCall{Call: call}
}

// MockSyncedDataOnHeadStateCall wrap *gomock.Call
type MockSyncedDataOnHeadStateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSyncedDataOnHeadStateCall) Return(arg0 error) *MockSyncedDataOnHeadStateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSyncedDataOnHeadStateCall) Do(f func(*state.CachingBeaconState) error) *MockSyncedDataOnHeadStateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSyncedDataOnHeadStateCall) DoAndReturn(f func(*state.CachingBeaconState) error) *MockSyncedDataOnHeadStateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Syncing mocks base method.
func (m *MockSyncedData) Syncing() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Syncing")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Syncing indicates an expected call of Syncing.
func (mr *MockSyncedDataMockRecorder) Syncing() *MockSyncedDataSyncingCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Syncing", reflect.TypeOf((*MockSyncedData)(nil).Syncing))
	return &MockSyncedDataSyncingCall{Call: call}
}

// MockSyncedDataSyncingCall wrap *gomock.Call
type MockSyncedDataSyncingCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSyncedDataSyncingCall) Return(arg0 bool) *MockSyncedDataSyncingCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSyncedDataSyncingCall) Do(f func() bool) *MockSyncedDataSyncingCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSyncedDataSyncingCall) DoAndReturn(f func() bool) *MockSyncedDataSyncingCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
