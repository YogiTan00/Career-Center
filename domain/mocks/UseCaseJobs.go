// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entity "CareerCenter/domain/entity"
	filter "CareerCenter/domain/entity/filter"
	context "context"

	mock "github.com/stretchr/testify/mock"

	utils "CareerCenter/utils"
)

// UseCaseJobs is an autogenerated mock type for the UseCaseJobs type
type UseCaseJobs struct {
	mock.Mock
}

// CreateJob provides a mock function with given fields: ctx, dto
func (_m *UseCaseJobs) CreateJob(ctx context.Context, dto *entity.JobsDTO) error {
	ret := _m.Called(ctx, dto)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.JobsDTO) error); ok {
		r0 = rf(ctx, dto)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteJob provides a mock function with given fields: ctx, id
func (_m *UseCaseJobs) DeleteJob(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetJobById provides a mock function with given fields: ctx, user, id
func (_m *UseCaseJobs) GetJobById(ctx context.Context, user *utils.User, id string) (*entity.JobsDTO, error) {
	ret := _m.Called(ctx, user, id)

	var r0 *entity.JobsDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *utils.User, string) (*entity.JobsDTO, error)); ok {
		return rf(ctx, user, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *utils.User, string) *entity.JobsDTO); ok {
		r0 = rf(ctx, user, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.JobsDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *utils.User, string) error); ok {
		r1 = rf(ctx, user, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListByEmail provides a mock function with given fields: ctx, user
func (_m *UseCaseJobs) GetListByEmail(ctx context.Context, user *utils.User) ([]*entity.JobsDTO, int, error) {
	ret := _m.Called(ctx, user)

	var r0 []*entity.JobsDTO
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *utils.User) ([]*entity.JobsDTO, int, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *utils.User) []*entity.JobsDTO); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.JobsDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *utils.User) int); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *utils.User) error); ok {
		r2 = rf(ctx, user)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetListJobs provides a mock function with given fields: ctx, f
func (_m *UseCaseJobs) GetListJobs(ctx context.Context, f *filter.FilterDTO) ([]*entity.JobsDTO, int, error) {
	ret := _m.Called(ctx, f)

	var r0 []*entity.JobsDTO
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *filter.FilterDTO) ([]*entity.JobsDTO, int, error)); ok {
		return rf(ctx, f)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *filter.FilterDTO) []*entity.JobsDTO); ok {
		r0 = rf(ctx, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.JobsDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *filter.FilterDTO) int); ok {
		r1 = rf(ctx, f)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *filter.FilterDTO) error); ok {
		r2 = rf(ctx, f)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateJob provides a mock function with given fields: ctx, id, dto
func (_m *UseCaseJobs) UpdateJob(ctx context.Context, id string, dto *entity.JobsDTO) error {
	ret := _m.Called(ctx, id, dto)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *entity.JobsDTO) error); ok {
		r0 = rf(ctx, id, dto)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUseCaseJobs interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCaseJobs creates a new instance of UseCaseJobs. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCaseJobs(t mockConstructorTestingTNewUseCaseJobs) *UseCaseJobs {
	mock := &UseCaseJobs{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
