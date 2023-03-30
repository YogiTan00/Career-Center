// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entity "CareerCenter/domain/entity"
	filter "CareerCenter/domain/entity/filter"
	context "context"

	mock "github.com/stretchr/testify/mock"

	valueobject "CareerCenter/domain/valueobject"
)

// RepoCompany is an autogenerated mock type for the RepoCompany type
type RepoCompany struct {
	mock.Mock
}

// GetCompanyById provides a mock function with given fields: ctx, id
func (_m *RepoCompany) GetCompanyById(ctx context.Context, id string) (*entity.CompanyDTO, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.CompanyDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.CompanyDTO, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.CompanyDTO); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.CompanyDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListCompany provides a mock function with given fields: ctx, typeSearch, f
func (_m *RepoCompany) GetListCompany(ctx context.Context, typeSearch *valueobject.TypeSearch, f *filter.Filter) ([]*entity.CompanyDTO, error) {
	ret := _m.Called(ctx, typeSearch, f)

	var r0 []*entity.CompanyDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *valueobject.TypeSearch, *filter.Filter) ([]*entity.CompanyDTO, error)); ok {
		return rf(ctx, typeSearch, f)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *valueobject.TypeSearch, *filter.Filter) []*entity.CompanyDTO); ok {
		r0 = rf(ctx, typeSearch, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.CompanyDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *valueobject.TypeSearch, *filter.Filter) error); ok {
		r1 = rf(ctx, typeSearch, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepoCompany interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepoCompany creates a new instance of RepoCompany. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepoCompany(t mockConstructorTestingTNewRepoCompany) *RepoCompany {
	mock := &RepoCompany{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}