// Code generated by MockGen. DO NOT EDIT.
// Source: thumbup.go

// Package client is a generated GoMock package.
package client

import (
	context "context"
	model "go-common/app/service/main/thumbup/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockThumbupRPC is a mock of ThumbupRPC interface
type MockThumbupRPC struct {
	ctrl     *gomock.Controller
	recorder *MockThumbupRPCMockRecorder
}

// MockThumbupRPCMockRecorder is the mock recorder for MockThumbupRPC
type MockThumbupRPCMockRecorder struct {
	mock *MockThumbupRPC
}

// NewMockThumbupRPC creates a new mock instance
func NewMockThumbupRPC(ctrl *gomock.Controller) *MockThumbupRPC {
	mock := &MockThumbupRPC{ctrl: ctrl}
	mock.recorder = &MockThumbupRPCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockThumbupRPC) EXPECT() *MockThumbupRPCMockRecorder {
	return m.recorder
}

// Like mocks base method
func (m *MockThumbupRPC) Like(c context.Context, arg *model.ArgLike) error {
	ret := m.ctrl.Call(m, "Like", c, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Like indicates an expected call of Like
func (mr *MockThumbupRPCMockRecorder) Like(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Like", reflect.TypeOf((*MockThumbupRPC)(nil).Like), c, arg)
}

// HasLike mocks base method
func (m *MockThumbupRPC) HasLike(c context.Context, arg *model.ArgHasLike) (map[int64]int8, error) {
	ret := m.ctrl.Call(m, "HasLike", c, arg)
	ret0, _ := ret[0].(map[int64]int8)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasLike indicates an expected call of HasLike
func (mr *MockThumbupRPCMockRecorder) HasLike(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasLike", reflect.TypeOf((*MockThumbupRPC)(nil).HasLike), c, arg)
}

// Stats mocks base method
func (m *MockThumbupRPC) Stats(c context.Context, arg *model.ArgStats) (map[int64]*model.Stats, error) {
	ret := m.ctrl.Call(m, "Stats", c, arg)
	ret0, _ := ret[0].(map[int64]*model.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stats indicates an expected call of Stats
func (mr *MockThumbupRPCMockRecorder) Stats(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockThumbupRPC)(nil).Stats), c, arg)
}

// UserLikes mocks base method
func (m *MockThumbupRPC) UserLikes(c context.Context, arg *model.ArgUserLikes) ([]*model.ItemLikeRecord, error) {
	ret := m.ctrl.Call(m, "UserLikes", c, arg)
	ret0, _ := ret[0].([]*model.ItemLikeRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserLikes indicates an expected call of UserLikes
func (mr *MockThumbupRPCMockRecorder) UserLikes(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLikes", reflect.TypeOf((*MockThumbupRPC)(nil).UserLikes), c, arg)
}

// UserDislikes mocks base method
func (m *MockThumbupRPC) UserDislikes(c context.Context, arg *model.ArgUserLikes) ([]*model.ItemLikeRecord, error) {
	ret := m.ctrl.Call(m, "UserDislikes", c, arg)
	ret0, _ := ret[0].([]*model.ItemLikeRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserDislikes indicates an expected call of UserDislikes
func (mr *MockThumbupRPCMockRecorder) UserDislikes(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserDislikes", reflect.TypeOf((*MockThumbupRPC)(nil).UserDislikes), c, arg)
}

// ItemLikes mocks base method
func (m *MockThumbupRPC) ItemLikes(c context.Context, arg *model.ArgItemLikes) ([]*model.UserLikeRecord, error) {
	ret := m.ctrl.Call(m, "ItemLikes", c, arg)
	ret0, _ := ret[0].([]*model.UserLikeRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ItemLikes indicates an expected call of ItemLikes
func (mr *MockThumbupRPCMockRecorder) ItemLikes(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ItemLikes", reflect.TypeOf((*MockThumbupRPC)(nil).ItemLikes), c, arg)
}

// ItemDislikes mocks base method
func (m *MockThumbupRPC) ItemDislikes(c context.Context, arg *model.ArgItemLikes) ([]*model.UserLikeRecord, error) {
	ret := m.ctrl.Call(m, "ItemDislikes", c, arg)
	ret0, _ := ret[0].([]*model.UserLikeRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ItemDislikes indicates an expected call of ItemDislikes
func (mr *MockThumbupRPCMockRecorder) ItemDislikes(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ItemDislikes", reflect.TypeOf((*MockThumbupRPC)(nil).ItemDislikes), c, arg)
}

// StatsWithLike mocks base method
func (m *MockThumbupRPC) StatsWithLike(c context.Context, arg *model.ArgStatsWithLike) (map[int64]*model.StatsWithLike, error) {
	ret := m.ctrl.Call(m, "StatsWithLike", c, arg)
	ret0, _ := ret[0].(map[int64]*model.StatsWithLike)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatsWithLike indicates an expected call of StatsWithLike
func (mr *MockThumbupRPCMockRecorder) StatsWithLike(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatsWithLike", reflect.TypeOf((*MockThumbupRPC)(nil).StatsWithLike), c, arg)
}

// UserTotalLike mocks base method
func (m *MockThumbupRPC) UserTotalLike(c context.Context, arg *model.ArgUserLikes) (*model.UserTotalLike, error) {
	ret := m.ctrl.Call(m, "UserTotalLike", c, arg)
	ret0, _ := ret[0].(*model.UserTotalLike)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserTotalLike indicates an expected call of UserTotalLike
func (mr *MockThumbupRPCMockRecorder) UserTotalLike(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserTotalLike", reflect.TypeOf((*MockThumbupRPC)(nil).UserTotalLike), c, arg)
}
