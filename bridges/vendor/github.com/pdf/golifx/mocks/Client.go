// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import time "time"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// GetRetryInterval provides a mock function with given fields:
func (_m *Client) GetRetryInterval() *time.Duration {
	ret := _m.Called()

	var r0 *time.Duration
	if rf, ok := ret.Get(0).(func() *time.Duration); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*time.Duration)
		}
	}

	return r0
}

// GetTimeout provides a mock function with given fields:
func (_m *Client) GetTimeout() *time.Duration {
	ret := _m.Called()

	var r0 *time.Duration
	if rf, ok := ret.Get(0).(func() *time.Duration); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*time.Duration)
		}
	}

	return r0
}