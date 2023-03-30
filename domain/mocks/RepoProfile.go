// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	profile "CareerCenter/domain/entity/profile"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// RepoProfile is an autogenerated mock type for the RepoProfile type
type RepoProfile struct {
	mock.Mock
}

// CreateProfile provides a mock function with given fields: ctx, data
func (_m *RepoProfile) CreateProfile(ctx context.Context, data *profile.ProfileUser) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *profile.ProfileUser) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetProfileByEmail provides a mock function with given fields: ctx, email
func (_m *RepoProfile) GetProfileByEmail(ctx context.Context, email string) (*profile.ProfileUserDTO, error) {
	ret := _m.Called(ctx, email)

	var r0 *profile.ProfileUserDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*profile.ProfileUserDTO, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *profile.ProfileUserDTO); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*profile.ProfileUserDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProfile provides a mock function with given fields: ctx, email, data
func (_m *RepoProfile) UpdateProfile(ctx context.Context, email string, data *profile.ProfileUser) error {
	ret := _m.Called(ctx, email, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *profile.ProfileUser) error); ok {
		r0 = rf(ctx, email, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepoProfile interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepoProfile creates a new instance of RepoProfile. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepoProfile(t mockConstructorTestingTNewRepoProfile) *RepoProfile {
	mock := &RepoProfile{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
