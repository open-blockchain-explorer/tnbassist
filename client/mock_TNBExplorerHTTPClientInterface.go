// Code generated by mockery v2.8.0. DO NOT EDIT.

package client

import (
	model "github.com/open-blockchain-explorer/tnbassist/model"
	mock "github.com/stretchr/testify/mock"
)

// MockTNBExplorerHTTPClientInterface is an autogenerated mock type for the TNBExplorerHTTPClientInterface type
type MockTNBExplorerHTTPClientInterface struct {
	mock.Mock
}

// PostStats provides a mock function with given fields: stats
func (_m *MockTNBExplorerHTTPClientInterface) PostStats(stats *model.LegacyStats) (int, []byte, error) {
	ret := _m.Called(stats)

	var r0 int
	if rf, ok := ret.Get(0).(func(*model.LegacyStats) int); ok {
		r0 = rf(stats)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(*model.LegacyStats) []byte); ok {
		r1 = rf(stats)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*model.LegacyStats) error); ok {
		r2 = rf(stats)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
