// Code generated by mockery v2.14.1. DO NOT EDIT.

package reader

import (
	model "github.com/pingcap/tiflow/cdc/model"
	mock "github.com/stretchr/testify/mock"
)

// mockFileReader is an autogenerated mock type for the fileReader type
type mockFileReader struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *mockFileReader) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Read provides a mock function with given fields:
func (_m *mockFileReader) Read() (*model.RedoLog, error) {
	ret := _m.Called()

	var r0 *model.RedoLog
	if rf, ok := ret.Get(0).(func() *model.RedoLog); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RedoLog)
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

type mockConstructorTestingTnewMockFileReader interface {
	mock.TestingT
	Cleanup(func())
}

// newMockFileReader creates a new instance of mockFileReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockFileReader(t mockConstructorTestingTnewMockFileReader) *mockFileReader {
	mock := &mockFileReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
