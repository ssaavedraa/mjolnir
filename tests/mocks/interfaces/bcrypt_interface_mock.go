// Code generated by mockery v2.43.2. DO NOT EDIT.

package interfaces_mocks

import mock "github.com/stretchr/testify/mock"

// MockBcryptInterface is an autogenerated mock type for the BcryptInterface type
type MockBcryptInterface struct {
	mock.Mock
}

type MockBcryptInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBcryptInterface) EXPECT() *MockBcryptInterface_Expecter {
	return &MockBcryptInterface_Expecter{mock: &_m.Mock}
}

// CompareHashAndPassword provides a mock function with given fields: hashedPassword, password
func (_m *MockBcryptInterface) CompareHashAndPassword(hashedPassword []byte, password []byte) error {
	ret := _m.Called(hashedPassword, password)

	if len(ret) == 0 {
		panic("no return value specified for CompareHashAndPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []byte) error); ok {
		r0 = rf(hashedPassword, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBcryptInterface_CompareHashAndPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompareHashAndPassword'
type MockBcryptInterface_CompareHashAndPassword_Call struct {
	*mock.Call
}

// CompareHashAndPassword is a helper method to define mock.On call
//   - hashedPassword []byte
//   - password []byte
func (_e *MockBcryptInterface_Expecter) CompareHashAndPassword(hashedPassword interface{}, password interface{}) *MockBcryptInterface_CompareHashAndPassword_Call {
	return &MockBcryptInterface_CompareHashAndPassword_Call{Call: _e.mock.On("CompareHashAndPassword", hashedPassword, password)}
}

func (_c *MockBcryptInterface_CompareHashAndPassword_Call) Run(run func(hashedPassword []byte, password []byte)) *MockBcryptInterface_CompareHashAndPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].([]byte))
	})
	return _c
}

func (_c *MockBcryptInterface_CompareHashAndPassword_Call) Return(_a0 error) *MockBcryptInterface_CompareHashAndPassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBcryptInterface_CompareHashAndPassword_Call) RunAndReturn(run func([]byte, []byte) error) *MockBcryptInterface_CompareHashAndPassword_Call {
	_c.Call.Return(run)
	return _c
}

// GenerateFromPassword provides a mock function with given fields: password, cost
func (_m *MockBcryptInterface) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	ret := _m.Called(password, cost)

	if len(ret) == 0 {
		panic("no return value specified for GenerateFromPassword")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, int) ([]byte, error)); ok {
		return rf(password, cost)
	}
	if rf, ok := ret.Get(0).(func([]byte, int) []byte); ok {
		r0 = rf(password, cost)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte, int) error); ok {
		r1 = rf(password, cost)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBcryptInterface_GenerateFromPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateFromPassword'
type MockBcryptInterface_GenerateFromPassword_Call struct {
	*mock.Call
}

// GenerateFromPassword is a helper method to define mock.On call
//   - password []byte
//   - cost int
func (_e *MockBcryptInterface_Expecter) GenerateFromPassword(password interface{}, cost interface{}) *MockBcryptInterface_GenerateFromPassword_Call {
	return &MockBcryptInterface_GenerateFromPassword_Call{Call: _e.mock.On("GenerateFromPassword", password, cost)}
}

func (_c *MockBcryptInterface_GenerateFromPassword_Call) Run(run func(password []byte, cost int)) *MockBcryptInterface_GenerateFromPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(int))
	})
	return _c
}

func (_c *MockBcryptInterface_GenerateFromPassword_Call) Return(_a0 []byte, _a1 error) *MockBcryptInterface_GenerateFromPassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBcryptInterface_GenerateFromPassword_Call) RunAndReturn(run func([]byte, int) ([]byte, error)) *MockBcryptInterface_GenerateFromPassword_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBcryptInterface creates a new instance of MockBcryptInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBcryptInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBcryptInterface {
	mock := &MockBcryptInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}