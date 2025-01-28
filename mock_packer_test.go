// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go (interfaces: Packer)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/quic-go/quic-go -destination mock_packer_test.go github.com/quic-go/quic-go Packer
//

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"
	time "time"

	ackhandler "github.com/quic-go/quic-go/internal/ackhandler"
	protocol "github.com/quic-go/quic-go/internal/protocol"
	qerr "github.com/quic-go/quic-go/internal/qerr"
	gomock "go.uber.org/mock/gomock"
)

// MockPacker is a mock of Packer interface.
type MockPacker struct {
	ctrl     *gomock.Controller
	recorder *MockPackerMockRecorder
	isgomock struct{}
}

// MockPackerMockRecorder is the mock recorder for MockPacker.
type MockPackerMockRecorder struct {
	mock *MockPacker
}

// NewMockPacker creates a new mock instance.
func NewMockPacker(ctrl *gomock.Controller) *MockPacker {
	mock := &MockPacker{ctrl: ctrl}
	mock.recorder = &MockPackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPacker) EXPECT() *MockPackerMockRecorder {
	return m.recorder
}

// AppendPacket mocks base method.
func (m *MockPacker) AppendPacket(buf *packetBuffer, maxPacketSize protocol.ByteCount, now time.Time, v protocol.Version) (shortHeaderPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendPacket", buf, maxPacketSize, now, v)
	ret0, _ := ret[0].(shortHeaderPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppendPacket indicates an expected call of AppendPacket.
func (mr *MockPackerMockRecorder) AppendPacket(buf, maxPacketSize, now, v any) *MockPackerAppendPacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendPacket", reflect.TypeOf((*MockPacker)(nil).AppendPacket), buf, maxPacketSize, now, v)
	return &MockPackerAppendPacketCall{Call: call}
}

// MockPackerAppendPacketCall wrap *gomock.Call
type MockPackerAppendPacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPackerAppendPacketCall) Return(arg0 shortHeaderPacket, arg1 error) *MockPackerAppendPacketCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPackerAppendPacketCall) Do(f func(*packetBuffer, protocol.ByteCount, time.Time, protocol.Version) (shortHeaderPacket, error)) *MockPackerAppendPacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPackerAppendPacketCall) DoAndReturn(f func(*packetBuffer, protocol.ByteCount, time.Time, protocol.Version) (shortHeaderPacket, error)) *MockPackerAppendPacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MaybePackPTOProbePacket mocks base method.
func (m *MockPacker) MaybePackPTOProbePacket(arg0 protocol.EncryptionLevel, arg1 protocol.ByteCount, arg2 time.Time, arg3 protocol.Version) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaybePackPTOProbePacket", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaybePackPTOProbePacket indicates an expected call of MaybePackPTOProbePacket.
func (mr *MockPackerMockRecorder) MaybePackPTOProbePacket(arg0, arg1, arg2, arg3 any) *MockPackerMaybePackPTOProbePacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybePackPTOProbePacket", reflect.TypeOf((*MockPacker)(nil).MaybePackPTOProbePacket), arg0, arg1, arg2, arg3)
	return &MockPackerMaybePackPTOProbePacketCall{Call: call}
}

// MockPackerMaybePackPTOProbePacketCall wrap *gomock.Call
type MockPackerMaybePackPTOProbePacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPackerMaybePackPTOProbePacketCall) Return(arg0 *coalescedPacket, arg1 error) *MockPackerMaybePackPTOProbePacketCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPackerMaybePackPTOProbePacketCall) Do(f func(protocol.EncryptionLevel, protocol.ByteCount, time.Time, protocol.Version) (*coalescedPacket, error)) *MockPackerMaybePackPTOProbePacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPackerMaybePackPTOProbePacketCall) DoAndReturn(f func(protocol.EncryptionLevel, protocol.ByteCount, time.Time, protocol.Version) (*coalescedPacket, error)) *MockPackerMaybePackPTOProbePacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PackAckOnlyPacket mocks base method.
func (m *MockPacker) PackAckOnlyPacket(maxPacketSize protocol.ByteCount, now time.Time, v protocol.Version) (shortHeaderPacket, *packetBuffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackAckOnlyPacket", maxPacketSize, now, v)
	ret0, _ := ret[0].(shortHeaderPacket)
	ret1, _ := ret[1].(*packetBuffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PackAckOnlyPacket indicates an expected call of PackAckOnlyPacket.
func (mr *MockPackerMockRecorder) PackAckOnlyPacket(maxPacketSize, now, v any) *MockPackerPackAckOnlyPacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackAckOnlyPacket", reflect.TypeOf((*MockPacker)(nil).PackAckOnlyPacket), maxPacketSize, now, v)
	return &MockPackerPackAckOnlyPacketCall{Call: call}
}

// MockPackerPackAckOnlyPacketCall wrap *gomock.Call
type MockPackerPackAckOnlyPacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPackerPackAckOnlyPacketCall) Return(arg0 shortHeaderPacket, arg1 *packetBuffer, arg2 error) *MockPackerPackAckOnlyPacketCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPackerPackAckOnlyPacketCall) Do(f func(protocol.ByteCount, time.Time, protocol.Version) (shortHeaderPacket, *packetBuffer, error)) *MockPackerPackAckOnlyPacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPackerPackAckOnlyPacketCall) DoAndReturn(f func(protocol.ByteCount, time.Time, protocol.Version) (shortHeaderPacket, *packetBuffer, error)) *MockPackerPackAckOnlyPacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PackApplicationClose mocks base method.
func (m *MockPacker) PackApplicationClose(arg0 *qerr.ApplicationError, arg1 protocol.ByteCount, arg2 protocol.Version) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackApplicationClose", arg0, arg1, arg2)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PackApplicationClose indicates an expected call of PackApplicationClose.
func (mr *MockPackerMockRecorder) PackApplicationClose(arg0, arg1, arg2 any) *MockPackerPackApplicationCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackApplicationClose", reflect.TypeOf((*MockPacker)(nil).PackApplicationClose), arg0, arg1, arg2)
	return &MockPackerPackApplicationCloseCall{Call: call}
}

// MockPackerPackApplicationCloseCall wrap *gomock.Call
type MockPackerPackApplicationCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPackerPackApplicationCloseCall) Return(arg0 *coalescedPacket, arg1 error) *MockPackerPackApplicationCloseCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPackerPackApplicationCloseCall) Do(f func(*qerr.ApplicationError, protocol.ByteCount, protocol.Version) (*coalescedPacket, error)) *MockPackerPackApplicationCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPackerPackApplicationCloseCall) DoAndReturn(f func(*qerr.ApplicationError, protocol.ByteCount, protocol.Version) (*coalescedPacket, error)) *MockPackerPackApplicationCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PackCoalescedPacket mocks base method.
func (m *MockPacker) PackCoalescedPacket(onlyAck bool, maxPacketSize protocol.ByteCount, now time.Time, v protocol.Version) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackCoalescedPacket", onlyAck, maxPacketSize, now, v)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PackCoalescedPacket indicates an expected call of PackCoalescedPacket.
func (mr *MockPackerMockRecorder) PackCoalescedPacket(onlyAck, maxPacketSize, now, v any) *MockPackerPackCoalescedPacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackCoalescedPacket", reflect.TypeOf((*MockPacker)(nil).PackCoalescedPacket), onlyAck, maxPacketSize, now, v)
	return &MockPackerPackCoalescedPacketCall{Call: call}
}

// MockPackerPackCoalescedPacketCall wrap *gomock.Call
type MockPackerPackCoalescedPacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPackerPackCoalescedPacketCall) Return(arg0 *coalescedPacket, arg1 error) *MockPackerPackCoalescedPacketCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPackerPackCoalescedPacketCall) Do(f func(bool, protocol.ByteCount, time.Time, protocol.Version) (*coalescedPacket, error)) *MockPackerPackCoalescedPacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPackerPackCoalescedPacketCall) DoAndReturn(f func(bool, protocol.ByteCount, time.Time, protocol.Version) (*coalescedPacket, error)) *MockPackerPackCoalescedPacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PackConnectionClose mocks base method.
func (m *MockPacker) PackConnectionClose(arg0 *qerr.TransportError, arg1 protocol.ByteCount, arg2 protocol.Version) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackConnectionClose", arg0, arg1, arg2)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PackConnectionClose indicates an expected call of PackConnectionClose.
func (mr *MockPackerMockRecorder) PackConnectionClose(arg0, arg1, arg2 any) *MockPackerPackConnectionCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackConnectionClose", reflect.TypeOf((*MockPacker)(nil).PackConnectionClose), arg0, arg1, arg2)
	return &MockPackerPackConnectionCloseCall{Call: call}
}

// MockPackerPackConnectionCloseCall wrap *gomock.Call
type MockPackerPackConnectionCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPackerPackConnectionCloseCall) Return(arg0 *coalescedPacket, arg1 error) *MockPackerPackConnectionCloseCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPackerPackConnectionCloseCall) Do(f func(*qerr.TransportError, protocol.ByteCount, protocol.Version) (*coalescedPacket, error)) *MockPackerPackConnectionCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPackerPackConnectionCloseCall) DoAndReturn(f func(*qerr.TransportError, protocol.ByteCount, protocol.Version) (*coalescedPacket, error)) *MockPackerPackConnectionCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PackMTUProbePacket mocks base method.
func (m *MockPacker) PackMTUProbePacket(ping ackhandler.Frame, size protocol.ByteCount, v protocol.Version) (shortHeaderPacket, *packetBuffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackMTUProbePacket", ping, size, v)
	ret0, _ := ret[0].(shortHeaderPacket)
	ret1, _ := ret[1].(*packetBuffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PackMTUProbePacket indicates an expected call of PackMTUProbePacket.
func (mr *MockPackerMockRecorder) PackMTUProbePacket(ping, size, v any) *MockPackerPackMTUProbePacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackMTUProbePacket", reflect.TypeOf((*MockPacker)(nil).PackMTUProbePacket), ping, size, v)
	return &MockPackerPackMTUProbePacketCall{Call: call}
}

// MockPackerPackMTUProbePacketCall wrap *gomock.Call
type MockPackerPackMTUProbePacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPackerPackMTUProbePacketCall) Return(arg0 shortHeaderPacket, arg1 *packetBuffer, arg2 error) *MockPackerPackMTUProbePacketCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPackerPackMTUProbePacketCall) Do(f func(ackhandler.Frame, protocol.ByteCount, protocol.Version) (shortHeaderPacket, *packetBuffer, error)) *MockPackerPackMTUProbePacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPackerPackMTUProbePacketCall) DoAndReturn(f func(ackhandler.Frame, protocol.ByteCount, protocol.Version) (shortHeaderPacket, *packetBuffer, error)) *MockPackerPackMTUProbePacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PackPathProbePacket mocks base method.
func (m *MockPacker) PackPathProbePacket(arg0 protocol.ConnectionID, arg1 ackhandler.Frame, arg2 protocol.Version) (shortHeaderPacket, *packetBuffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackPathProbePacket", arg0, arg1, arg2)
	ret0, _ := ret[0].(shortHeaderPacket)
	ret1, _ := ret[1].(*packetBuffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PackPathProbePacket indicates an expected call of PackPathProbePacket.
func (mr *MockPackerMockRecorder) PackPathProbePacket(arg0, arg1, arg2 any) *MockPackerPackPathProbePacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackPathProbePacket", reflect.TypeOf((*MockPacker)(nil).PackPathProbePacket), arg0, arg1, arg2)
	return &MockPackerPackPathProbePacketCall{Call: call}
}

// MockPackerPackPathProbePacketCall wrap *gomock.Call
type MockPackerPackPathProbePacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPackerPackPathProbePacketCall) Return(arg0 shortHeaderPacket, arg1 *packetBuffer, arg2 error) *MockPackerPackPathProbePacketCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPackerPackPathProbePacketCall) Do(f func(protocol.ConnectionID, ackhandler.Frame, protocol.Version) (shortHeaderPacket, *packetBuffer, error)) *MockPackerPackPathProbePacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPackerPackPathProbePacketCall) DoAndReturn(f func(protocol.ConnectionID, ackhandler.Frame, protocol.Version) (shortHeaderPacket, *packetBuffer, error)) *MockPackerPackPathProbePacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetToken mocks base method.
func (m *MockPacker) SetToken(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetToken", arg0)
}

// SetToken indicates an expected call of SetToken.
func (mr *MockPackerMockRecorder) SetToken(arg0 any) *MockPackerSetTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetToken", reflect.TypeOf((*MockPacker)(nil).SetToken), arg0)
	return &MockPackerSetTokenCall{Call: call}
}

// MockPackerSetTokenCall wrap *gomock.Call
type MockPackerSetTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPackerSetTokenCall) Return() *MockPackerSetTokenCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPackerSetTokenCall) Do(f func([]byte)) *MockPackerSetTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPackerSetTokenCall) DoAndReturn(f func([]byte)) *MockPackerSetTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
