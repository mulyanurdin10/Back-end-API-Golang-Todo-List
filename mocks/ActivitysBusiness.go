// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	activitys "testcode/features/activitys"

	mock "github.com/stretchr/testify/mock"
)

// ActivityBusiness is an autogenerated mock type for the Business type
type ActivityBusiness struct {
	mock.Mock
}

// DeleteData provides a mock function with given fields: id
func (_m *ActivityBusiness) DeleteData(id int) (int, error) {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllData provides a mock function with given fields:
func (_m *ActivityBusiness) GetAllData() ([]activitys.Core, error) {
	ret := _m.Called()

	var r0 []activitys.Core
	if rf, ok := ret.Get(0).(func() []activitys.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]activitys.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetData provides a mock function with given fields: id
func (_m *ActivityBusiness) GetData(id int) (activitys.Core, int, error) {
	ret := _m.Called(id)

	var r0 activitys.Core
	if rf, ok := ret.Get(0).(func(int) activitys.Core); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(activitys.Core)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int) int); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// InsertData provides a mock function with given fields: insert
func (_m *ActivityBusiness) InsertData(insert activitys.Core) (activitys.Core, int, error) {
	ret := _m.Called(insert)

	var r0 activitys.Core
	if rf, ok := ret.Get(0).(func(activitys.Core) activitys.Core); ok {
		r0 = rf(insert)
	} else {
		r0 = ret.Get(0).(activitys.Core)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(activitys.Core) int); ok {
		r1 = rf(insert)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(activitys.Core) error); ok {
		r2 = rf(insert)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateData provides a mock function with given fields: id, insert
func (_m *ActivityBusiness) UpdateData(id int, insert activitys.Core) (activitys.Core, int, error) {
	ret := _m.Called(id, insert)

	var r0 activitys.Core
	if rf, ok := ret.Get(0).(func(int, activitys.Core) activitys.Core); ok {
		r0 = rf(id, insert)
	} else {
		r0 = ret.Get(0).(activitys.Core)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, activitys.Core) int); ok {
		r1 = rf(id, insert)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, activitys.Core) error); ok {
		r2 = rf(id, insert)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewActivityBusiness interface {
	mock.TestingT
	Cleanup(func())
}

// NewActivityBusiness creates a new instance of ActivityBusiness. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewActivityBusiness(t mockConstructorTestingTNewActivityBusiness) *ActivityBusiness {
	mock := &ActivityBusiness{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
