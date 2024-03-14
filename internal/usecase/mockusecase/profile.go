// Code generated by MockGen. DO NOT EDIT.
// Source: profile.go
//
// Generated by this command:
//
//	mockgen -source=profile.go -destination=mockusecase/profile.go -package=mockusecase
//

// Package mockusecase is a generated GoMock package.
package mockusecase

import (
	context "context"
	reflect "reflect"

	dto "github.com/Hidayathamir/go-user/internal/dto"
	gomock "go.uber.org/mock/gomock"
)

// MockIProfile is a mock of IProfile interface.
type MockIProfile struct {
	ctrl     *gomock.Controller
	recorder *MockIProfileMockRecorder
}

// MockIProfileMockRecorder is the mock recorder for MockIProfile.
type MockIProfileMockRecorder struct {
	mock *MockIProfile
}

// NewMockIProfile creates a new mock instance.
func NewMockIProfile(ctrl *gomock.Controller) *MockIProfile {
	mock := &MockIProfile{ctrl: ctrl}
	mock.recorder = &MockIProfileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProfile) EXPECT() *MockIProfileMockRecorder {
	return m.recorder
}

// GetProfileByUsername mocks base method.
func (m *MockIProfile) GetProfileByUsername(ctx context.Context, username string) (dto.ResGetProfileByUsername, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileByUsername", ctx, username)
	ret0, _ := ret[0].(dto.ResGetProfileByUsername)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileByUsername indicates an expected call of GetProfileByUsername.
func (mr *MockIProfileMockRecorder) GetProfileByUsername(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileByUsername", reflect.TypeOf((*MockIProfile)(nil).GetProfileByUsername), ctx, username)
}

// UpdateProfileByUserID mocks base method.
func (m *MockIProfile) UpdateProfileByUserID(ctx context.Context, req dto.ReqUpdateProfileByUserID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfileByUserID", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfileByUserID indicates an expected call of UpdateProfileByUserID.
func (mr *MockIProfileMockRecorder) UpdateProfileByUserID(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfileByUserID", reflect.TypeOf((*MockIProfile)(nil).UpdateProfileByUserID), ctx, req)
}
