// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Fooer is an autogenerated mock type for the Fooer type
type Fooer struct {
	mock.Mock
}

// Bar provides a mock function with given fields: f
func (_m *Fooer) Bar(f func([]int)) {
	_m.Called(f)
}

// Baz provides a mock function with given fields: path
func (_m *Fooer) Baz(path string) func(string) string {
	ret := _m.Called(path)

	var r0 func(string) string
	if rf, ok := ret.Get(0).(func(string) func(string) string); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(string) string)
		}
	}

	return r0
}

// Foo provides a mock function with given fields: f
func (_m *Fooer) Foo(f func(string) string) error {
	ret := _m.Called(f)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(string) string) error); ok {
		r0 = rf(f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
