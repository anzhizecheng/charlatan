// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Faker is an autogenerated mock type for the Faker type
type Faker struct {
	mock.Mock
}

// Generate provides a mock function with given fields: data
func (_m *Faker) Generate(data string) interface{} {
	ret := _m.Called(data)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}
