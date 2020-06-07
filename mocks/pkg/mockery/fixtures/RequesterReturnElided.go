// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// RequesterReturnElided is an autogenerated mock type for the RequesterReturnElided type
type RequesterReturnElided struct {
	mock.Mock
}

// Get provides a mock function with given fields: path
func (_m *RequesterReturnElided) Get(path string) (int, int, int, error) {
	ret := _m.Called(path)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(string) int); ok {
		r2 = rf(path)
	} else {
		r2 = ret.Get(2).(int)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(string) error); ok {
		r3 = rf(path)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}
