// Code generated by mockery v2.26.1. DO NOT EDIT.

package repomocks

import (
	models "github.com/mrtkmynsndev/go-test-mockery-demo/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// ProductRepositoryInterface is an autogenerated mock type for the ProductRepositoryInterface type
type ProductRepositoryInterface struct {
	mock.Mock
}

// Add provides a mock function with given fields: product
func (_m *ProductRepositoryInterface) Add(product models.Product) error {
	ret := _m.Called(product)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Product) error); ok {
		r0 = rf(product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewProductRepositoryInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductRepositoryInterface creates a new instance of ProductRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductRepositoryInterface(t mockConstructorTestingTNewProductRepositoryInterface) *ProductRepositoryInterface {
	mock := &ProductRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
