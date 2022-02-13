// Code generated by mockery v2.8.0. DO NOT EDIT.

package client

import mock "github.com/stretchr/testify/mock"

// MockGitHubGraphQLClientInterface is an autogenerated mock type for the GitHubGraphQLClientInterface type
type MockGitHubGraphQLClientInterface struct {
	mock.Mock
}

// CheckRateLimit provides a mock function with given fields:
func (_m *MockGitHubGraphQLClientInterface) CheckRateLimit() (*RateLimit, error) {
	ret := _m.Called()

	var r0 *RateLimit
	if rf, ok := ret.Get(0).(func() *RateLimit); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RateLimit)
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

// FetchNIssues provides a mock function with given fields: organization, repository, filters
func (_m *MockGitHubGraphQLClientInterface) FetchNIssues(organization string, repository string, filters Filters) (*Issues, error) {
	ret := _m.Called(organization, repository, filters)

	var r0 *Issues
	if rf, ok := ret.Get(0).(func(string, string, Filters) *Issues); ok {
		r0 = rf(organization, repository, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Issues)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, Filters) error); ok {
		r1 = rf(organization, repository, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WhoAmI provides a mock function with given fields:
func (_m *MockGitHubGraphQLClientInterface) WhoAmI() (*Viewer, error) {
	ret := _m.Called()

	var r0 *Viewer
	if rf, ok := ret.Get(0).(func() *Viewer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Viewer)
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
