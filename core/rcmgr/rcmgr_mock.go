// Code generated by MockGen. DO NOT EDIT.
// Source: ./rcmgr.go

// Package rcmgr is a generated GoMock package.
package rcmgr

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockResourceManager is a mock of ResourceManager interface.
type MockResourceManager struct {
	ctrl     *gomock.Controller
	recorder *MockResourceManagerMockRecorder
}

// MockResourceManagerMockRecorder is the mock recorder for MockResourceManager.
type MockResourceManagerMockRecorder struct {
	mock *MockResourceManager
}

// NewMockResourceManager creates a new mock instance.
func NewMockResourceManager(ctrl *gomock.Controller) *MockResourceManager {
	mock := &MockResourceManager{ctrl: ctrl}
	mock.recorder = &MockResourceManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceManager) EXPECT() *MockResourceManagerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockResourceManager) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockResourceManagerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockResourceManager)(nil).Close))
}

// OpenService mocks base method.
func (m *MockResourceManager) OpenService(svc string) (ResourceScope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenService", svc)
	ret0, _ := ret[0].(ResourceScope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenService indicates an expected call of OpenService.
func (mr *MockResourceManagerMockRecorder) OpenService(svc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenService", reflect.TypeOf((*MockResourceManager)(nil).OpenService), svc)
}

// ServiceState mocks base method.
func (m *MockResourceManager) ServiceState(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceState", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// ServiceState indicates an expected call of ServiceState.
func (mr *MockResourceManagerMockRecorder) ServiceState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceState", reflect.TypeOf((*MockResourceManager)(nil).ServiceState), arg0)
}

// SystemState mocks base method.
func (m *MockResourceManager) SystemState() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SystemState")
	ret0, _ := ret[0].(string)
	return ret0
}

// SystemState indicates an expected call of SystemState.
func (mr *MockResourceManagerMockRecorder) SystemState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SystemState", reflect.TypeOf((*MockResourceManager)(nil).SystemState))
}

// TransientState mocks base method.
func (m *MockResourceManager) TransientState() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransientState")
	ret0, _ := ret[0].(string)
	return ret0
}

// TransientState indicates an expected call of TransientState.
func (mr *MockResourceManagerMockRecorder) TransientState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransientState", reflect.TypeOf((*MockResourceManager)(nil).TransientState))
}

// ViewService mocks base method.
func (m *MockResourceManager) ViewService(arg0 string, arg1 func(ResourceScope) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewService", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ViewService indicates an expected call of ViewService.
func (mr *MockResourceManagerMockRecorder) ViewService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewService", reflect.TypeOf((*MockResourceManager)(nil).ViewService), arg0, arg1)
}

// ViewSystem mocks base method.
func (m *MockResourceManager) ViewSystem(arg0 func(ResourceScope) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewSystem", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ViewSystem indicates an expected call of ViewSystem.
func (mr *MockResourceManagerMockRecorder) ViewSystem(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewSystem", reflect.TypeOf((*MockResourceManager)(nil).ViewSystem), arg0)
}

// ViewTransient mocks base method.
func (m *MockResourceManager) ViewTransient(arg0 func(ResourceScope) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewTransient", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ViewTransient indicates an expected call of ViewTransient.
func (mr *MockResourceManagerMockRecorder) ViewTransient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewTransient", reflect.TypeOf((*MockResourceManager)(nil).ViewTransient), arg0)
}

// MockResourceScopeViewer is a mock of ResourceScopeViewer interface.
type MockResourceScopeViewer struct {
	ctrl     *gomock.Controller
	recorder *MockResourceScopeViewerMockRecorder
}

// MockResourceScopeViewerMockRecorder is the mock recorder for MockResourceScopeViewer.
type MockResourceScopeViewerMockRecorder struct {
	mock *MockResourceScopeViewer
}

// NewMockResourceScopeViewer creates a new mock instance.
func NewMockResourceScopeViewer(ctrl *gomock.Controller) *MockResourceScopeViewer {
	mock := &MockResourceScopeViewer{ctrl: ctrl}
	mock.recorder = &MockResourceScopeViewerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceScopeViewer) EXPECT() *MockResourceScopeViewerMockRecorder {
	return m.recorder
}

// ServiceState mocks base method.
func (m *MockResourceScopeViewer) ServiceState(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceState", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// ServiceState indicates an expected call of ServiceState.
func (mr *MockResourceScopeViewerMockRecorder) ServiceState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceState", reflect.TypeOf((*MockResourceScopeViewer)(nil).ServiceState), arg0)
}

// SystemState mocks base method.
func (m *MockResourceScopeViewer) SystemState() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SystemState")
	ret0, _ := ret[0].(string)
	return ret0
}

// SystemState indicates an expected call of SystemState.
func (mr *MockResourceScopeViewerMockRecorder) SystemState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SystemState", reflect.TypeOf((*MockResourceScopeViewer)(nil).SystemState))
}

// TransientState mocks base method.
func (m *MockResourceScopeViewer) TransientState() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransientState")
	ret0, _ := ret[0].(string)
	return ret0
}

// TransientState indicates an expected call of TransientState.
func (mr *MockResourceScopeViewerMockRecorder) TransientState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransientState", reflect.TypeOf((*MockResourceScopeViewer)(nil).TransientState))
}

// ViewService mocks base method.
func (m *MockResourceScopeViewer) ViewService(arg0 string, arg1 func(ResourceScope) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewService", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ViewService indicates an expected call of ViewService.
func (mr *MockResourceScopeViewerMockRecorder) ViewService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewService", reflect.TypeOf((*MockResourceScopeViewer)(nil).ViewService), arg0, arg1)
}

// ViewSystem mocks base method.
func (m *MockResourceScopeViewer) ViewSystem(arg0 func(ResourceScope) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewSystem", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ViewSystem indicates an expected call of ViewSystem.
func (mr *MockResourceScopeViewerMockRecorder) ViewSystem(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewSystem", reflect.TypeOf((*MockResourceScopeViewer)(nil).ViewSystem), arg0)
}

// ViewTransient mocks base method.
func (m *MockResourceScopeViewer) ViewTransient(arg0 func(ResourceScope) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewTransient", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ViewTransient indicates an expected call of ViewTransient.
func (mr *MockResourceScopeViewerMockRecorder) ViewTransient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewTransient", reflect.TypeOf((*MockResourceScopeViewer)(nil).ViewTransient), arg0)
}

// MockResourceScope is a mock of ResourceScope interface.
type MockResourceScope struct {
	ctrl     *gomock.Controller
	recorder *MockResourceScopeMockRecorder
}

// MockResourceScopeMockRecorder is the mock recorder for MockResourceScope.
type MockResourceScopeMockRecorder struct {
	mock *MockResourceScope
}

// NewMockResourceScope creates a new mock instance.
func NewMockResourceScope(ctrl *gomock.Controller) *MockResourceScope {
	mock := &MockResourceScope{ctrl: ctrl}
	mock.recorder = &MockResourceScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceScope) EXPECT() *MockResourceScopeMockRecorder {
	return m.recorder
}

// AddConn mocks base method.
func (m *MockResourceScope) AddConn(dir Direction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddConn", dir)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddConn indicates an expected call of AddConn.
func (mr *MockResourceScopeMockRecorder) AddConn(dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddConn", reflect.TypeOf((*MockResourceScope)(nil).AddConn), dir)
}

// AddTask mocks base method.
func (m *MockResourceScope) AddTask(num int, prio ReserveTaskPriority) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTask", num, prio)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTask indicates an expected call of AddTask.
func (mr *MockResourceScopeMockRecorder) AddTask(num, prio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTask", reflect.TypeOf((*MockResourceScope)(nil).AddTask), num, prio)
}

// BeginSpan mocks base method.
func (m *MockResourceScope) BeginSpan() (ResourceScopeSpan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginSpan")
	ret0, _ := ret[0].(ResourceScopeSpan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginSpan indicates an expected call of BeginSpan.
func (mr *MockResourceScopeMockRecorder) BeginSpan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginSpan", reflect.TypeOf((*MockResourceScope)(nil).BeginSpan))
}

// Name mocks base method.
func (m *MockResourceScope) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockResourceScopeMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockResourceScope)(nil).Name))
}

// Release mocks base method.
func (m *MockResourceScope) Release() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Release")
}

// Release indicates an expected call of Release.
func (mr *MockResourceScopeMockRecorder) Release() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Release", reflect.TypeOf((*MockResourceScope)(nil).Release))
}

// ReleaseMemory mocks base method.
func (m *MockResourceScope) ReleaseMemory(size int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReleaseMemory", size)
}

// ReleaseMemory indicates an expected call of ReleaseMemory.
func (mr *MockResourceScopeMockRecorder) ReleaseMemory(size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseMemory", reflect.TypeOf((*MockResourceScope)(nil).ReleaseMemory), size)
}

// RemainingResource mocks base method.
func (m *MockResourceScope) RemainingResource() (Limit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemainingResource")
	ret0, _ := ret[0].(Limit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemainingResource indicates an expected call of RemainingResource.
func (mr *MockResourceScopeMockRecorder) RemainingResource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemainingResource", reflect.TypeOf((*MockResourceScope)(nil).RemainingResource))
}

// RemoveConn mocks base method.
func (m *MockResourceScope) RemoveConn(dir Direction) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveConn", dir)
}

// RemoveConn indicates an expected call of RemoveConn.
func (mr *MockResourceScopeMockRecorder) RemoveConn(dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveConn", reflect.TypeOf((*MockResourceScope)(nil).RemoveConn), dir)
}

// RemoveTask mocks base method.
func (m *MockResourceScope) RemoveTask(num int, prio ReserveTaskPriority) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveTask", num, prio)
}

// RemoveTask indicates an expected call of RemoveTask.
func (mr *MockResourceScopeMockRecorder) RemoveTask(num, prio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTask", reflect.TypeOf((*MockResourceScope)(nil).RemoveTask), num, prio)
}

// ReserveMemory mocks base method.
func (m *MockResourceScope) ReserveMemory(size int64, prio uint8) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveMemory", size, prio)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReserveMemory indicates an expected call of ReserveMemory.
func (mr *MockResourceScopeMockRecorder) ReserveMemory(size, prio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReserveMemory", reflect.TypeOf((*MockResourceScope)(nil).ReserveMemory), size, prio)
}

// ReserveResources mocks base method.
func (m *MockResourceScope) ReserveResources(st *ScopeStat) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveResources", st)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReserveResources indicates an expected call of ReserveResources.
func (mr *MockResourceScopeMockRecorder) ReserveResources(st interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReserveResources", reflect.TypeOf((*MockResourceScope)(nil).ReserveResources), st)
}

// Stat mocks base method.
func (m *MockResourceScope) Stat() ScopeStat {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat")
	ret0, _ := ret[0].(ScopeStat)
	return ret0
}

// Stat indicates an expected call of Stat.
func (mr *MockResourceScopeMockRecorder) Stat() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockResourceScope)(nil).Stat))
}

// MockResourceScopeSpan is a mock of ResourceScopeSpan interface.
type MockResourceScopeSpan struct {
	ctrl     *gomock.Controller
	recorder *MockResourceScopeSpanMockRecorder
}

// MockResourceScopeSpanMockRecorder is the mock recorder for MockResourceScopeSpan.
type MockResourceScopeSpanMockRecorder struct {
	mock *MockResourceScopeSpan
}

// NewMockResourceScopeSpan creates a new mock instance.
func NewMockResourceScopeSpan(ctrl *gomock.Controller) *MockResourceScopeSpan {
	mock := &MockResourceScopeSpan{ctrl: ctrl}
	mock.recorder = &MockResourceScopeSpanMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceScopeSpan) EXPECT() *MockResourceScopeSpanMockRecorder {
	return m.recorder
}

// AddConn mocks base method.
func (m *MockResourceScopeSpan) AddConn(dir Direction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddConn", dir)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddConn indicates an expected call of AddConn.
func (mr *MockResourceScopeSpanMockRecorder) AddConn(dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddConn", reflect.TypeOf((*MockResourceScopeSpan)(nil).AddConn), dir)
}

// AddTask mocks base method.
func (m *MockResourceScopeSpan) AddTask(num int, prio ReserveTaskPriority) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTask", num, prio)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTask indicates an expected call of AddTask.
func (mr *MockResourceScopeSpanMockRecorder) AddTask(num, prio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTask", reflect.TypeOf((*MockResourceScopeSpan)(nil).AddTask), num, prio)
}

// BeginSpan mocks base method.
func (m *MockResourceScopeSpan) BeginSpan() (ResourceScopeSpan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginSpan")
	ret0, _ := ret[0].(ResourceScopeSpan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginSpan indicates an expected call of BeginSpan.
func (mr *MockResourceScopeSpanMockRecorder) BeginSpan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginSpan", reflect.TypeOf((*MockResourceScopeSpan)(nil).BeginSpan))
}

// Done mocks base method.
func (m *MockResourceScopeSpan) Done() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Done")
}

// Done indicates an expected call of Done.
func (mr *MockResourceScopeSpanMockRecorder) Done() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockResourceScopeSpan)(nil).Done))
}

// Name mocks base method.
func (m *MockResourceScopeSpan) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockResourceScopeSpanMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockResourceScopeSpan)(nil).Name))
}

// Release mocks base method.
func (m *MockResourceScopeSpan) Release() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Release")
}

// Release indicates an expected call of Release.
func (mr *MockResourceScopeSpanMockRecorder) Release() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Release", reflect.TypeOf((*MockResourceScopeSpan)(nil).Release))
}

// ReleaseMemory mocks base method.
func (m *MockResourceScopeSpan) ReleaseMemory(size int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReleaseMemory", size)
}

// ReleaseMemory indicates an expected call of ReleaseMemory.
func (mr *MockResourceScopeSpanMockRecorder) ReleaseMemory(size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseMemory", reflect.TypeOf((*MockResourceScopeSpan)(nil).ReleaseMemory), size)
}

// RemainingResource mocks base method.
func (m *MockResourceScopeSpan) RemainingResource() (Limit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemainingResource")
	ret0, _ := ret[0].(Limit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemainingResource indicates an expected call of RemainingResource.
func (mr *MockResourceScopeSpanMockRecorder) RemainingResource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemainingResource", reflect.TypeOf((*MockResourceScopeSpan)(nil).RemainingResource))
}

// RemoveConn mocks base method.
func (m *MockResourceScopeSpan) RemoveConn(dir Direction) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveConn", dir)
}

// RemoveConn indicates an expected call of RemoveConn.
func (mr *MockResourceScopeSpanMockRecorder) RemoveConn(dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveConn", reflect.TypeOf((*MockResourceScopeSpan)(nil).RemoveConn), dir)
}

// RemoveTask mocks base method.
func (m *MockResourceScopeSpan) RemoveTask(num int, prio ReserveTaskPriority) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveTask", num, prio)
}

// RemoveTask indicates an expected call of RemoveTask.
func (mr *MockResourceScopeSpanMockRecorder) RemoveTask(num, prio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTask", reflect.TypeOf((*MockResourceScopeSpan)(nil).RemoveTask), num, prio)
}

// ReserveMemory mocks base method.
func (m *MockResourceScopeSpan) ReserveMemory(size int64, prio uint8) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveMemory", size, prio)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReserveMemory indicates an expected call of ReserveMemory.
func (mr *MockResourceScopeSpanMockRecorder) ReserveMemory(size, prio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReserveMemory", reflect.TypeOf((*MockResourceScopeSpan)(nil).ReserveMemory), size, prio)
}

// ReserveResources mocks base method.
func (m *MockResourceScopeSpan) ReserveResources(st *ScopeStat) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveResources", st)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReserveResources indicates an expected call of ReserveResources.
func (mr *MockResourceScopeSpanMockRecorder) ReserveResources(st interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReserveResources", reflect.TypeOf((*MockResourceScopeSpan)(nil).ReserveResources), st)
}

// Stat mocks base method.
func (m *MockResourceScopeSpan) Stat() ScopeStat {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat")
	ret0, _ := ret[0].(ScopeStat)
	return ret0
}

// Stat indicates an expected call of Stat.
func (mr *MockResourceScopeSpanMockRecorder) Stat() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockResourceScopeSpan)(nil).Stat))
}
