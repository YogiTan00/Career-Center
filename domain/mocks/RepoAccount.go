// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	account "CareerCenter/domain/entity/account"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// RepoAccount is an autogenerated mock type for the RepoAccount type
type RepoAccount struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: ctx, data
func (_m *RepoAccount) CreateAccount(ctx context.Context, data *account.Account) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *account.Account) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *RepoAccount) GetByEmail(ctx context.Context, email string) (*account.AccountDTO, error) {
	ret := _m.Called(ctx, email)

	var r0 *account.AccountDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*account.AccountDTO, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *account.AccountDTO); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*account.AccountDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePassword provides a mock function with given fields: ctx, email, password
func (_m *RepoAccount) UpdatePassword(ctx context.Context, email string, password string) error {
	ret := _m.Called(ctx, email, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, email, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRole provides a mock function with given fields: ctx, data
func (_m *RepoAccount) UpdateRole(ctx context.Context, data *account.Account) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *account.Account) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepoAccount interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepoAccount creates a new instance of RepoAccount. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepoAccount(t mockConstructorTestingTNewRepoAccount) *RepoAccount {
	mock := &RepoAccount{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
