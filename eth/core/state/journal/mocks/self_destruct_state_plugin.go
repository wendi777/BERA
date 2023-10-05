// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"
)

// SelfDestructStatePlugin is an autogenerated mock type for the selfDestructStatePlugin type
type SelfDestructStatePlugin struct {
	mock.Mock
}

type SelfDestructStatePlugin_Expecter struct {
	mock *mock.Mock
}

func (_m *SelfDestructStatePlugin) EXPECT() *SelfDestructStatePlugin_Expecter {
	return &SelfDestructStatePlugin_Expecter{mock: &_m.Mock}
}

// GetBalance provides a mock function with given fields: _a0
func (_m *SelfDestructStatePlugin) GetBalance(_a0 common.Address) *big.Int {
	ret := _m.Called(_a0)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(common.Address) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// SelfDestructStatePlugin_GetBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBalance'
type SelfDestructStatePlugin_GetBalance_Call struct {
	*mock.Call
}

// GetBalance is a helper method to define mock.On call
//   - _a0 common.Address
func (_e *SelfDestructStatePlugin_Expecter) GetBalance(_a0 interface{}) *SelfDestructStatePlugin_GetBalance_Call {
	return &SelfDestructStatePlugin_GetBalance_Call{Call: _e.mock.On("GetBalance", _a0)}
}

func (_c *SelfDestructStatePlugin_GetBalance_Call) Run(run func(_a0 common.Address)) *SelfDestructStatePlugin_GetBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *SelfDestructStatePlugin_GetBalance_Call) Return(_a0 *big.Int) *SelfDestructStatePlugin_GetBalance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SelfDestructStatePlugin_GetBalance_Call) RunAndReturn(run func(common.Address) *big.Int) *SelfDestructStatePlugin_GetBalance_Call {
	_c.Call.Return(run)
	return _c
}

// GetCodeHash provides a mock function with given fields: _a0
func (_m *SelfDestructStatePlugin) GetCodeHash(_a0 common.Address) common.Hash {
	ret := _m.Called(_a0)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(common.Address) common.Hash); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	return r0
}

// SelfDestructStatePlugin_GetCodeHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCodeHash'
type SelfDestructStatePlugin_GetCodeHash_Call struct {
	*mock.Call
}

// GetCodeHash is a helper method to define mock.On call
//   - _a0 common.Address
func (_e *SelfDestructStatePlugin_Expecter) GetCodeHash(_a0 interface{}) *SelfDestructStatePlugin_GetCodeHash_Call {
	return &SelfDestructStatePlugin_GetCodeHash_Call{Call: _e.mock.On("GetCodeHash", _a0)}
}

func (_c *SelfDestructStatePlugin_GetCodeHash_Call) Run(run func(_a0 common.Address)) *SelfDestructStatePlugin_GetCodeHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *SelfDestructStatePlugin_GetCodeHash_Call) Return(_a0 common.Hash) *SelfDestructStatePlugin_GetCodeHash_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SelfDestructStatePlugin_GetCodeHash_Call) RunAndReturn(run func(common.Address) common.Hash) *SelfDestructStatePlugin_GetCodeHash_Call {
	_c.Call.Return(run)
	return _c
}

// SubBalance provides a mock function with given fields: _a0, _a1
func (_m *SelfDestructStatePlugin) SubBalance(_a0 common.Address, _a1 *big.Int) {
	_m.Called(_a0, _a1)
}

// SelfDestructStatePlugin_SubBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubBalance'
type SelfDestructStatePlugin_SubBalance_Call struct {
	*mock.Call
}

// SubBalance is a helper method to define mock.On call
//   - _a0 common.Address
//   - _a1 *big.Int
func (_e *SelfDestructStatePlugin_Expecter) SubBalance(_a0 interface{}, _a1 interface{}) *SelfDestructStatePlugin_SubBalance_Call {
	return &SelfDestructStatePlugin_SubBalance_Call{Call: _e.mock.On("SubBalance", _a0, _a1)}
}

func (_c *SelfDestructStatePlugin_SubBalance_Call) Run(run func(_a0 common.Address, _a1 *big.Int)) *SelfDestructStatePlugin_SubBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(*big.Int))
	})
	return _c
}

func (_c *SelfDestructStatePlugin_SubBalance_Call) Return() *SelfDestructStatePlugin_SubBalance_Call {
	_c.Call.Return()
	return _c
}

func (_c *SelfDestructStatePlugin_SubBalance_Call) RunAndReturn(run func(common.Address, *big.Int)) *SelfDestructStatePlugin_SubBalance_Call {
	_c.Call.Return(run)
	return _c
}

// NewSelfDestructStatePlugin creates a new instance of SelfDestructStatePlugin. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSelfDestructStatePlugin(t interface {
	mock.TestingT
	Cleanup(func())
}) *SelfDestructStatePlugin {
	mock := &SelfDestructStatePlugin{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
