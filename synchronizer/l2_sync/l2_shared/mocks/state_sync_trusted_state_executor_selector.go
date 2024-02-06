// Code generated by mockery. DO NOT EDIT.

package mock_l2_shared

import (
	state "github.com/0xPolygonHermez/zkevm-node/state"
	mock "github.com/stretchr/testify/mock"
)

// stateSyncTrustedStateExecutorSelector is an autogenerated mock type for the stateSyncTrustedStateExecutorSelector type
type stateSyncTrustedStateExecutorSelector struct {
	mock.Mock
}

type stateSyncTrustedStateExecutorSelector_Expecter struct {
	mock *mock.Mock
}

func (_m *stateSyncTrustedStateExecutorSelector) EXPECT() *stateSyncTrustedStateExecutorSelector_Expecter {
	return &stateSyncTrustedStateExecutorSelector_Expecter{mock: &_m.Mock}
}

// GetForkIDInMemory provides a mock function with given fields: forkId
func (_m *stateSyncTrustedStateExecutorSelector) GetForkIDInMemory(forkId uint64) *state.ForkIDInterval {
	ret := _m.Called(forkId)

	if len(ret) == 0 {
		panic("no return value specified for GetForkIDInMemory")
	}

	var r0 *state.ForkIDInterval
	if rf, ok := ret.Get(0).(func(uint64) *state.ForkIDInterval); ok {
		r0 = rf(forkId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.ForkIDInterval)
		}
	}

	return r0
}

// stateSyncTrustedStateExecutorSelector_GetForkIDInMemory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetForkIDInMemory'
type stateSyncTrustedStateExecutorSelector_GetForkIDInMemory_Call struct {
	*mock.Call
}

// GetForkIDInMemory is a helper method to define mock.On call
//   - forkId uint64
func (_e *stateSyncTrustedStateExecutorSelector_Expecter) GetForkIDInMemory(forkId interface{}) *stateSyncTrustedStateExecutorSelector_GetForkIDInMemory_Call {
	return &stateSyncTrustedStateExecutorSelector_GetForkIDInMemory_Call{Call: _e.mock.On("GetForkIDInMemory", forkId)}
}

func (_c *stateSyncTrustedStateExecutorSelector_GetForkIDInMemory_Call) Run(run func(forkId uint64)) *stateSyncTrustedStateExecutorSelector_GetForkIDInMemory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *stateSyncTrustedStateExecutorSelector_GetForkIDInMemory_Call) Return(_a0 *state.ForkIDInterval) *stateSyncTrustedStateExecutorSelector_GetForkIDInMemory_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *stateSyncTrustedStateExecutorSelector_GetForkIDInMemory_Call) RunAndReturn(run func(uint64) *state.ForkIDInterval) *stateSyncTrustedStateExecutorSelector_GetForkIDInMemory_Call {
	_c.Call.Return(run)
	return _c
}

// newStateSyncTrustedStateExecutorSelector creates a new instance of stateSyncTrustedStateExecutorSelector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newStateSyncTrustedStateExecutorSelector(t interface {
	mock.TestingT
	Cleanup(func())
}) *stateSyncTrustedStateExecutorSelector {
	mock := &stateSyncTrustedStateExecutorSelector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}