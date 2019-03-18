// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import contracts "github.com/guiyomh/charlatan/pkg/tree/contracts"
import mock "github.com/stretchr/testify/mock"

// Tree is an autogenerated mock type for the Tree type
type Tree struct {
	mock.Mock
}

// Add provides a mock function with given fields: node
func (_m *Tree) Add(node contracts.Node) {
	_m.Called(node)
}

// Find provides a mock function with given fields: _a0
func (_m *Tree) Find(_a0 string) contracts.Node {
	ret := _m.Called(_a0)

	var r0 contracts.Node
	if rf, ok := ret.Get(0).(func(string) contracts.Node); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(contracts.Node)
		}
	}

	return r0
}

// First provides a mock function with given fields:
func (_m *Tree) First() contracts.Node {
	ret := _m.Called()

	var r0 contracts.Node
	if rf, ok := ret.Get(0).(func() contracts.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(contracts.Node)
		}
	}

	return r0
}

// Last provides a mock function with given fields:
func (_m *Tree) Last() contracts.Node {
	ret := _m.Called()

	var r0 contracts.Node
	if rf, ok := ret.Get(0).(func() contracts.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(contracts.Node)
		}
	}

	return r0
}

// Walk provides a mock function with given fields: iterator, forward
func (_m *Tree) Walk(iterator contracts.Iterator, forward bool) {
	_m.Called(iterator, forward)
}