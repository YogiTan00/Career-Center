// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entity "CareerCenter/domain/entity"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// RepoApplication is an autogenerated mock type for the RepoApplication type
type RepoApplication struct {
	mock.Mock
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *RepoApplication) GetByEmail(ctx context.Context, email string) (*entity.ApplicationDTO, error) {
	ret := _m.Called(ctx, email)

	var r0 *entity.ApplicationDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.ApplicationDTO, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.ApplicationDTO); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.ApplicationDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByJobId provides a mock function with given fields: ctx, id
func (_m *RepoApplication) GetByJobId(ctx context.Context, id string) ([]*entity.ApplicationDTO, error) {
	ret := _m.Called(ctx, id)

	var r0 []*entity.ApplicationDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*entity.ApplicationDTO, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*entity.ApplicationDTO); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.ApplicationDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListApplication provides a mock function with given fields: ctx
func (_m *RepoApplication) GetListApplication(ctx context.Context) ([]*entity.ApplicationDTO, error) {
	ret := _m.Called(ctx)

	var r0 []*entity.ApplicationDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*entity.ApplicationDTO, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.ApplicationDTO); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.ApplicationDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListByEmail provides a mock function with given fields: ctx, email
func (_m *RepoApplication) GetListByEmail(ctx context.Context, email string) ([]*entity.ApplicationDTO, error) {
	ret := _m.Called(ctx, email)

	var r0 []*entity.ApplicationDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*entity.ApplicationDTO, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*entity.ApplicationDTO); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.ApplicationDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendApplication provides a mock function with given fields: ctx, application
func (_m *RepoApplication) SendApplication(ctx context.Context, application *entity.Application) error {
	ret := _m.Called(ctx, application)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Application) error); ok {
		r0 = rf(ctx, application)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepoApplication interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepoApplication creates a new instance of RepoApplication. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepoApplication(t mockConstructorTestingTNewRepoApplication) *RepoApplication {
	mock := &RepoApplication{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
