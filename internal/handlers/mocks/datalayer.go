// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/ogniloud/madr/internal/models"
)

// Datalayer is an autogenerated mock type for the Datalayer type
type Datalayer struct {
	mock.Mock
}

type Datalayer_Expecter struct {
	mock *mock.Mock
}

func (_m *Datalayer) EXPECT() *Datalayer_Expecter {
	return &Datalayer_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *Datalayer) CreateUser(ctx context.Context, user models.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Datalayer_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type Datalayer_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user models.User
func (_e *Datalayer_Expecter) CreateUser(ctx interface{}, user interface{}) *Datalayer_CreateUser_Call {
	return &Datalayer_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, user)}
}

func (_c *Datalayer_CreateUser_Call) Run(run func(ctx context.Context, user models.User)) *Datalayer_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.User))
	})
	return _c
}

func (_c *Datalayer_CreateUser_Call) Return(_a0 error) *Datalayer_CreateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Datalayer_CreateUser_Call) RunAndReturn(run func(context.Context, models.User) error) *Datalayer_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// SignInUser provides a mock function with given fields: ctx, username, password
func (_m *Datalayer) SignInUser(ctx context.Context, username string, password string) (string, error) {
	ret := _m.Called(ctx, username, password)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, username, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, username, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Datalayer_SignInUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SignInUser'
type Datalayer_SignInUser_Call struct {
	*mock.Call
}

// SignInUser is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
//   - password string
func (_e *Datalayer_Expecter) SignInUser(ctx interface{}, username interface{}, password interface{}) *Datalayer_SignInUser_Call {
	return &Datalayer_SignInUser_Call{Call: _e.mock.On("SignInUser", ctx, username, password)}
}

func (_c *Datalayer_SignInUser_Call) Run(run func(ctx context.Context, username string, password string)) *Datalayer_SignInUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Datalayer_SignInUser_Call) Return(_a0 string, _a1 error) *Datalayer_SignInUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Datalayer_SignInUser_Call) RunAndReturn(run func(context.Context, string, string) (string, error)) *Datalayer_SignInUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewDatalayer creates a new instance of Datalayer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDatalayer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Datalayer {
	mock := &Datalayer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
