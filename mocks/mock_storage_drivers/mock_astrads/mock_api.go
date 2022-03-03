// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/storage_drivers/astrads/api (interfaces: AstraDS)

// Package mock_api is a generated GoMock package.
package mock_api

import (
	context "context"
	reflect "reflect"
	time "time"

	roaring "github.com/RoaringBitmap/roaring"
	gomock "github.com/golang/mock/gomock"
	api "github.com/netapp/trident/storage_drivers/astrads/api"
)

// MockAstraDS is a mock of AstraDS interface.
type MockAstraDS struct {
	ctrl     *gomock.Controller
	recorder *MockAstraDSMockRecorder
}

// MockAstraDSMockRecorder is the mock recorder for MockAstraDS.
type MockAstraDSMockRecorder struct {
	mock *MockAstraDS
}

// NewMockAstraDS creates a new mock instance.
func NewMockAstraDS(ctrl *gomock.Controller) *MockAstraDS {
	mock := &MockAstraDS{ctrl: ctrl}
	mock.recorder = &MockAstraDSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAstraDS) EXPECT() *MockAstraDSMockRecorder {
	return m.recorder
}

// CreateSnapshot mocks base method.
func (m *MockAstraDS) CreateSnapshot(arg0 context.Context, arg1 *api.Snapshot) (*api.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshot", arg0, arg1)
	ret0, _ := ret[0].(*api.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSnapshot indicates an expected call of CreateSnapshot.
func (mr *MockAstraDSMockRecorder) CreateSnapshot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshot", reflect.TypeOf((*MockAstraDS)(nil).CreateSnapshot), arg0, arg1)
}

// CreateVolume mocks base method.
func (m *MockAstraDS) CreateVolume(arg0 context.Context, arg1 *api.Volume) (*api.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolume", arg0, arg1)
	ret0, _ := ret[0].(*api.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVolume indicates an expected call of CreateVolume.
func (mr *MockAstraDSMockRecorder) CreateVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockAstraDS)(nil).CreateVolume), arg0, arg1)
}

// DeleteExportPolicy mocks base method.
func (m *MockAstraDS) DeleteExportPolicy(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExportPolicy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExportPolicy indicates an expected call of DeleteExportPolicy.
func (mr *MockAstraDSMockRecorder) DeleteExportPolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExportPolicy", reflect.TypeOf((*MockAstraDS)(nil).DeleteExportPolicy), arg0, arg1)
}

// DeleteSnapshot mocks base method.
func (m *MockAstraDS) DeleteSnapshot(arg0 context.Context, arg1 *api.Snapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshot", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSnapshot indicates an expected call of DeleteSnapshot.
func (mr *MockAstraDSMockRecorder) DeleteSnapshot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshot", reflect.TypeOf((*MockAstraDS)(nil).DeleteSnapshot), arg0, arg1)
}

// DeleteVolume mocks base method.
func (m *MockAstraDS) DeleteVolume(arg0 context.Context, arg1 *api.Volume) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVolume indicates an expected call of DeleteVolume.
func (mr *MockAstraDSMockRecorder) DeleteVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockAstraDS)(nil).DeleteVolume), arg0, arg1)
}

// EnsureExportPolicyExists mocks base method.
func (m *MockAstraDS) EnsureExportPolicyExists(arg0 context.Context, arg1 string) (*api.ExportPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureExportPolicyExists", arg0, arg1)
	ret0, _ := ret[0].(*api.ExportPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureExportPolicyExists indicates an expected call of EnsureExportPolicyExists.
func (mr *MockAstraDSMockRecorder) EnsureExportPolicyExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureExportPolicyExists", reflect.TypeOf((*MockAstraDS)(nil).EnsureExportPolicyExists), arg0, arg1)
}

// ExportPolicyExists mocks base method.
func (m *MockAstraDS) ExportPolicyExists(arg0 context.Context, arg1 string) (bool, *api.ExportPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportPolicyExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*api.ExportPolicy)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExportPolicyExists indicates an expected call of ExportPolicyExists.
func (mr *MockAstraDSMockRecorder) ExportPolicyExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportPolicyExists", reflect.TypeOf((*MockAstraDS)(nil).ExportPolicyExists), arg0, arg1)
}

// Init mocks base method.
func (m *MockAstraDS) Init(arg0 context.Context, arg1, arg2, arg3 string) (*api.Cluster, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*api.Cluster)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Init indicates an expected call of Init.
func (mr *MockAstraDSMockRecorder) Init(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockAstraDS)(nil).Init), arg0, arg1, arg2, arg3)
}

// QosPolicies mocks base method.
func (m *MockAstraDS) QosPolicies(arg0 context.Context) ([]*api.QosPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QosPolicies", arg0)
	ret0, _ := ret[0].([]*api.QosPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QosPolicies indicates an expected call of QosPolicies.
func (mr *MockAstraDSMockRecorder) QosPolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QosPolicies", reflect.TypeOf((*MockAstraDS)(nil).QosPolicies), arg0)
}

// SetExportPolicyAttributes mocks base method.
func (m *MockAstraDS) SetExportPolicyAttributes(arg0 context.Context, arg1 *api.ExportPolicy, arg2 *roaring.Bitmap) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetExportPolicyAttributes", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetExportPolicyAttributes indicates an expected call of SetExportPolicyAttributes.
func (mr *MockAstraDSMockRecorder) SetExportPolicyAttributes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetExportPolicyAttributes", reflect.TypeOf((*MockAstraDS)(nil).SetExportPolicyAttributes), arg0, arg1, arg2)
}

// SetSnapshotAttributes mocks base method.
func (m *MockAstraDS) SetSnapshotAttributes(arg0 context.Context, arg1 *api.Snapshot, arg2 *roaring.Bitmap) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSnapshotAttributes", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSnapshotAttributes indicates an expected call of SetSnapshotAttributes.
func (mr *MockAstraDSMockRecorder) SetSnapshotAttributes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSnapshotAttributes", reflect.TypeOf((*MockAstraDS)(nil).SetSnapshotAttributes), arg0, arg1, arg2)
}

// SetVolumeAttributes mocks base method.
func (m *MockAstraDS) SetVolumeAttributes(arg0 context.Context, arg1 *api.Volume, arg2 *roaring.Bitmap) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetVolumeAttributes", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetVolumeAttributes indicates an expected call of SetVolumeAttributes.
func (mr *MockAstraDSMockRecorder) SetVolumeAttributes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVolumeAttributes", reflect.TypeOf((*MockAstraDS)(nil).SetVolumeAttributes), arg0, arg1, arg2)
}

// Snapshot mocks base method.
func (m *MockAstraDS) Snapshot(arg0 context.Context, arg1 string) (*api.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot", arg0, arg1)
	ret0, _ := ret[0].(*api.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Snapshot indicates an expected call of Snapshot.
func (mr *MockAstraDSMockRecorder) Snapshot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockAstraDS)(nil).Snapshot), arg0, arg1)
}

// SnapshotExists mocks base method.
func (m *MockAstraDS) SnapshotExists(arg0 context.Context, arg1 string) (bool, *api.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*api.Snapshot)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SnapshotExists indicates an expected call of SnapshotExists.
func (mr *MockAstraDSMockRecorder) SnapshotExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotExists", reflect.TypeOf((*MockAstraDS)(nil).SnapshotExists), arg0, arg1)
}

// Snapshots mocks base method.
func (m *MockAstraDS) Snapshots(arg0 context.Context, arg1 *api.Volume) ([]*api.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshots", arg0, arg1)
	ret0, _ := ret[0].([]*api.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Snapshots indicates an expected call of Snapshots.
func (mr *MockAstraDSMockRecorder) Snapshots(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshots", reflect.TypeOf((*MockAstraDS)(nil).Snapshots), arg0, arg1)
}

// Volume mocks base method.
func (m *MockAstraDS) Volume(arg0 context.Context, arg1 string) (*api.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Volume", arg0, arg1)
	ret0, _ := ret[0].(*api.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Volume indicates an expected call of Volume.
func (mr *MockAstraDSMockRecorder) Volume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Volume", reflect.TypeOf((*MockAstraDS)(nil).Volume), arg0, arg1)
}

// VolumeExists mocks base method.
func (m *MockAstraDS) VolumeExists(arg0 context.Context, arg1 string) (bool, *api.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*api.Volume)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VolumeExists indicates an expected call of VolumeExists.
func (mr *MockAstraDSMockRecorder) VolumeExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeExists", reflect.TypeOf((*MockAstraDS)(nil).VolumeExists), arg0, arg1)
}

// Volumes mocks base method.
func (m *MockAstraDS) Volumes(arg0 context.Context) ([]*api.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Volumes", arg0)
	ret0, _ := ret[0].([]*api.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Volumes indicates an expected call of Volumes.
func (mr *MockAstraDSMockRecorder) Volumes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Volumes", reflect.TypeOf((*MockAstraDS)(nil).Volumes), arg0)
}

// WaitForSnapshotDeleted mocks base method.
func (m *MockAstraDS) WaitForSnapshotDeleted(arg0 context.Context, arg1 *api.Snapshot, arg2 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForSnapshotDeleted", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForSnapshotDeleted indicates an expected call of WaitForSnapshotDeleted.
func (mr *MockAstraDSMockRecorder) WaitForSnapshotDeleted(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForSnapshotDeleted", reflect.TypeOf((*MockAstraDS)(nil).WaitForSnapshotDeleted), arg0, arg1, arg2)
}

// WaitForSnapshotReady mocks base method.
func (m *MockAstraDS) WaitForSnapshotReady(arg0 context.Context, arg1 *api.Snapshot, arg2 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForSnapshotReady", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForSnapshotReady indicates an expected call of WaitForSnapshotReady.
func (mr *MockAstraDSMockRecorder) WaitForSnapshotReady(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForSnapshotReady", reflect.TypeOf((*MockAstraDS)(nil).WaitForSnapshotReady), arg0, arg1, arg2)
}

// WaitForVolumeDeleted mocks base method.
func (m *MockAstraDS) WaitForVolumeDeleted(arg0 context.Context, arg1 *api.Volume, arg2 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForVolumeDeleted", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForVolumeDeleted indicates an expected call of WaitForVolumeDeleted.
func (mr *MockAstraDSMockRecorder) WaitForVolumeDeleted(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForVolumeDeleted", reflect.TypeOf((*MockAstraDS)(nil).WaitForVolumeDeleted), arg0, arg1, arg2)
}

// WaitForVolumeReady mocks base method.
func (m *MockAstraDS) WaitForVolumeReady(arg0 context.Context, arg1 *api.Volume, arg2 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForVolumeReady", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForVolumeReady indicates an expected call of WaitForVolumeReady.
func (mr *MockAstraDSMockRecorder) WaitForVolumeReady(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForVolumeReady", reflect.TypeOf((*MockAstraDS)(nil).WaitForVolumeReady), arg0, arg1, arg2)
}

// WaitForVolumeResize mocks base method.
func (m *MockAstraDS) WaitForVolumeResize(arg0 context.Context, arg1 *api.Volume, arg2 int64, arg3 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForVolumeResize", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForVolumeResize indicates an expected call of WaitForVolumeResize.
func (mr *MockAstraDSMockRecorder) WaitForVolumeResize(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForVolumeResize", reflect.TypeOf((*MockAstraDS)(nil).WaitForVolumeResize), arg0, arg1, arg2, arg3)
}
