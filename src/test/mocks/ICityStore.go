// Code generated by mockery v2.22.1. DO NOT EDIT.
package mocks

import (
	mock "github.com/stretchr/testify/mock"

	domain "tinderclone_back/src/pkg/domain"
)

// ICityStore is an autogenerated mock type for the ICityStore type
type ICityStore struct {
	mock.Mock
}

// IsCityAlreadyAvailable provides a mock function with given fields: cityName
func (_m *ICityStore) IsCityAlreadyAvailable(cityName string) bool {
	ret := _m.Called(cityName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(cityName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SelectCityByName provides a mock function with given fields: cityName
func (_m *ICityStore) SelectCityByName(cityName string) *domain.City {
	ret := _m.Called(cityName)

	var r0 *domain.City
	if rf, ok := ret.Get(0).(func(string) *domain.City); ok {
		r0 = rf(cityName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.City)
		}
	}

	return r0
}

type mockConstructorTestingTNewICityStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewICityStore creates a new instance of ICityStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewICityStore(t mockConstructorTestingTNewICityStore) *ICityStore {
	mock := &ICityStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}