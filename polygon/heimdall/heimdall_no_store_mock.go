// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ledgerwatch/erigon/polygon/heimdall (interfaces: HeimdallNoStore)

// Package heimdall is a generated GoMock package.
package heimdall

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHeimdallNoStore is a mock of HeimdallNoStore interface.
type MockHeimdallNoStore struct {
	ctrl     *gomock.Controller
	recorder *MockHeimdallNoStoreMockRecorder
}

// MockHeimdallNoStoreMockRecorder is the mock recorder for MockHeimdallNoStore.
type MockHeimdallNoStoreMockRecorder struct {
	mock *MockHeimdallNoStore
}

// NewMockHeimdallNoStore creates a new mock instance.
func NewMockHeimdallNoStore(ctrl *gomock.Controller) *MockHeimdallNoStore {
	mock := &MockHeimdallNoStore{ctrl: ctrl}
	mock.recorder = &MockHeimdallNoStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHeimdallNoStore) EXPECT() *MockHeimdallNoStoreMockRecorder {
	return m.recorder
}

// FetchCheckpoints mocks base method.
func (m *MockHeimdallNoStore) FetchCheckpoints(arg0 context.Context, arg1, arg2 CheckpointId) ([]*Checkpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCheckpoints", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*Checkpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCheckpoints indicates an expected call of FetchCheckpoints.
func (mr *MockHeimdallNoStoreMockRecorder) FetchCheckpoints(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCheckpoints", reflect.TypeOf((*MockHeimdallNoStore)(nil).FetchCheckpoints), arg0, arg1, arg2)
}

// FetchCheckpointsFromBlock mocks base method.
func (m *MockHeimdallNoStore) FetchCheckpointsFromBlock(arg0 context.Context, arg1 uint64) (Waypoints, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCheckpointsFromBlock", arg0, arg1)
	ret0, _ := ret[0].(Waypoints)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCheckpointsFromBlock indicates an expected call of FetchCheckpointsFromBlock.
func (mr *MockHeimdallNoStoreMockRecorder) FetchCheckpointsFromBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCheckpointsFromBlock", reflect.TypeOf((*MockHeimdallNoStore)(nil).FetchCheckpointsFromBlock), arg0, arg1)
}

// FetchMilestones mocks base method.
func (m *MockHeimdallNoStore) FetchMilestones(arg0 context.Context, arg1, arg2 MilestoneId) ([]*Milestone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMilestones", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*Milestone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMilestones indicates an expected call of FetchMilestones.
func (mr *MockHeimdallNoStoreMockRecorder) FetchMilestones(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMilestones", reflect.TypeOf((*MockHeimdallNoStore)(nil).FetchMilestones), arg0, arg1, arg2)
}

// FetchMilestonesFromBlock mocks base method.
func (m *MockHeimdallNoStore) FetchMilestonesFromBlock(arg0 context.Context, arg1 uint64) (Waypoints, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMilestonesFromBlock", arg0, arg1)
	ret0, _ := ret[0].(Waypoints)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMilestonesFromBlock indicates an expected call of FetchMilestonesFromBlock.
func (mr *MockHeimdallNoStoreMockRecorder) FetchMilestonesFromBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMilestonesFromBlock", reflect.TypeOf((*MockHeimdallNoStore)(nil).FetchMilestonesFromBlock), arg0, arg1)
}

// FetchSpans mocks base method.
func (m *MockHeimdallNoStore) FetchSpans(arg0 context.Context, arg1, arg2 SpanId) ([]*Span, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSpans", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*Span)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSpans indicates an expected call of FetchSpans.
func (mr *MockHeimdallNoStoreMockRecorder) FetchSpans(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSpans", reflect.TypeOf((*MockHeimdallNoStore)(nil).FetchSpans), arg0, arg1, arg2)
}

// FetchSpansFromBlock mocks base method.
func (m *MockHeimdallNoStore) FetchSpansFromBlock(arg0 context.Context, arg1 uint64) ([]*Span, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSpansFromBlock", arg0, arg1)
	ret0, _ := ret[0].([]*Span)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSpansFromBlock indicates an expected call of FetchSpansFromBlock.
func (mr *MockHeimdallNoStoreMockRecorder) FetchSpansFromBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSpansFromBlock", reflect.TypeOf((*MockHeimdallNoStore)(nil).FetchSpansFromBlock), arg0, arg1)
}

// LastCheckpointId mocks base method.
func (m *MockHeimdallNoStore) LastCheckpointId(arg0 context.Context) (CheckpointId, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastCheckpointId", arg0)
	ret0, _ := ret[0].(CheckpointId)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastCheckpointId indicates an expected call of LastCheckpointId.
func (mr *MockHeimdallNoStoreMockRecorder) LastCheckpointId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastCheckpointId", reflect.TypeOf((*MockHeimdallNoStore)(nil).LastCheckpointId), arg0)
}

// LastMilestoneId mocks base method.
func (m *MockHeimdallNoStore) LastMilestoneId(arg0 context.Context) (MilestoneId, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastMilestoneId", arg0)
	ret0, _ := ret[0].(MilestoneId)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastMilestoneId indicates an expected call of LastMilestoneId.
func (mr *MockHeimdallNoStoreMockRecorder) LastMilestoneId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastMilestoneId", reflect.TypeOf((*MockHeimdallNoStore)(nil).LastMilestoneId), arg0)
}

// LastSpanId mocks base method.
func (m *MockHeimdallNoStore) LastSpanId(arg0 context.Context) (SpanId, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastSpanId", arg0)
	ret0, _ := ret[0].(SpanId)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastSpanId indicates an expected call of LastSpanId.
func (mr *MockHeimdallNoStoreMockRecorder) LastSpanId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastSpanId", reflect.TypeOf((*MockHeimdallNoStore)(nil).LastSpanId), arg0)
}

// OnCheckpointEvent mocks base method.
func (m *MockHeimdallNoStore) OnCheckpointEvent(arg0 context.Context, arg1 func(*Checkpoint)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnCheckpointEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnCheckpointEvent indicates an expected call of OnCheckpointEvent.
func (mr *MockHeimdallNoStoreMockRecorder) OnCheckpointEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnCheckpointEvent", reflect.TypeOf((*MockHeimdallNoStore)(nil).OnCheckpointEvent), arg0, arg1)
}

// OnMilestoneEvent mocks base method.
func (m *MockHeimdallNoStore) OnMilestoneEvent(arg0 context.Context, arg1 func(*Milestone)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnMilestoneEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnMilestoneEvent indicates an expected call of OnMilestoneEvent.
func (mr *MockHeimdallNoStoreMockRecorder) OnMilestoneEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnMilestoneEvent", reflect.TypeOf((*MockHeimdallNoStore)(nil).OnMilestoneEvent), arg0, arg1)
}

// OnSpanEvent mocks base method.
func (m *MockHeimdallNoStore) OnSpanEvent(arg0 context.Context, arg1 func(*Span)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnSpanEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnSpanEvent indicates an expected call of OnSpanEvent.
func (mr *MockHeimdallNoStoreMockRecorder) OnSpanEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnSpanEvent", reflect.TypeOf((*MockHeimdallNoStore)(nil).OnSpanEvent), arg0, arg1)
}